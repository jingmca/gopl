package gopl

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

//GFetch will start many goroutines for fetching url from cmd
func GFetch() {

	start := time.Now()
	ch := make(chan string)

	for _, url := range flag.Args() {
		go fetch(url, ch)
	}

	for range flag.Args() {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		ch <- fmt.Sprintf("fetch %s status error : %d", url, resp.StatusCode)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("%s reading error: %s", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%s %.2fs %7dbytes \n-----\n%v\n-----", url, secs, nbytes, resp.Header)
}
