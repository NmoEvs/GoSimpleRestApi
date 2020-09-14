package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/debug", func(c *gin.Context) {
		uuid, err := uuid.NewRandom()
		if err != nil {
			panic("Oops something goes wrong")
		}
		log.WithField("DiscusionID", uuid).Debug("Debug message appears !!")
		c.JSON(200, gin.H{
			"message": "debug",
		})
	})
	r.Run(":8080")
}
