package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var questionList = map[string]question{
	"1": {"How are you today?"},
	"2": {"What are you up to?"},
	"3": {"Can I come and visit?"},
	"4": {"What is your favorite color?"},
	"5": {"What do you need help with most often?"},
}

type question struct {
	Question string `json: "How are you today?"`
}

func main() {
	router := gin.Default()
	router.GET("/", getAllQuestions)
	router.GET("/:id", getQuestionByID)
	router.Run("localhost:8080")
}

func getAllQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, questionList)
}

func getQuestionByID(c *gin.Context) {
	params := c.Param("id")
	//if questionList[params] doesnt exist return error
	for k, _ := range questionList {
		if params == k {
			c.JSON(http.StatusOK, questionList[params])
			return
		}
		if params != k {
			fmt.Println(k)
			fmt.Println(params)
		}
	}
	c.JSON(http.StatusBadRequest, "message: id was not found")
}
