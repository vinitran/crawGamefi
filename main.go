package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"
)

//type Data struct {
//	Props struct {
//		PageProps struct {
//			Data struct {
//				MostPopular []struct {
//					Title          string `json:"title"`
//					Slug           string `json:"slug"`
//					Description    string `json:"description"`
//					Image          string `json:"image"`
//					ImageWebp      string `json:"image_webp"`
//					Developer      string `json:"developer"`
//					Live           bool   `json:"live"`
//					CommunitySize  string `json:"community_size"`
//					Rating         string `json:"rating"`
//					LargeImage     string `json:"large_image"`
//					MediumImage    string `json:"medium_image"`
//					ThumbnailImage string `json:"thumbnail_image"`
//				} `json:"mostPopular"`
//				EmergingGames []struct {
//					Title          string `json:"title"`
//					Slug           string `json:"slug"`
//					Description    string `json:"description"`
//					Image          string `json:"image"`
//					ImageWebp      string `json:"image_webp"`
//					Developer      string `json:"developer"`
//					Live           bool   `json:"live"`
//					CommunitySize  string `json:"community_size"`
//					Rating         string `json:"rating"`
//					LargeImage     string `json:"large_image"`
//					MediumImage    string `json:"medium_image"`
//					ThumbnailImage string `json:"thumbnail_image"`
//				} `json:"emerging_games"`
//				Developers []struct {
//					Title       string `json:"title"`
//					Slug        string `json:"slug"`
//					Image       string `json:"image"`
//					ImageWebp   string `json:"image_webp"`
//					Description string `json:"description"`
//					MediumImage string `json:"medium_image"`
//					Verified    bool   `json:"verified"`
//					GamesAmount int    `json:"gamesAmount"`
//				} `json:"developers"`
//				Seo struct {
//					ID                       int    `json:"id"`
//					Permalink                string `json:"permalink"`
//					PermalinkHash            string `json:"permalink_hash"`
//					ObjectID                 int    `json:"object_id"`
//					ObjectType               string `json:"object_type"`
//					ObjectSubType            string `json:"object_sub_type"`
//					AuthorID                 int    `json:"author_id"`
//					PostParent               int    `json:"post_parent"`
//					Title                    string `json:"title"`
//					Description              string `json:"description"`
//					BreadcrumbTitle          string `json:"breadcrumb_title"`
//					PostStatus               string `json:"post_status"`
//					IsPublic                 any    `json:"is_public"`
//					IsProtected              int    `json:"is_protected"`
//					HasPublicPosts           any    `json:"has_public_posts"`
//					NumberOfPages            any    `json:"number_of_pages"`
//					Canonical                any    `json:"canonical"`
//					PrimaryFocusKeyword      string `json:"primary_focus_keyword"`
//					PrimaryFocusKeywordScore int    `json:"primary_focus_keyword_score"`
//					ReadabilityScore         int    `json:"readability_score"`
//					IsCornerstone            int    `json:"is_cornerstone"`
//					IsRobotsNoindex          any    `json:"is_robots_noindex"`
//					IsRobotsNofollow         int    `json:"is_robots_nofollow"`
//					IsRobotsNoarchive        any    `json:"is_robots_noarchive"`
//					IsRobotsNoimageindex     any    `json:"is_robots_noimageindex"`
//					IsRobotsNosnippet        any    `json:"is_robots_nosnippet"`
//					TwitterTitle             string `json:"twitter_title"`
//					TwitterImage             string `json:"twitter_image"`
//					TwitterDescription       string `json:"twitter_description"`
//					TwitterImageID           string `json:"twitter_image_id"`
//					TwitterImageSource       string `json:"twitter_image_source"`
//					OpenGraphTitle           string `json:"open_graph_title"`
//					OpenGraphDescription     string `json:"open_graph_description"`
//					OpenGraphImage           string `json:"open_graph_image"`
//					OpenGraphImageID         string `json:"open_graph_image_id"`
//					OpenGraphImageSource     string `json:"open_graph_image_source"`
//					OpenGraphImageMeta       struct {
//						Width    int    `json:"width"`
//						Height   int    `json:"height"`
//						Filesize int    `json:"filesize"`
//						URL      string `json:"url"`
//						Path     string `json:"path"`
//						Size     string `json:"size"`
//						ID       int    `json:"id"`
//						Alt      string `json:"alt"`
//						Pixels   int    `json:"pixels"`
//						Type     string `json:"type"`
//					} `json:"open_graph_image_meta"`
//					LinkCount                   int       `json:"link_count"`
//					IncomingLinkCount           int       `json:"incoming_link_count"`
//					ProminentWordsVersion       int       `json:"prominent_words_version"`
//					CreatedAt                   time.Time `json:"created_at"`
//					UpdatedAt                   time.Time `json:"updated_at"`
//					BlogID                      int       `json:"blog_id"`
//					Language                    any       `json:"language"`
//					Region                      any       `json:"region"`
//					SchemaPageType              any       `json:"schema_page_type"`
//					SchemaArticleType           any       `json:"schema_article_type"`
//					HasAncestors                int       `json:"has_ancestors"`
//					EstimatedReadingTimeMinutes int       `json:"estimated_reading_time_minutes"`
//					Version                     int       `json:"version"`
//					ObjectLastModified          string    `json:"object_last_modified"`
//					ObjectPublishedAt           string    `json:"object_published_at"`
//				} `json:"seo"`
//				News []struct {
//					Title          string `json:"title"`
//					Slug           string `json:"slug"`
//					Image          string `json:"image"`
//					Excerpt        string `json:"excerpt"`
//					Created        string `json:"created"`
//					UpdatedAt      string `json:"updated_at"`
//					Author         string `json:"author"`
//					AuthorAvatar   string `json:"author_avatar"`
//					LargeImage     string `json:"large_image"`
//					MediumImage    string `json:"medium_image"`
//					ThumbnailImage string `json:"thumbnail_image"`
//				} `json:"news"`
//				Videos []struct {
//					Title          string `json:"title"`
//					Slug           string `json:"slug"`
//					Excerpt        string `json:"excerpt"`
//					Created        string `json:"created"`
//					Author         string `json:"author"`
//					AuthorAvatar   string `json:"author_avatar"`
//					LargeImage     string `json:"large_image"`
//					MediumImage    string `json:"medium_image"`
//					Youtube        string `json:"youtube"`
//					ThumbnailImage string `json:"thumbnail_image"`
//				} `json:"videos"`
//			} `json:"data"`
//			URL  string   `json:"url"`
//			Name []string `json:"name"`
//		} `json:"pageProps"`
//		NSsp bool `json:"__N_SSP"`
//	} `json:"props"`
//	Page  string `json:"page"`
//	Query struct {
//	} `json:"query"`
//	BuildID      string `json:"buildId"`
//	IsFallback   bool   `json:"isFallback"`
//	DynamicIds   []int  `json:"dynamicIds"`
//	Gssp         bool   `json:"gssp"`
//	ScriptLoader []any  `json:"scriptLoader"`
//}

