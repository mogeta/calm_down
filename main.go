package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"

	"github.com/gorilla/websocket"
)

type (
	talk struct {
		Count     int    `json:"count"`
		WrongWord string `json:"wrong_word"`
	}
)

var (
	upgrader = websocket.Upgrader{}
	counter  = 0
	word     = ""
)

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	notifiedCoutner := 0
	for {
		// Write
		if notifiedCoutner<counter{
			notifiedCoutner = counter
			data,_ := json.Marshal(t)
			err := ws.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				c.Logger().Error(err)
			}
		}


		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

var t = &talk{}
func receive(c echo.Context) error {

	if err := c.Bind(t); err != nil {
		return err
	}
	counter = t.Count
	word = t.WrongWord
	fmt.Println(t.WrongWord)
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "public")
	e.GET("/ws", hello)
	e.POST("/receive", receive)
	e.Logger.Fatal(e.Start(":1323"))
}
