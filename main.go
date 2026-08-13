package main

import (
	"finance-planning-go/config/env_cfg"
	"finance-planning-go/database/mysql_db"
	"finance-planning-go/middleware/cors_middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	env_cfg.Load()
	app := gin.Default()
	defer func() {
		fmt.Println("== API STARTED ==")
		app.Run(fmt.Sprintf("localhost:%s", env_cfg.Get("API_PORT")))
	}()

	mysql_db.Start()
	app.Use(cors_middleware.UseCORS())
}
