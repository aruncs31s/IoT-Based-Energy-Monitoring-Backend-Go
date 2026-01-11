package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Readings struct {
	Voltage float64 `json:"voltage"`
	Current float64 `json:"current"`
}

func main() {
	r := gin.Default()

	r.GET("/hi", HelloWorld)
	r.Run()

}
func HelloWorld(c *gin.Context) {

	random := rand.Float64()
	reading := Readings{
		Voltage: random,
		Current: 10.0,
	}
	c.JSON(http.StatusOK, reading)
}
