package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	r "router"
)

var router = r.NewRouter()

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln("Unable to get response")
	}
	router.Process(resp)
}

func init() {
	router.Register(200, func(resp *http.Response) {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("Unable to read content")
		}
		fmt.Println(string(content))
	})
	router.Register(404, func(resp *http.Response) {
		log.Fatalln("Not Found (404): ", r.Request.URL.String())

	})
}