type DataCrawl struct {
	Props struct {
		PageProps struct {
			URL         string `json:"url"`
			HomePageURL string `json:"homePageUrl"`
			GamesData   struct {
				Items []struct {
					Title          string `json:"title"`
					Slug           string `json:"slug"`
					Description    string `json:"description"`
					Image          string `json:"image"`
					ImageWebp      string `json:"image_webp"`
					Developer      string `json:"developer"`
					Live           bool   `json:"live"`
					CommunitySize  string `json:"community_size"`
					Rating         string `json:"rating"`
					LargeImage     string `json:"large_image"`
					MediumImage    string `json:"medium_image"`
					ThumbnailImage string `json:"thumbnail_image"`
				} `json:"items"`
				Links struct {
					Next string `json:"next"`
					Prev string `json:"prev"`
				} `json:"links"`
				Filters []struct {
					Name string `json:"name"`
					Item []struct {
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Active bool   `json:"active"`
						Count  int    `json:"count"`
						ID     int    `json:"id"`
					} `json:"item"`
				} `json:"filters"`
				Total int `json:"total"`
				Seo   struct {
					ID                       int    `json:"id"`
					Permalink                string `json:"permalink"`
					PermalinkHash            string `json:"permalink_hash"`
					ObjectID                 int    `json:"object_id"`
					ObjectType               string `json:"object_type"`
					ObjectSubType            string `json:"object_sub_type"`
					AuthorID                 int    `json:"author_id"`
					PostParent               int    `json:"post_parent"`
					Title                    string `json:"title"`
					Description              string `json:"description"`
					BreadcrumbTitle          string `json:"breadcrumb_title"`
					PostStatus               string `json:"post_status"`
					IsPublic                 any    `json:"is_public"`
					IsProtected              int    `json:"is_protected"`
					HasPublicPosts           any    `json:"has_public_posts"`
					NumberOfPages            any    `json:"number_of_pages"`
					Canonical                any    `json:"canonical"`
					PrimaryFocusKeyword      string `json:"primary_focus_keyword"`
					PrimaryFocusKeywordScore int    `json:"primary_focus_keyword_score"`
					ReadabilityScore         int    `json:"readability_score"`
					IsCornerstone            int    `json:"is_cornerstone"`
					IsRobotsNoindex          any    `json:"is_robots_noindex"`
					IsRobotsNofollow         int    `json:"is_robots_nofollow"`
					IsRobotsNoarchive        any    `json:"is_robots_noarchive"`
					IsRobotsNoimageindex     any    `json:"is_robots_noimageindex"`
					IsRobotsNosnippet        any    `json:"is_robots_nosnippet"`
					TwitterTitle             string `json:"twitter_title"`
					TwitterImage             string `json:"twitter_image"`
					TwitterDescription       string `json:"twitter_description"`
					TwitterImageID           string `json:"twitter_image_id"`
					TwitterImageSource       string `json:"twitter_image_source"`
					OpenGraphTitle           string `json:"open_graph_title"`
					OpenGraphDescription     string `json:"open_graph_description"`
					OpenGraphImage           string `json:"open_graph_image"`
					OpenGraphImageID         string `json:"open_graph_image_id"`
					OpenGraphImageSource     string `json:"open_graph_image_source"`
					OpenGraphImageMeta       struct {
						Width    int    `json:"width"`
						Height   int    `json:"height"`
						Filesize int    `json:"filesize"`
						Path     string `json:"path"`
						URL      string `json:"url"`
						Size     string `json:"size"`
						ID       int    `json:"id"`
						Alt      string `json:"alt"`
						Pixels   int    `json:"pixels"`
						Type     string `json:"type"`
					} `json:"open_graph_image_meta"`
					LinkCount                   int       `json:"link_count"`
					IncomingLinkCount           int       `json:"incoming_link_count"`
					ProminentWordsVersion       int       `json:"prominent_words_version"`
					CreatedAt                   time.Time `json:"created_at"`
					UpdatedAt                   time.Time `json:"updated_at"`
					BlogID                      int       `json:"blog_id"`
					Language                    any       `json:"language"`
					Region                      any       `json:"region"`
					SchemaPageType              any       `json:"schema_page_type"`
					SchemaArticleType           any       `json:"schema_article_type"`
					HasAncestors                int       `json:"has_ancestors"`
					EstimatedReadingTimeMinutes int       `json:"estimated_reading_time_minutes"`
					Version                     int       `json:"version"`
					ObjectLastModified          string    `json:"object_last_modified"`
					ObjectPublishedAt           string    `json:"object_published_at"`
				} `json:"seo"`
			} `json:"gamesData"`
			GetParam int `json:"getParam"`
			Links    struct {
				PreviousLink string `json:"previousLink"`
				NextLink     string `json:"nextLink"`
			} `json:"links"`
			TotalPages int      `json:"totalPages"`
			URLFilters []any    `json:"Url_filters"`
			Names      []string `json:"Names"`
			IsMobile   bool     `json:"isMobile"`
		} `json:"pageProps"`
		NSsp bool `json:"__N_SSP"`
	} `json:"props"`
	Page  string `json:"page"`
	Query struct {
		Index string `json:"index"`
	} `json:"query"`
	BuildID      string `json:"buildId"`
	IsFallback   bool   `json:"isFallback"`
	DynamicIds   []int  `json:"dynamicIds"`
	Gssp         bool   `json:"gssp"`
	ScriptLoader []any  `json:"scriptLoader"`
}

