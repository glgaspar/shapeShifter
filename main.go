package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool
	Message string
	Data    interface{}
}

func main() {
	router := gin.Default()
	setRoutes(router)

	router.Run("0.0.0.0:8080")
}

func processor(c *gin.Context) {
	var err error
	var bodyJson []byte
	var body Request

	bodyJson, err = io.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			Response{
				Status:  false,
				Message: "Request body processing error: " + err.Error(),
				Data:    nil})
		return
	}

	err = json.Unmarshal(bodyJson, &body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			Response{
				Status:  false,
				Message: "Error parsing body json: " + err.Error(),
				Data:    nil})
		return
	}

	if len(body.Data) == 0 {
		c.IndentedJSON(
			http.StatusBadRequest,
			Response{
				Status:  false,
				Message: "No data to shape shift",
				Data:    nil})
		return
	}

	if len(body.ResponseLang) == 0 {
		c.IndentedJSON(http.StatusBadRequest,
			Response{
				Status:  false,
				Message: "No language to shape shift to",
				Data:    nil})
		return
	}

	for lang := 0; lang < len(body.ResponseLang); lang++ {
		if err = body.Parse(lang); err != nil {
			c.IndentedJSON(http.StatusUnprocessableEntity,
				Response{
					Status:  false,
					Message: "Shape shifting error: " + err.Error(),
					Data:    nil})
		}
	}

	c.IndentedJSON(http.StatusOK,
		Response{
			Status:  true,
			Message: "Shape shifting successful",
			Data:    body.ResponseData})

}

func setRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) { c.Done() })
	router.POST("/", processor)
}
