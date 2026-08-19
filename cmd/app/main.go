package main

import (
	"finance-planning-go/internal/app/config"
	"finance-planning-go/internal/app/database"
	"finance-planning-go/internal/app/middleware"
	"finance-planning-go/internal/app/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()
	config.ENV{}.Load()
	database.Database{DBType: "MYSQL"}.Start()
	app.Use(middleware.CORS{}.Use())
	routes.Init(app)

	fmt.Println("== API STARTED ==")
	app.Run(fmt.Sprintf("localhost:%s", config.ENV{}.Get("API_PORT")))
}
