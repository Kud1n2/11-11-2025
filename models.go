package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type URLrequest struct {
	Links []string `json:"links"`
}

type URLresponse struct {
	Links     map[string]string `json:"links"`
	Links_num int               `json:"links_num"`
}

type ListsRequest struct {
	Links_list []int `json:"links_list"`
}

var URLrequests []URLrequest
var ListsRequests []ListsRequest
var URLresponses []URLresponse

func makeURLresponses() []URLresponse {
	file, err := os.ReadFile("log.json")
	var URLresponses []URLresponse
	if err != nil {
		URLresponses = []URLresponse{}
	} else {
		err = json.Unmarshal(file, &URLresponses)
		if err != nil {
			fmt.Println("Ошибка чтения")
		}
	}
	return URLresponses
}
