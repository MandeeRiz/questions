package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var questionList = map[int]question{
	1: {"How are you today?"},
	2: {"What are you up to?"},
	3: {"Can I come and visit?"},
	4: {"What is your favorite color?"},
	5: {"What do you need help with most often?"},
}

type question struct {
	Question string `json: "How are you today?"`
}

func main() {
	router := gin.Default()
	router.GET("/", getAllQuestions)
	router.Run("localhost:8080")
}

func getAllQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, questionList)
}
