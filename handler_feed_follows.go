package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Saakhr/rssaggr/internal/database"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (apiCFG *apiconfig) handleCreateFeedFollow(c echo.Context, user database.User) error {
	type parametres struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	param := new(parametres)

	if err := c.Bind(param); err != nil {
		return err
	}
	feedfollow, err := apiCFG.DB.CreateFeedFollows(c.Request().Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    param.FeedID,
	})
	if err != nil {
		er := errResponsed{
			Erorr: fmt.Sprintf("couldn't create a feed follow :%v", err),
		}
		return c.JSON(http.StatusInternalServerError, er)
	}

	return c.JSON(http.StatusCreated, databaseToFeedFollow(feedfollow))
}
func (apiCFG *apiconfig) handleGetFeedsFollows(c echo.Context, user database.User) error {

	feedfollows, err := apiCFG.DB.GetFeedFollows(c.Request().Context(), user.ID)
	if err != nil {
		er := errResponsed{
			Erorr: fmt.Sprintf("couldn't get feed follows :%v", err),
		}
		return c.JSON(http.StatusInternalServerError, er)
	}

	return c.JSON(http.StatusOK, databaseToFeedFollows(feedfollows))
}
func (apiCFG *apiconfig) handleUnfollowFeed(c echo.Context, user database.User) error {

	para := c.Param("id")
	param, err := uuid.Parse(para)
	if err != nil {
		er := errResponsed{
			Erorr: fmt.Sprintf("couldn't parse follow feed id :%v", err),
		}
		return c.JSON(http.StatusInternalServerError, er)
	}
	ad, err := apiCFG.DB.UnfollowFeed(c.Request().Context(), database.UnfollowFeedParams{
		ID:     param,
		UserID: user.ID,
	})
	if err != nil || ad == 0 {
		return handleAnError(c, fmt.Sprintf("Couldn't delete feed follow: %v", err))
	}
	return handleAMessage(c, "Feed Unfollowed", http.StatusOK)
}
