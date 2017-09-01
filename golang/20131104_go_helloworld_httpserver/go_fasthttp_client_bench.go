
package main

import (
	"log"
	// "net"
	// "io"
	// "io/ioutil"
	// "net/http"
	"sync"
	// "time"

	"github.com/valyala/fasthttp"
)

var (
	n   = 1000000
	c   = 100
	url = "http://127.0.0.1:8081"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < c; i++ {
		wg.Add(1)
		go func() {
			worker()
			wg.Done()
		}()
	}
	log.Printf("started...")

	wg.Wait()
}

func worker() {
	// client := &http.Client{}
	client := &fasthttp.Client{}
	buf := make([]byte, 1024)
	for i := 0; i < n/c; i++ {
		// res, err := client.Get(url)
		_, _, err := client.Get(buf, url)
		/*
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Printf("newrequest error: %v", err)
			continue
		}
		res, err := client.Do(req)
		*/
		if err != nil {
			log.Printf("err: %v", err)
			continue
		}
		// io.Copy(ioutil.Discard, res.Body)
		// res.Body.Close()
	}
	log.Printf("concurrent finish")
}

