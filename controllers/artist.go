package controllers

import (
	"github.com/RubenPari/playlist_with_all_artists/database"
	"github.com/RubenPari/playlist_with_all_artists/models"
	"github.com/RubenPari/playlist_with_all_artists/utils"
	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify/v2"
)

func Add(c echo.Context) error {
	client, ctx := utils.GetSpotifyClientCtx()

	// get name
	name := c.FormValue("name")

	// search artist in spotify
	artistSpotify, err := client.Search(ctx, name, spotify.SearchTypeArtist)
	if err != nil {
		return c.JSON(500, err)
	}

	// create artist object
	var artistObj = models.Artist{
		Name:      artistSpotify.Artists.Artists[0].Name,
		SpotifyId: artistSpotify.Artists.Artists[0].ID,
	}

	// insert artist in the db
	inserted := database.InsertArtist(artistObj)

	if inserted {
		return c.JSON(200, map[string]string{
			"status":  "ok",
			"message": "inserted successfully",
		})
	} else {
		return c.JSON(500, map[string]string{
			"status":  "error",
			"message": "error to insert in the db",
		})
	}
}
