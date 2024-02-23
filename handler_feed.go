package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Saakhr/rssaggr/internal/database"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (apiCFG *apiconfig) handleCreatefeed(c echo.Context, user database.User) error {
	type parametres struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	param := new(parametres)

	if err := c.Bind(param); err != nil {
		return err
	}
	feed, err := apiCFG.DB.CreateFeed(c.Request().Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      param.Name,
		Url:       param.Url,
		UserID:    user.ID,
	})
	if err != nil {
		er := errResponsed{
			Erorr: fmt.Sprintf("couldn't create a feed :%v", err),
		}
		return c.JSON(http.StatusInternalServerError, er)
	}

	return c.JSON(http.StatusCreated, databaseToFeed(feed))
}
func (apiCFG *apiconfig) handleGetFeeds(c echo.Context) error {

	feeds, err := apiCFG.DB.GetFeeds(c.Request().Context())
	if err != nil {
		er := errResponsed{
			Erorr: fmt.Sprintf("couldn't create user :%v", err),
		}
		return c.JSON(http.StatusInternalServerError, er)
	}

	return c.JSON(http.StatusCreated, databaseToFeeds(feeds))
}
