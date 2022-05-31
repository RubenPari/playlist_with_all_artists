package models

import "github.com/zmb3/spotify/v2"

type Artist struct {
	Id        int        `json:"id"`
	SpotifyId spotify.ID `json:"spotify_id"`
	Name      string     `json:"name"`
}
