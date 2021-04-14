package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	url := "http://localhost:8080"

	maxConn := make(chan bool, 200)
	wg := &sync.WaitGroup{}

	count := 0
	start := time.Now()
	for maxRequest := 0; maxRequest < 10000; maxRequest++ {
		wg.Add(1)
		maxConn <- true
		go func() {
			defer wg.Done()

			r, err := http.Get(url)
			if err != nil {
				log.Println(err)
				return
			}
			defer r.Body.Close()

			count++
			<-maxConn
		}()
		wg.Wait()
		end := time.Now()
		log.Printf("%d回のリクエストに成功しました\n", count)
		log.Printf("%f秒処理に時間がかかりました\n", (end.Sub(start)).Seconds())
	}
}
