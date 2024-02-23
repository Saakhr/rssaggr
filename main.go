package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load()
	portstring := os.Getenv("PORT")
	if portstring == "" {
		portstring = "8000"
	}
	e := echo.New()
	e.Use(middleware.CORS())
	v1 := e.Group("/v1")
	v1.GET("/health", handleHealth)
	v1.GET("/err", handleError)

	fmt.Println("server starting...")
	if err := e.Start(":" + portstring); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
