package controllers

import (
	"github.com/RubenPari/playlist_with_all_artists/database"
	"github.com/RubenPari/playlist_with_all_artists/models"
	"github.com/RubenPari/playlist_with_all_artists/utils"
	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify/v2"
)

// GetArtists
// this function get in input the artist name,
// search it on spotify and return the artist object model
func getArtistObjByName(name string) models.Artist {

	client, ctx := utils.GetSpotifyClientCtx()

	// search artist in spotify
	artistSpotify, _ := client.Search(ctx, name, spotify.SearchTypeArtist)

	artist := models.Artist{
		Name:      artistSpotify.Artists.Artists[0].Name,
		SpotifyId: artistSpotify.Artists.Artists[0].ID,
	}

	return artist
}

func Add(c echo.Context) error {
	artistObj := getArtistObjByName(c.Param("name"))

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

func CheckIfExists(c echo.Context) error {
	artistObj := getArtistObjByName(c.Param("name"))

	// check if artist exists in the db
	exists := database.CheckIfArtistExists(artistObj)

	if exists {
		return c.JSON(200, map[string]string{
			"status":  "ok",
			"message": "artist exists",
		})
	} else {
		return c.JSON(404, map[string]string{
			"status":  "error",
			"message": "artist not found",
		})
	}
}

func Delete(c echo.Context) error {
	artistObj := getArtistObjByName(c.Param("name"))

	// check if artist exists in the db
	exists := database.CheckIfArtistExists(artistObj)

	if exists {
		// delete artist in the db
		deleted := database.DeleteArtist(artistObj)

		if deleted {
			return c.JSON(200, map[string]string{
				"status":  "ok",
				"message": "artist deleted",
			})
		} else {
			return c.JSON(500, map[string]string{
				"status":  "error",
				"message": "error to delete in the db",
			})
		}
	} else {
		return c.JSON(404, map[string]string{
			"status":  "error",
			"message": "artist not found",
		})
	}
}
