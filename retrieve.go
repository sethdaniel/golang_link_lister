package main

import (
	"fmt"
	//"github.com/jackdanger/collectlinks"
	"io/ioutil"
	"net/http"
)

func retrieveMain() {
	resp, err := http.Get("https://www.reddit.com/r/nottheonion/")

	fmt.Println("error", err)

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println("read err is ", err)

	fmt.Println(string(body))
}
