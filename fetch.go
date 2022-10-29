package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	links := os.Args[1:]

	if len(links) == 0 {
		fmt.Println("You have to provide one or more links")
		os.Exit(1)
	}

	for _, link := range links {
		if strings.HasPrefix(link, "http://") {
			link = link[7:]
		}

		if !strings.HasPrefix(link, "https://") {
			link = "https://" + link
		}

		resp, err := http.Get(link)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %s: %v\n", link, err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %s: %v\n", link, err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "\nResponse status: %s\n", resp.Status)
	}
}
