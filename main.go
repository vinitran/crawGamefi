package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const regrexString = "<script id=\"__NEXT_DATA__\" type=\"application\\/json\">(.+?)<\\/script>"

func getDataFromUrl(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}

	myRegex := regexp.MustCompile(regrexString)
	var imgTags = myRegex.FindAllStringSubmatch(string(content), -1)

	return imgTags[0][1], nil
}

// <script id="__NEXT_DATA__" type="application\/json">(.+?)<\/script>
func main() {
	var slugs []string
	for i := 1; i < 18; i++ {
		link := fmt.Sprintf("https://playtoearngames.com/games/page/%d", i)
		dtCrawl, err := getDataFromUrl(link)
		var dataCrawl DataCrawl
		err = json.Unmarshal([]byte(dtCrawl), &dataCrawl)
		if err != nil {
			log.Fatal(err)
		}

		game := dataCrawl.Props.PageProps.GamesData.Items

		for _, gm := range game {
			slugs = append(slugs, gm.Slug)
		}
	}

	var dataCsv [][]string
	for _, slug := range slugs {
		detailLink := fmt.Sprintf("https://playtoearngames.com/games/%s", slug)
		dataWeb, err := getDataFromUrl(detailLink)
		if err != nil {
			log.Fatal(err)
		}

		var dataCrawlDetail DataGameDetail
		err = json.Unmarshal([]byte(dataWeb), &dataCrawlDetail)
		if err != nil {
			fmt.Println(detailLink)
			fmt.Println(err)
			continue
		}

		content := dataCrawlDetail.Props.PageProps.Data
		title := strings.Replace(content.Title, " - Game Review", " ", -1)
		var rowCsv []string
		rowCsv = append(rowCsv, fmt.Sprintf("what is the genre of %s?", title))
		rowCsv = append(rowCsv, content.Genre)
		dataCsv = append(dataCsv, rowCsv)
	}

	exportCsv(dataCsv)
}

func exportCsv(data [][]string) error {
	file, err := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}

	defer file.Close()
	csvWriter := csv.NewWriter(file)
	csvWriter.WriteAll(data)
	csvWriter.Flush()
	return nil
}
