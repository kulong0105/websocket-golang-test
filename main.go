package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// webSocket returns text format
func textApi(c *gin.Context) {
	//Upgrade get request to webSocket protocol
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Println(err)
		return
	}
	defer ws.Close()
	//Read data in ws
	mt, message, err := ws.ReadMessage()
	if err != nil {
		log.Println("error read message")
		log.Println(err)
		return
	}

	log.Println("receive data", message)

	//Write ws data, pong 10 times
	var count = 0
	for {
		count++
		if count > 10 {
			break
		}

		message = []byte(string(message) + " " + strconv.Itoa(count))
		//message := []byte(strconv.Itoa(count))
		//err := ws.WriteMessage(2, message)
		err := ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("error write message: " + err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}

//webSocket returns json format
func jsonApi(c *gin.Context) {
	//Upgrade get request to webSocket protocol
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Println(err)
		return
	}
	defer ws.Close()
	var data struct {
		A string `json:"a"`
		B int    `json:"b"`
	}
	//Read data in ws
	err = ws.ReadJSON(&data)
	if err != nil {
		log.Println("error read json")
		log.Println(err)
		return
	}

	log.Println("receive data", data)

	//Write ws data, pong 10 times
	var count = 0
	for {
		count++
		if count > 10 {
			break
		}

		err = ws.WriteJSON(struct {
			A string `json:"a"`
			B int    `json:"b"`
			C int    `json:"c"`
		}{
			A: data.A,
			B: data.B,
			C: count,
		})
		if err != nil {
			log.Println("error write json: " + err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	r := gin.Default()
	r.GET("/json", jsonApi)
	r.GET("/text", textApi)

	r.Run(":8080")
}
