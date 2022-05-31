package models

type Artist struct {
	Id        int    `json:"id"`
	SpotifyId string `json:"spotify_id"`
	Name      string `json:"name"`
}
