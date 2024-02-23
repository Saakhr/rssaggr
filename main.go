package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Saakhr/rssaggr/internal/database"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type apiconfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	portstring := os.Getenv("PORT")
	if portstring == "" {
		portstring = "8000"
	}

	dburl := os.Getenv("DB_URL")
	if dburl == "" {
		log.Fatal("no db url in env")
	}
	con, err := sql.Open("postgres", dburl)
	if err != nil {
		log.Fatal("Couldn't connect to db", err)
	}

	queries := database.New(con)

	apiCFG := apiconfig{
		DB: queries,
	}

	e := echo.New()
	e.Use(middleware.CORS())
	v1 := e.Group("/v1")
	v1.GET("/health", handleHealth)
	v1.GET("/err", handleError)
	v1.POST("/users", apiCFG.handleCreateUser)
	v1.GET("/users", apiCFG.middlewareAuth(apiCFG.handleGetUserByApiKey))
	v1.POST("/feeds", apiCFG.middlewareAuth(apiCFG.handleCreatefeed))
	v1.GET("/feeds", apiCFG.handleGetFeeds)
	v1.POST("/feed_follow", apiCFG.middlewareAuth(apiCFG.handleCreateFeedFollow))
	v1.GET("/feed_follow", apiCFG.middlewareAuth(apiCFG.handleGetFeedsFollows))
	v1.DELETE("/feed_follow/:id", apiCFG.middlewareAuth(apiCFG.handleUnfollowFeed))
	v1.GET("/posts", apiCFG.middlewareAuth(apiCFG.handleGetPostsForUser))

	fmt.Println("server starting...")
	go startScraping(apiCFG.DB, 10, time.Minute)
	if err := e.Start(":" + portstring); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
