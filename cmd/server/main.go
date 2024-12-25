package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"scheduler-server/internal"
	"scheduler-server/internal/database"
	"scheduler-server/internal/handler"
)

func main() {
	r := gin.Default()

	// middleware

	// DB
	db := database.NewMongoDb()

	err := db.Open()
	if err != nil {
		log.Fatalf("[main] not able to connect to database err: %+v", err)
	}

	defer db.Close()
	// handlers
	ginHandler := handler.NewHandler(db)

	r.GET("/health", ginHandler.Health)

	port := internal.GetConfigString("server.port")
	fmt.Println("[main] starting http server on port:", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("[main] failed to start server: %v", err)
	}
}
