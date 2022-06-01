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
	artistObj := getArtistObjByName(c.QueryParam("name"))

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
	artistObj := getArtistObjByName(c.QueryParam("name"))

	// check if artist exists in the db
	idExists := database.CheckIfArtistExists(&artistObj)

	if idExists != 0 {
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

func GetAll(c echo.Context) error {
	artists := database.GetAllArtists()

	return c.JSON(200, map[string]interface{}{
		"artists": artists,
	})
}

func Delete(c echo.Context) error {
	artistObj := getArtistObjByName(c.QueryParam("name"))

	// check if artist exists in the db
	idExists := database.CheckIfArtistExists(&artistObj)

	if idExists != 0 {
		// setup id artistObj
		artistObj.Id = idExists

		// delete artist in the db
		deleted := database.DeleteArtist(&artistObj)

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
