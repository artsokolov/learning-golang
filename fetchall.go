package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const outputPath = "./output/"

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, link := range os.Args[1:] {
		go fetch(link, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("Time spent: %0.2fs\n", time.Since(start).Seconds())
}

func fetch(link string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(link)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	saveResponseToFile(resp, link, ch)

	ch <- fmt.Sprintf("%0.2fs %s", time.Since(start).Seconds(), link)
}

func saveResponseToFile(resp *http.Response, link string, ch chan string) {
	fileName := generateFileName(link)
	f, err := os.Create(fileName)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	err = f.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
}

func generateFileName(link string) string {
	convertedDomain := strings.ReplaceAll(link, ".", "_")
	convertedDomain = strings.ReplaceAll(convertedDomain, "https://", "")
	convertedDomain = strings.ReplaceAll(convertedDomain, "http://", "")

	return outputPath + strconv.FormatInt(time.Now().UTC().UnixNano(), 5) + "_" + convertedDomain
}
