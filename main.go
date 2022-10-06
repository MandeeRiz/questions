package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func randomNum() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(len(questionList))
}

var questionList = map[string]question{
	"0": {"How are you today?"},
	"1": {"What are you up to?"},
	"2": {"Can I come and visit?"},
	"3": {"What is your favorite color?"},
	"4": {"What do you need help with most often?"},
}

type question struct {
	Question string `json: "How are you today?"`
}

func main() {
	router := gin.Default()
	router.GET("/", getAllQuestions)
	router.GET("/:id", getQuestionByID)
	router.GET("/surprise", getSurpriseQuestion)
	router.Run("localhost:8080")
}

func getAllQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, questionList)
}

func getQuestionByID(c *gin.Context) {
	params := c.Param("id")
	for k, _ := range questionList {
		if params == k {
			c.JSON(http.StatusOK, questionList[params])
			return
		}
	}
	c.JSON(http.StatusBadRequest, "message: id was not found")
}

func getSurpriseQuestion(c *gin.Context) {
	num := randomNum()
	numm := fmt.Sprintf("%v", num)
	c.JSON(http.StatusOK, questionList[numm])
}
