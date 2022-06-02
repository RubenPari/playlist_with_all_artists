package routes

import (
	"github.com/RubenPari/playlist_with_all_artists/controllers"
	"github.com/labstack/echo/v4"
)

func SetUpUserRoutes(e *echo.Echo) {
	e.GET("/user/getAllSongs", controllers.GetAllSongs)
}
