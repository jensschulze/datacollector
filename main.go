package main

import (
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	queue := NewQueue(1000000) // TODO: Make configurable
	defer queue.RemoveBatch()

	go func(queue *Queue) {
		ip := make([]byte, 4)
		udpAddr := net.UDPAddr{IP: ip, Port: 8080, Zone: ""} // TODO: Make configurable

		udpConn, err := net.ListenUDP("udp", &udpAddr)
		if err != nil {
			panic(err)
		}
		defer udpConn.Close()

		for {
			buffer := make([]byte, 1024)
			length, err := udpConn.Read(buffer)
			if err != nil {
				panic(err)
			}

			if length > 0 {
				queue.Insert(string(buffer))
			}
		}
	}(queue)

	r := gin.Default()
	r.GET("/metrics", func(c *gin.Context) {
		items := queue.RemoveBatch()
		body := strings.Join((*items)[:], "\n")

		c.String(http.StatusOK, body)
	})

	r.Run(":9090") // TODO: Make configurable. Listen and serve on 0.0.0.0:9090 (for windows "localhost:9090")
}
