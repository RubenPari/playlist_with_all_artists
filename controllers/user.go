package controllers

import (
	"github.com/RubenPari/playlist_with_all_artists/utils"
	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify/v2"
)

func GetAllSongs(c echo.Context) error {
	client, ctx := utils.GetSpotifyClientCtx()

	// Get all songs from the user
	songs, _ := client.AddTracksToPlaylist(ctx, spotify.Offset())

	return c.JSON(200, songs)
}
