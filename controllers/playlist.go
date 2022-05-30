package controllers

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"os"
)

func UpdatePlaylist(c echo.Context) error {
	_ = godotenv.Load()

	userId := os.Getenv("USER_ID")

	ctx := context.Background()

	config := &clientcredentials.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		TokenURL:     "https://accounts.spotify.com/api/token",
	}

	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)

	client := spotify.New(httpClient)

	playlists, errPlaylists := client.GetPlaylistsForUser(ctx, userId)
	if errPlaylists != nil {
		log.Fatalf("couldn't get playlists: %v", errPlaylists)
	}

	return c.JSON(200, map[string]interface{}{
		"status":    "success",
		"playlists": playlists,
	})
}
