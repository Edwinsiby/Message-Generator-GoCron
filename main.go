package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	c.AddFunc("*/5 * * * * *", generateMessage)
	c.AddFunc("*/10 * * * * *", generateMessage2)

	c.Start()

	select {}
}

func generateMessage() {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	message := "Message generated for every 5 seconds"
	fmt.Printf("[%s] %s\n", timestamp, message)
}
func generateMessage2() {

	timestamp := time.Now().Format("2006-01-02 15:04:05")

	message := "Message generated for every 10 seconds"

	fmt.Printf("[%s] %s\n", timestamp, message)
}

// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jasonlvhit/gocron"
// )

// func main() {
// 	r := gin.Default()
// 	r.LoadHTMLGlob("templates/*.html")
// 	r.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.html", nil)
// 	})

// 	r.GET("/events", func(c *gin.Context) {
// 		c.Header("Content-Type", "text/event-stream")
// 		c.Header("Cache-Control", "no-cache")
// 		c.Header("Connection", "keep-alive")
// 		c.Header("Access-Control-Allow-Origin", "*")

// 		flusher, ok := c.Writer.(http.Flusher)
// 		if !ok {
// 			http.Error(c.Writer, "Streaming unsupported", http.StatusInternalServerError)
// 			return
// 		}

// 		messageChan := make(chan string)

// 		gocron.Every(5).Seconds().Do(func() {
// 			messageChan <- "Hello! This is the message."
// 		})

// 		gocron.Start()

// 		for {
// 			select {
// 			case message := <-messageChan:
// 				fmt.Fprintf(c.Writer, "data: %s\n\n", message)
// 				flusher.Flush()
// 			case <-c.Writer.CloseNotify():
// 				fmt.Println("page closed")
// 				return
// 			}
// 		}
// 	})

// 	r.Run(":8080")
// }
