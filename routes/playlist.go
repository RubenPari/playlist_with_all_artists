package routes

import (
	"github.com/RubenPari/playlist_with_all_artists/controllers"
	"github.com/labstack/echo/v4"
)

func SetUpPlaylistRoutes(e *echo.Echo) {
	e.POST("update-playlist", controllers.UpdatePlaylist)
}
