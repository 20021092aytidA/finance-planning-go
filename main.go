package main

import (
	"finance-planning-go/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	defer func() {
		fmt.Println("== API STOPPED ==")
	}()
	app.Use(middleware.UseCORS())
}
