package controllers

import (
	"github.com/RubenPari/playlist_with_all_artists/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify/v2"
	"log"
	"os"
)

func UpdatePlaylist(c echo.Context) error {
	_ = godotenv.Load()

	userId := os.Getenv("USER_ID")

	client, ctx := utils.GetSpotifyClientCtx()

	// get all playlist of the user logged
	playlists, errPlaylists := client.GetPlaylistsForUser(ctx, userId)
	if errPlaylists != nil {
		log.Fatalf("couldn't get playlists: %v", errPlaylists)
	}

	// get specific playlist id
	var idPlaylist spotify.ID

	for i := 0; i < len(playlists.Playlists); i++ {
		if playlists.Playlists[i].Name == os.Getenv("PLAYLIST_NAME") {
			idPlaylist = playlists.Playlists[i].ID
			break
		}
	}

	// TODO: continue... get all tracks and add the to playlist

	return c.JSON(200, idPlaylist)
}

// GetAllTracks
// Get all tracks for every artist saved in the db
func GetAllTracks(idPlaylist spotify.ID) {
	//client, ctx := utils.GetSpotifyClientCtx()
	//db := database.GetDatabase()
}
