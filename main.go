package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Request struct {
	App       string `json:"app,omitempty"`
	Sender    string `json:"sender,omitempty"`
	Message   string `json:"message,omitempty"`
	GroupName string `json:"group_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type Response struct {
	Reply string `json:"reply,omitempty"`
}

func HandlerWhats(c echo.Context) error {
	req := new(Request)

	if err := c.Bind(req); err != nil {
		return err
	}

	log.Println(*req)

	return c.JSON(http.StatusCreated, Response{Reply: req.Message})
}

func main() {
	echoHttp := echo.New()

	echoHttp.POST("/message", HandlerWhats)
	echoHttp.Logger.Fatal(echoHttp.Start(":3007"))
}
