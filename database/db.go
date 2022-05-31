package database

import (
	"database/sql"
	"github.com/RubenPari/playlist_with_all_artists/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/zmb3/spotify/v2"
)

func GetDatabase() *sql.DB {
	_ = godotenv.Load()

	//host := os.Getenv("DB_HOST")
	//port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	//user := os.Getenv("DB_USER")
	//password := os.Getenv("DB_PASSWORD")
	//dbname := os.Getenv("DB_NAME")

	// connection string
	//pSqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	pSqlConn := "postgresql://doadmin:AVNS_ZfCTZ4BWfndxI5L@playlist-do-user-10671938-0.b.db.ondigitalocean.com:25060/playlist?sslmode=require"
	db, err := sql.Open("postgres", pSqlConn)
	if err != nil {
		panic(err)
	}

	return db
}

func GetAllArtists() []models.Artist {
	db := GetDatabase()
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	rows, err := db.Query("SELECT * FROM artists")
	if err != nil {
		panic(err)
	}

	var artists []models.Artist

	for rows.Next() {
		var id int
		var spotifyId spotify.ID
		var name string

		err = rows.Scan(&id, &spotifyId, &name)
		if err != nil {
			panic(err)
		}

		artists = append(artists, models.Artist{
			Id:        id,
			SpotifyId: spotifyId,
			Name:      name,
		})
	}

	return artists
}

func InsertArtist(artist models.Artist) bool {
	db := GetDatabase()
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	res, err := db.Exec("INSERT INTO artists (spotify_id, name) VALUES ($1, $2)", artist.SpotifyId, artist.Name)
	if err != nil {
		panic(err)
	}
	if res != nil {
		return true
	} else {
		return false
	}
}

func DeleteArtist(artist models.Artist) {
	db := GetDatabase()
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	_, err := db.Exec("DELETE FROM artists WHERE id = $1", artist.Id)
	if err != nil {
		panic(err)
	}
}

func CheckIfArtistExists(artist models.Artist) bool {
	db := GetDatabase()
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	var id int
	err := db.QueryRow("SELECT id FROM artists WHERE spotify_id = $1", artist.SpotifyId).Scan(&id)
	if err != nil {
		return false
	}

	return true
}
