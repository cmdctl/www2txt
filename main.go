package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/gocolly/colly"
)

type ArrayFlags []string

func (i *ArrayFlags) String() string {
	return "my string representation"
}

func (i *ArrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var tags ArrayFlags

func main() {

	flag.Var(&tags, "tag", "html tags to extract")
	flag.Parse()

  fmt.Println(tags)
  if len(tags) == 0 {
    fmt.Println("No tags specified")
    return
  }

	www := os.Args[1]
	fmt.Println("Visiting", www)

	parsed_url, err := url.Parse(www)
	fmt.Println(parsed_url.Host)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(fmt.Sprintf("%s.txt", parsed_url.Host))
	defer file.Close()

	c := colly.NewCollector()
  for _, tag := range tags {
    c.OnHTML(tag, func(e *colly.HTMLElement) {
      file.WriteString(e.Text)
    })
  }
	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		parsed_href, err := url.Parse(href)
		if err != nil {
			fmt.Printf("Error parsing url: %s", href)
			return
		}
		if parsed_href.Host == "" || parsed_href.Host == parsed_url.Host {
			e.Request.Visit(e.Attr("href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(www)
}