type DataGameDetail struct {
	Props struct {
		PageProps struct {
			Data struct {
				Title             string    `json:"title"`
				Content           string    `json:"content"`
				YoutubeURL        string    `json:"youtube_url"`
				Image             string    `json:"image"`
				Developer         string    `json:"developer"`
				DeveloperSlug     string    `json:"developer_slug"`
				DeveloperVerified string    `json:"developer_verified"`
				Type              string    `json:"type"`
				Genre             string    `json:"genre"`
				Platform          string    `json:"platform"`
				Chains            string    `json:"chains"`
				GameStatus        string    `json:"game_status"`
				GameURL           string    `json:"game_url"`
				CoingeckoID       string    `json:"coingecko_id"`
				Twitter           string    `json:"twitter"`
				TwitterFollowers  string    `json:"twitter_followers"`
				Discord           string    `json:"discord"`
				DiscordSize       string    `json:"discord_size"`
				DiscordUpdated    string    `json:"discord_updated"`
				Token             string    `json:"token"`
				TokenUpdated      string    `json:"token_updated"`
				Dau               string    `json:"dau"`
				Whitepaper        string    `json:"whitepaper"`
				Telegram          string    `json:"telegram"`
				TelegramUpdated   string    `json:"telegram_updated"`
				TelegramCount     string    `json:"telegram_count"`
				CommunitySize     string    `json:"community_size"`
				Rating            string    `json:"rating"`
				Live              bool      `json:"live"`
				Verified          bool      `json:"verified"`
				StartDate         time.Time `json:"start_date"`
				EndDate           time.Time `json:"end_date"`
				TokenName         string    `json:"token_name"`
				Nfts              string    `json:"nfts"`
				Votes             struct {
					Up   int `json:"up"`
					Down int `json:"down"`
				} `json:"votes"`
				UpdatedAt      time.Time `json:"updated_at"`
				ImageGallery   []string  `json:"image_gallery"`
				CommunityPosts []any     `json:"community_posts"`
				MightLike      []struct {
					Title          string `json:"title"`
					Slug           string `json:"slug"`
					Description    string `json:"description"`
					Image          string `json:"image"`
					ImageWebp      string `json:"image_webp"`
					Developer      string `json:"developer"`
					Live           bool   `json:"live"`
					CommunitySize  string `json:"community_size"`
					Rating         string `json:"rating"`
					LargeImage     string `json:"large_image"`
					MediumImage    string `json:"medium_image"`
					ThumbnailImage string `json:"thumbnail_image"`
				} `json:"might_like"`
				Next string `json:"next"`
				Prev string `json:"prev"`
				Seo  struct {
					ID                       int    `json:"id"`
					Permalink                string `json:"permalink"`
					PermalinkHash            string `json:"permalink_hash"`
					ObjectID                 int    `json:"object_id"`
					ObjectType               string `json:"object_type"`
					ObjectSubType            string `json:"object_sub_type"`
					AuthorID                 int    `json:"author_id"`
					PostParent               int    `json:"post_parent"`
					Title                    any    `json:"title"`
					Description              string `json:"description"`
					BreadcrumbTitle          string `json:"breadcrumb_title"`
					PostStatus               string `json:"post_status"`
					IsPublic                 any    `json:"is_public"`
					IsProtected              int    `json:"is_protected"`
					HasPublicPosts           any    `json:"has_public_posts"`
					NumberOfPages            any    `json:"number_of_pages"`
					Canonical                any    `json:"canonical"`
					PrimaryFocusKeyword      string `json:"primary_focus_keyword"`
					PrimaryFocusKeywordScore int    `json:"primary_focus_keyword_score"`
					ReadabilityScore         int    `json:"readability_score"`
					IsCornerstone            int    `json:"is_cornerstone"`
					IsRobotsNoindex          any    `json:"is_robots_noindex"`
					IsRobotsNofollow         int    `json:"is_robots_nofollow"`
					IsRobotsNoarchive        any    `json:"is_robots_noarchive"`
					IsRobotsNoimageindex     any    `json:"is_robots_noimageindex"`
					IsRobotsNosnippet        any    `json:"is_robots_nosnippet"`
					TwitterTitle             any    `json:"twitter_title"`
					TwitterImage             string `json:"twitter_image"`
					TwitterDescription       string `json:"twitter_description"`
					TwitterImageID           string `json:"twitter_image_id"`
					TwitterImageSource       string `json:"twitter_image_source"`
					OpenGraphTitle           any    `json:"open_graph_title"`
					OpenGraphDescription     string `json:"open_graph_description"`
					OpenGraphImage           string `json:"open_graph_image"`
					OpenGraphImageID         string `json:"open_graph_image_id"`
					OpenGraphImageSource     string `json:"open_graph_image_source"`
					OpenGraphImageMeta       struct {
						Width    int    `json:"width"`
						Height   int    `json:"height"`
						Filesize int    `json:"filesize"`
						URL      string `json:"url"`
						Path     string `json:"path"`
						Size     string `json:"size"`
						ID       int    `json:"id"`
						Alt      string `json:"alt"`
						Pixels   int    `json:"pixels"`
						Type     string `json:"type"`
					} `json:"open_graph_image_meta"`
					LinkCount                   int       `json:"link_count"`
					IncomingLinkCount           int       `json:"incoming_link_count"`
					ProminentWordsVersion       int       `json:"prominent_words_version"`
					CreatedAt                   time.Time `json:"created_at"`
					UpdatedAt                   time.Time `json:"updated_at"`
					BlogID                      int       `json:"blog_id"`
					Language                    any       `json:"language"`
					Region                      any       `json:"region"`
					SchemaPageType              any       `json:"schema_page_type"`
					SchemaArticleType           any       `json:"schema_article_type"`
					HasAncestors                int       `json:"has_ancestors"`
					EstimatedReadingTimeMinutes int       `json:"estimated_reading_time_minutes"`
					Version                     int       `json:"version"`
					ObjectLastModified          string    `json:"object_last_modified"`
					ObjectPublishedAt           string    `json:"object_published_at"`
				} `json:"seo"`
				Faq []struct {
					Question string `json:"question"`
					Answer   string `json:"answer"`
				} `json:"faq"`
			} `json:"data"`
			DeviceType string `json:"deviceType"`
			SeoData    struct {
				ID                       int    `json:"id"`
				Permalink                string `json:"permalink"`
				PermalinkHash            string `json:"permalink_hash"`
				ObjectID                 int    `json:"object_id"`
				ObjectType               string `json:"object_type"`
				ObjectSubType            string `json:"object_sub_type"`
				AuthorID                 int    `json:"author_id"`
				PostParent               int    `json:"post_parent"`
				Title                    any    `json:"title"`
				Description              string `json:"description"`
				BreadcrumbTitle          string `json:"breadcrumb_title"`
				PostStatus               string `json:"post_status"`
				IsPublic                 any    `json:"is_public"`
				IsProtected              int    `json:"is_protected"`
				HasPublicPosts           any    `json:"has_public_posts"`
				NumberOfPages            any    `json:"number_of_pages"`
				Canonical                any    `json:"canonical"`
				PrimaryFocusKeyword      string `json:"primary_focus_keyword"`
				PrimaryFocusKeywordScore int    `json:"primary_focus_keyword_score"`
				ReadabilityScore         int    `json:"readability_score"`
				IsCornerstone            int    `json:"is_cornerstone"`
				IsRobotsNoindex          any    `json:"is_robots_noindex"`
				IsRobotsNofollow         int    `json:"is_robots_nofollow"`
				IsRobotsNoarchive        any    `json:"is_robots_noarchive"`
				IsRobotsNoimageindex     any    `json:"is_robots_noimageindex"`
				IsRobotsNosnippet        any    `json:"is_robots_nosnippet"`
				TwitterTitle             any    `json:"twitter_title"`
				TwitterImage             string `json:"twitter_image"`
				TwitterDescription       string `json:"twitter_description"`
				TwitterImageID           string `json:"twitter_image_id"`
				TwitterImageSource       string `json:"twitter_image_source"`
				OpenGraphTitle           any    `json:"open_graph_title"`
				OpenGraphDescription     string `json:"open_graph_description"`
				OpenGraphImage           string `json:"open_graph_image"`
				OpenGraphImageID         string `json:"open_graph_image_id"`
				OpenGraphImageSource     string `json:"open_graph_image_source"`
				OpenGraphImageMeta       struct {
					Width    int    `json:"width"`
					Height   int    `json:"height"`
					Filesize int    `json:"filesize"`
					URL      string `json:"url"`
					Path     string `json:"path"`
					Size     string `json:"size"`
					ID       int    `json:"id"`
					Alt      string `json:"alt"`
					Pixels   int    `json:"pixels"`
					Type     string `json:"type"`
				} `json:"open_graph_image_meta"`
				LinkCount                   int       `json:"link_count"`
				IncomingLinkCount           int       `json:"incoming_link_count"`
				ProminentWordsVersion       int       `json:"prominent_words_version"`
				CreatedAt                   time.Time `json:"created_at"`
				UpdatedAt                   time.Time `json:"updated_at"`
				BlogID                      int       `json:"blog_id"`
				Language                    any       `json:"language"`
				Region                      any       `json:"region"`
				SchemaPageType              any       `json:"schema_page_type"`
				SchemaArticleType           any       `json:"schema_article_type"`
				HasAncestors                int       `json:"has_ancestors"`
				EstimatedReadingTimeMinutes int       `json:"estimated_reading_time_minutes"`
				Version                     int       `json:"version"`
				ObjectLastModified          string    `json:"object_last_modified"`
				ObjectPublishedAt           string    `json:"object_published_at"`
			} `json:"seoData"`
			Host     string `json:"host"`
			PathName string `json:"pathName"`
			IP       string `json:"ip"`
			Slug     string `json:"slug"`
			Proto    string `json:"proto"`
		} `json:"pageProps"`
		NSsp bool `json:"__N_SSP"`
	} `json:"props"`
	Page  string `json:"page"`
	Query struct {
		Game string `json:"game"`
	} `json:"query"`
	BuildID      string `json:"buildId"`
	IsFallback   bool   `json:"isFallback"`
	DynamicIds   []int  `json:"dynamicIds"`
	Gssp         bool   `json:"gssp"`
	ScriptLoader []any  `json:"scriptLoader"`
}

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

	fmt.Print(len(slugs))

	//detailLink := fmt.Sprintf("https://playtoearngames.com/games/%s", "heroes-chained")
	//dataWeb, err := getDataFromUrl(detailLink)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var dataCrawlDetail DataGameDetail
	//err = json.Unmarshal([]byte(dataWeb), &dataCrawlDetail)
	//if err != nil {
	//	log.Fatal(err)
	//}

}
