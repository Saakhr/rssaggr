package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Saakhr/rssaggr/internal/database"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (apiCFG *apiconfig) handleCreateUser(c echo.Context) error {
	type parametres struct {
		Name string `json:"name"`
	}
	param := new(parametres)

	if err := c.Bind(param); err != nil {
		return err
	}
	user, err := apiCFG.DB.CreateUser(c.Request().Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      param.Name,
	})
	if err != nil {
		er := errResponsed{
			Erorr: fmt.Sprintf("couldn't create user :%v", err),
		}
		return c.JSON(http.StatusInternalServerError, er)
	}

	return c.JSON(http.StatusCreated, databaseToUser(user))
}
func (apiCFG *apiconfig) handleGetUserByApiKey(c echo.Context, user database.User) error {
	return c.JSON(http.StatusOK, databaseToUser(user))
}

func (apiCFG *apiconfig) handleGetPostsForUser(c echo.Context, user database.User) error {
	posts, err := apiCFG.DB.GetPostsForUser(c.Request().Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		return handleAnError(c, fmt.Sprintf("couldn't get posts: %v", err))
	}
	return c.JSON(http.StatusOK, databaseToPosts(posts))
}
