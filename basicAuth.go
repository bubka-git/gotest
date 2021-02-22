package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/passw0rd", nil)
	if err != nil {
		log.Fatalln("Unable to create a request")
	}
	req.SetBasicAuth("user", "passw0rd")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("Unable to get response")
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to read content")
	}
	fmt.Println(string(content))
	fmt.Println(resp.StatusCode)
}
