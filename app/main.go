package main

import "github.com/labstack/echo"

func main() {
	eo := echo.New()
	err := eo.Start(":8080")
	if err != nil {
		return
	}
}
