package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type URLrequest struct {
	Links []string `json:"links"`
}

type URLresponse struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

var URLrequests []URLrequest

func addURL(context *gin.Context) {
	var request URLrequest

	if err := context.BindJSON(&request); err != nil {
		return
	}
	URLrequests = append(URLrequests, request)
	CheckAvailable()
	context.IndentedJSON(http.StatusOK, URLrequests)

}
func getURLs(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, URLresponses)
}

var URLresponses []URLresponse

func CheckAvailable() {
	if len(URLrequests) > 0 {
		used := URLrequests[len(URLrequests)-1]
		for i := 0; i < len(used.Links); i++ {
			_, err := http.Get(used.Links[i])
			if err != nil {
				URLresponses = append(URLresponses, URLresponse{URL: used.Links[i], Status: "not available"})
			} else {
				URLresponses = append(URLresponses, URLresponse{URL: used.Links[i], Status: "available"})
			}
		}
	}
}

func main() {
	router := gin.Default()
	router.POST("/webservice", addURL)
	router.GET("/webservice", getURLs)

	router.Run("localhost:1010")
}
