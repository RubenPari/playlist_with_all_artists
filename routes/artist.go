package routes

import (
	"github.com/RubenPari/playlist_with_all_artists/controllers"
	"github.com/labstack/echo/v4"
)

func SetUpArtistRoutes(e *echo.Echo) {
	e.POST("/artist/add", controllers.Add)
	e.GET("/artist/all", controllers.GetAll)
	e.GET("/artist/exists", controllers.CheckIfExists)
	e.DELETE("/artist/delete", controllers.Delete)
}
