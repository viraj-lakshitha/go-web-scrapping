package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gocolly/colly"
)

type Repository struct {
	Name string `json:"name"`
	Description string `json:"description"`
	URL string `json:"url"`
}

func main() {
	repos := []Repository{}

	c := colly.NewCollector(
		colly.AllowedDomains("github.com"),
	)

	c.OnHTML("li[itemprop=owns]", func(h *colly.HTMLElement) {
		repos = append(repos, Repository{
			Name: h.ChildText("h3.wb-break-all a"),
			Description: h.ChildText("p[itemprop=description]"),
			URL: fmt.Sprintf("https://github.com%s", h.ChildAttr("h3.wb-break-all a","href")),
		})
	})

	c.Visit("https://github.com/viraj-lakshitha?tab=repositories")

	marshalRepos, marshalErr := json.Marshal(repos)
	if marshalErr != nil {
		fmt.Print(marshalErr)
		return
	}
	
	writeErr :=	ioutil.WriteFile("repo-details.json", marshalRepos, 0777)
	if writeErr != nil {
		fmt.Println(writeErr)
		return
	}
}