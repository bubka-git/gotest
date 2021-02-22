package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GetResponse struct {
	Origin string            `json:"origin"`
	URL    string            `json: "url"`
	Header map[string]string `json:"headers"`
}

func (r *GetResponse) ToString() string {
	s := fmt.Sprintf("GET Response\nOrigin IP: %s \nRequest URL: %s\n",
		r.Origin, r.URL)
	for k, v := range r.Header {
		s += fmt.Sprintf("Header: %s = %s\n", k, v)
	}
	return s
}

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln("Unable to get response")
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to read content")
	}

	/*fmt.Println(string(content))*/
	respContent := &GetResponse{}
	json.Unmarshal(content, respContent)
	fmt.Println(respContent.ToString())

}
