package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	lorem "github.com/drhodes/golorem"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var waitGroup sync.WaitGroup
var data chan string

func main() {

	data = make(chan string)

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

		for i := 0; i < nbLogs; i++ {
			GenerateRandomLog(i)
		}

		c.JSON(200, gin.H{
			"message": "Log Created",
		})
	})

	r.GET("/load/start", func(c *gin.Context) {
		waitGroup.Add(1)
		go worker()
	})
	r.GET("/load/stop", func(c *gin.Context) {
		close(data)
		waitGroup.Wait()
	})

	r.Run(":8080")
}

func worker() {
	fmt.Println("Goroutine worker is now starting...")
	defer func() {
		fmt.Println("Destroying the worker...")
		waitGroup.Done()
	}()
	for {
		select {

		case <-data:
			return
		default:
			for i := 0; i < rand.Intn(500); i++ {
				GenerateRandomLog(i)
			}
			time.Sleep(time.Second * 10)
		}
	}
}

//GenerateRandomLog - Create A random log level
func GenerateRandomLog(i int) {

	uuid, err := uuid.NewRandom()
	if err != nil {
		panic("Oops something goes wrong")
	}

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
