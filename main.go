package main

import (
	"github.com/RubenPari/playlist_with_all_artists/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	var e = echo.New()

	routes.SetUpPlaylistRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
