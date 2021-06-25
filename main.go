package main

import (
	"fmt"
	"net/http"
)

func checkResStatus(chanel chan string, link string) {
	res, err := http.Get(link)
	if err != nil {
		chanel <- fmt.Sprintf("Request %s is failed, status: %s, prototype: %s, server name: %s.", link, res.Status, res.Proto, res.TLS.ServerName)
		return
	}
	chanel <- fmt.Sprintf("Request %s is success, status: %s, prototype: %s, server name: %s.", link, res.Status, res.Proto, res.TLS.ServerName)
}
func main() {
	listSite := []string{"https://www.youtube.com", "https://www.facebook.com"}
	chanel := make(chan string)
	for _, site := range listSite {
		go checkResStatus(chanel, site)
	}
	responses := map[int]string{}
	for index, _ := range listSite {
		responses[index] = <-chanel
	}
	for _, res := range responses {
		fmt.Println(res)
	}
}
