package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"freefolks-fc/config"
	"freefolks-fc/db/gorm"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Read()
	port := os.Getenv("APP_PORT")
	fmt.Println(port)
	server := &http.Server{
		Addr:         port,
		ReadTimeout:  15 * time.Minute,
		WriteTimeout: 15 * time.Minute,
	}
	e := echo.New()

	_ = gorm.ConnectPostgreSQL()

	if err := e.StartServer(server); err != nil {
		e.Logger.Fatal("Shutting down the server")
	}
}
