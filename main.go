package main

import (
	"fmt"
	"net/url"

	"github.com/Sylphy0052/go_crawler"
)

func main() {
	url, _ := url.Parse("http://example.com")
	err := go_crawler.Start(url, 1, "./test/")
	if err != nil {
		fmt.Println(err)
	}
}
