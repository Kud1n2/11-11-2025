package main

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
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
var URLresponses []URLresponse
var ListsRequests []ListsRequest

func addURL(context *gin.Context) {
	var request URLrequest
	if err := context.BindJSON(&request); err == nil {
		URLrequests = append(URLrequests, request)
		CheckAvailable()
		context.IndentedJSON(http.StatusOK, URLresponses[len(URLresponses)-1])
	} else {
		var request ListsRequest

		if err := context.BindJSON(&request); err != nil {
			return
		}
		ListsRequests = append(ListsRequests, request)
		context.IndentedJSON(http.StatusOK, ListsRequests)
	}
}

func getURLs(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, URLresponses)
}

func CheckAvailable() {
	dict := make(map[string]string)
	if len(URLrequests) > 0 {
		used := URLrequests[len(URLrequests)-1]
		for i := 0; i < len(used.Links); i++ {
			_, err := net.Dial("tcp", used.Links[i]+":http")
			if err != nil {
				dict[used.Links[i]] = "not available"
			} else {
				dict[used.Links[i]] = "available"
			}

		}
		URLresponses = append(URLresponses, URLresponse{Links: dict, Links_num: len(URLrequests)})
	}
}

func main() {
	router := gin.Default()
	router.POST("/webservice", addURL)
	router.GET("/webservice", getURLs)
	// router.GET("/webservice", getPDF)

	router.Run("localhost:1010")
}
