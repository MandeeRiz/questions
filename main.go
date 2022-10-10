package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func randomNum() int {
	return rand.Intn(len(questionList))
}

var questionList = map[string]question{
	"0": {0, "How are you today?"},
	"1": {1, "What are you up to?"},
	"2": {2, "Can I come and visit?"},
	"3": {3, "What is your favorite color?"},
	"4": {4, "What do you need help with most often?"},
}

type question struct {
	Id       int    `json: 0`
	Question string `json: "How are you today?"`
}

func main() {
	router := gin.Default()
	router.GET("/questions", getAllQuestions)
	router.GET("question/:id", getQuestionByID)
	router.GET("question/random", getRandomQuestion)
	router.Run("localhost:8080")
}

func getAllQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, questionList)
}

func getQuestionByID(c *gin.Context) {
	param := c.Param("id")
	_, exists := questionList[param]
	if exists {
		c.JSON(http.StatusOK, questionList[param])
		return
	}
	c.JSON(http.StatusBadRequest, "message: id was not found")
}

func getRandomQuestion(c *gin.Context) {
	num := randomNum()
	stringNum := fmt.Sprint(num)
	c.JSON(http.StatusOK, questionList[stringNum])
}
