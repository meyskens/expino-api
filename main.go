package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	createDatabase()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Send your API data!")
	})
	e.POST("/metric", handleAPI)
	e.Logger.Fatal(e.Start(":8080"))
}

func handleAPI(c echo.Context) error {
	input := []APIData{}
	c.Bind(&input)

	if len(input) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": "no data sent"})
	}

	for _, data := range input {
		if len(data.Data) == 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": "no sensor data sent"})
		}
		err := addPoint(data)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
