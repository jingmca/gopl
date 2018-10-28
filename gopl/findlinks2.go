package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func findLinks(u string) ([]string, error) {
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting :%s:%s", u, resp.StatusCode)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	mess := make(chan string)
	mess2 := make(chan string)

	for _, url := range os.Args[1:] {
		go func(u string) {

			mess <- u
			mess <- "OK"

			for {
				_, ok := <-mess2
				if !ok {
					close(mess)
					return
				}
			}
		}(url)
	}
	a := 0

	for url := range mess {
		if url == "OK" {
			a++
			if a == 1 {
				close(mess2)
				break
			}
		} else {
			fmt.Println(url)
		}

	}

}
