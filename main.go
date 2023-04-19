package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"golang.org/x/sync/semaphore"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

const regrexString = "<script type=\"application\\/ld\\+json\">(.+?)<\\/script>"

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
	if len(imgTags) < 1 {
		return "", nil
	}

	return imgTags[1][1], nil
}

// <script type="application\/ld\+json">(.+?)<\/script>

func getDataDetail(slug string) ([][]string, error) {
	var dataList [][]string
	link := fmt.Sprintf("https://chainplay.gg/games/%s/", slug)
	dtCrawl, err := getDataFromUrl(link)
	if err != nil {
		return nil, err
	}

	var dataCrawl DataCrawl
	err = json.Unmarshal([]byte(dtCrawl), &dataCrawl)
	if err != nil {
		return nil, err
	}

	listQuest := dataCrawl.MainEntity

	if len(listQuest) == 0 {
		return nil, nil
	}

	for _, ques := range listQuest {
		var data []string
		data = append(data, ques.Name)
		data = append(data, ques.AcceptedAnswer.Text)
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func GetDataAll() []string {
	postBody, _ := json.Marshal(map[string]string{
		"pageNumber": "1",
		"pageSize":   "2181",
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("https://api.chainplay.gg/api/app/project/get-games", "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var dtCrawlAll DataAll
	err = json.Unmarshal(body, &dtCrawlAll)
	if err != nil {
		log.Fatal(err)
	}

	var code []string
	for _, cd := range dtCrawlAll.Payload.Games {
		code = append(code, cd.Code)
	}
	return code
}

func main() {

	dtList := GetDataAll()
	fmt.Println(dtList)

	sem := semaphore.NewWeighted(10)
	var dataCsv [][]string
	for _, slug := range dtList {
		err := sem.Acquire(context.Background(), 1)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(slug)

		go func(slug string) {
			data, err := getDataDetail(slug)
			if err != nil {
				fmt.Println("abcasd")
				sem.Release(1)
				return
			}

			if data == nil {
				fmt.Println("abcasd")
				sem.Release(1)
				return
			}

			for _, dt := range data {
				dataCsv = append(dataCsv, dt)
			}

			sem.Release(1)
		}(slug)
	}

	//var slugs []string
	//for i := 1; i < 18; i++ {
	//	link := fmt.Sprintf("https://playtoearngames.com/games/page/%d", i)
	//	dtCrawl, err := getDataFromUrl(link)
	//	var dataCrawl DataCrawl
	//	err = json.Unmarshal([]byte(dtCrawl), &dataCrawl)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	game := dataCrawl.Props.PageProps.GamesData.Items
	//
	//	for _, gm := range game {
	//		slugs = append(slugs, gm.Slug)
	//	}
	//}
	//
	//maxGoroutines := 10
	//guard := make(chan struct{}, maxGoroutines)
	//var wg sync.WaitGroup
	//
	//var dataCsv [][]string
	//for _, slug := range slugs {
	//	wg.Add(1)
	//	go func(slug string) {
	//		detailLink := fmt.Sprintf("https://playtoearngames.com/games/%s", slug)
	//		dataWeb, err := getDataFromUrl(detailLink)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//
	//		var dataCrawlDetail DataGameDetail
	//		err = json.Unmarshal([]byte(dataWeb), &dataCrawlDetail)
	//		if err != nil {
	//			fmt.Println(detailLink)
	//			fmt.Println(err)
	//			return
	//		}
	//
	//		content := dataCrawlDetail.Props.PageProps.Data
	//		title := strings.Replace(content.Title, " - Game Review", "", -1)
	//		var rowCsv []string
	//		data := content.TokenName
	//		//data = strings.Replace(data, " - Game Developer", "", -1)
	//		if data == "" {
	//			fmt.Println("not exist")
	//			wg.Done()
	//			<-guard
	//			return
	//		}
	//
	//		rowCsv = append(rowCsv, fmt.Sprintf("token name of %s", title))
	//		rowCsv = append(rowCsv, data)
	//		dataCsv = append(dataCsv, rowCsv)
	//		fmt.Println("appending")
	//		wg.Done()
	//		<-guard
	//	}(slug)
	//}
	//
	//close(guard) // This tells the goroutines there's nothing else to do
	//wg.Wait()
	//
	exportCsv(dataCsv)
}

func exportCsv(data [][]string) error {
	file, err := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	csvWriter := csv.NewWriter(file)

	defer file.Close()
	csvWriter.WriteAll(data)
	csvWriter.Flush()
	return nil
}
