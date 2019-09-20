package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// People is
type People struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// Response is
type Response struct {
	Message string `json:"message"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		res, err := http.Get("http://03be40f4.ngrok.io")
		if err != nil {
			log.Println(err)
		}
		defer res.Body.Close()
		var listPeople []People
		json.NewDecoder(res.Body).Decode(&listPeople)
		c.JSON(200, gin.H{
			"data": listPeople,
		})
	})
	r.POST("/", func(c *gin.Context) {
		body, _ := c.GetRawData()
		// var raw People
		// _ = json.Unmarshal(body, &raw)
		res, err := http.Post("http://test-fail-thirdparty:2222", "application/json", bytes.NewBuffer(body))
		if err != nil {
			log.Println(err)
		}
		defer res.Body.Close()
		var hasil Response
		json.NewDecoder(res.Body).Decode(&hasil)
		c.JSON(200, gin.H{
			"data": hasil,
		})
	})
	r.Run(":2121")
}
