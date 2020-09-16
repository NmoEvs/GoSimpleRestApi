package main

import (
	"os"
	"strconv"

	lorem "github.com/drhodes/golorem"

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
	r.GET("/logs/:number/create", func(c *gin.Context) {
		nbLogs, err := strconv.Atoi(c.Param("number"))
		if err != nil {
			panic("Oops something goes wrong")
		}
		uuid, err := uuid.NewRandom()
		if err != nil {
			panic("Oops something goes wrong")
		}

		for i := 0; i < nbLogs; i++ {
			log.WithField("DiscusionID", uuid)

			switch {
			case i%2 == 0:
				log.Debug(lorem.Sentence(5, 10))
				break
			case i%3 == 0:
				log.Info("Info Log Appears !!!")
				break
			case i == 4:
				log.Warn("Warn Log Appears !!!")
				break
			case i == 5:
				log.Error(lorem.Paragraph(20, 40))
				break

			default:
				log.Trace(lorem.Word(5, 10))
			}
		}

		c.JSON(200, gin.H{
			"message": "Log Created",
		})
	})
	r.Run(":5000")
}
