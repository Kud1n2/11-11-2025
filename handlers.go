package main

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlerRequest(context *gin.Context) {
	body, err := context.GetRawData()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Reading body error"})
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(body)) //создание копии запроса
	var requestList ListsRequest
	if err := context.BindJSON(&requestList); err == nil && len(requestList.Links_list) > 0 {
		ListsRequests = append(ListsRequests, requestList)
		err := makePDF(requestList.Links_list)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Making PDF error"})
		}
		context.Header("Content-Disposition", "attachment; filename=file.pdf")
		context.Header("Content-Type", "application/pdf")
		context.File("file.pdf")
		return
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	var requestURL URLrequest
	if err := context.BindJSON(&requestURL); err == nil && len(requestURL.Links) > 0 {
		URLrequests = append(URLrequests, requestURL)
		writeToLog(CheckAvailable())
		context.IndentedJSON(http.StatusOK, URLresponses[len(URLresponses)-1])
		return
	}
	context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
}

func getURLs(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, URLresponses)
}
