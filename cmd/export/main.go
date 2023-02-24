package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"robanohashi/cmd/import/wanikani"
)

type Pages struct {
	PerPage int    `json:"per_page"`
	NextUrl string `json:"next_url"`
}

type Body struct {
	Url        string `json:"url"`
	Pages      Pages  `json:"pages"`
	TotalCount int    `json:"total_count"`
	Data       []any  `json:"data"`
}

func main() {
	wanikani := http.Client{}

	f, err := os.Create("subjects.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	req := createRequest("https://api.wanikani.com/v2/subjects")

	for {
		res, err := wanikani.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode != http.StatusOK {
			log.Fatalf("Failed to fetch subjects with status code %d", res.StatusCode)
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		jsonRes := Body{}
		err = json.Unmarshal(body, &jsonRes)

		if err != nil {
			log.Fatal(err)
		}

		appendPageToFile(jsonRes.Data, f)

		if jsonRes.Pages.NextUrl != "" {
			nextUrl, err := url.Parse(jsonRes.Pages.NextUrl)
			if err != nil {
				log.Fatal(err)
			}
			req.URL = nextUrl
			fmt.Printf("Finished page... %s\n", jsonRes.Url)
		} else {
			break
		}
	}
}

func createRequest(url string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Wanikani-Revision", "20170710")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("WANIKANI_API_KEY")))

	return req
}

func appendPageToFile(data []any, file *os.File) {
	for _, item := range data {
		// download character image for radicals
		if item.(map[string]any)["object"] == "radical" {
			radical := wanikani.Subject[wanikani.Radical]{}
			b, err := json.Marshal(item)
			if err != nil {
				log.Fatal(err)
			}
			err = json.Unmarshal(b, &radical)
			if err != nil {
				log.Fatal(err)
			}
			url := getCharacterImageUrl(radical)
			if url != "" {
				item.(map[string]any)["data"].(map[string]any)["character_svg_image"] = fetchCharacterImage(url)
			}
		}
		b, err := json.Marshal(item)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := file.Write(b); err != nil {
			log.Fatal(err)
		}
		if _, err := file.Write([]byte("\n")); err != nil {
			log.Fatal(err)
		}
	}
}

func getCharacterImageUrl(subject wanikani.Subject[wanikani.Radical]) string {
	for _, img := range subject.Data.CharacterImages {
		if img.ContentType == "image/svg+xml" && img.Metadata.InlineStyles {
			return img.URL
		}
	}
	return ""
}

func fetchCharacterImage(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
