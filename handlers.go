package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type tiny_url struct {
	URL  string `json:"url"`
	Tiny string `json:"tiny_url"`
}

var tinyUrls = make(map[string]string)
var recoURLPrefix string = "http://localhost:8080/"

func getTinyURLs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tinyUrls)
}

func createTinyURL(c *gin.Context) {
	var newTinyURL tiny_url

	if err := c.BindJSON(&newTinyURL); err != nil {
		return
	}

	newTinyURL.Tiny = getSHA1Hash(&newTinyURL.URL)
	_, exists := tinyUrls[newTinyURL.Tiny]

	if exists {
		c.IndentedJSON(http.StatusAlreadyReported, recoURLPrefix+newTinyURL.Tiny)
		return
	}

	tinyUrls[newTinyURL.Tiny] = newTinyURL.URL
	c.IndentedJSON(http.StatusCreated, recoURLPrefix+newTinyURL.Tiny)
}

func getURL(c *gin.Context) {
	tinyURL := c.Param("tiny")

	results, exists := tinyUrls[tinyURL]

	if exists {
		c.Redirect(http.StatusMovedPermanently, results)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Tiny URL not found"})
	}
}
