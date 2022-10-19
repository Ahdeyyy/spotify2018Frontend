package db

import (
	"github.com/Ahdeyyy/spotify2018Frontend/models"
	"database/sql"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
)


// open the database
func OpenDatabase() *sql.DB {
	if _, err := os.Stat("songs.db"); err == nil {
		// database exists
		db, err := sql.Open("sqlite3", "./songs.db")
		// if _,err = db.Exec("PRAGMA journal_Mode = WAL"); err != nil {
		// 	fmt.Println(err)
		// }
		// if _,err := db.Exec("PRAGMA busy_timeout = 5000"); err != nil {
		// 	fmt.Println(err)
		// }

		if err != nil {
			log.Println(err)
		}
		return db

	 } else {
		// database does not exist so create a table and return the database
		db, err := sql.Open("sqlite3", "./songs.db")
		if _,err = db.Exec("PRAGMA journal_Mode = WAL"); err != nil {
			log.Println(err)
		}
		if _,err := db.Exec("PRAGMA busy_timeout = 5000"); err != nil {
			log.Println(err)
		}
		if err != nil {
			log.Println(err)
		}
		createTable(db)
		return db
	 }
}

// create table in database
func createTable(db *sql.DB) {

	sqlStmt := `
  CREATE TABLE IF NOT EXISTS songs (
    Id TEXT NOT NULL PRIMARY Key,
    Name TEXT,
    Artists TEXT,
    Danceability FLOAT,
    Energy FLOAT,
    Key INT,
    Loudness FLOAT,
    Mode INT,
    Speechiness FLOAT,
    Acousticness FLOAT,
    Instrumentalness FLOAT,
    Liveness FLOAT,
    Valence FLOAT,
    Tempo FLOAT,
    Duration_ms INT,
    Time_signature INT
  );
  `
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Println(err)
	}
}

// insert data into database
func InsertData(db *sql.DB, songs []models.Song) {
  	for _, song := range songs {

		sqlStmt := `
    INSERT INTO songs(Id, Name, Artists, Danceability, Energy, Key, Loudness, Mode, Speechiness, Acousticness, Instrumentalness, Liveness, Valence, Tempo, Duration_ms, Time_signature) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
		_, err := db.Exec(sqlStmt, song.Id, song.Name, song.Artists, song.Danceability, song.Energy, song.Key, song.Loudness, song.Mode, song.Speechiness, song.Acousticness, song.Instrumentalness, song.Liveness, song.Valence, song.Tempo, song.Duration_ms, song.Time_signature)
		if err != nil {
			log.Println(err)
		}
	}
}

// get all data in database
func FindAll(db *sql.DB) []models.Song {
	rows, err := db.Query("SELECT * FROM songs")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var songs []models.Song
	for rows.Next() {
		var song models.Song
		err := rows.Scan(&song.Id, &song.Name, &song.Artists, &song.Danceability, &song.Energy, &song.Key, &song.Loudness, &song.Mode, &song.Speechiness, &song.Acousticness, &song.Instrumentalness, &song.Liveness, &song.Valence, &song.Tempo, &song.Duration_ms, &song.Time_signature)
		if err != nil {
			log.Println(err)
		}
		songs = append(songs, song)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
	return songs
}

// search database for artist
func FindArtist(db *sql.DB, artist string) []models.Song {
	

	sqlStmt := `
	SELECT * FROM songs WHERE Artists LIKE ?
	`
	rows, err := db.Query(sqlStmt, "%"+artist+"%")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var songs []models.Song
	for rows.Next() {
		var song models.Song
		err := rows.Scan(&song.Id, &song.Name, &song.Artists, &song.Danceability, &song.Energy, &song.Key, &song.Loudness, &song.Mode, &song.Speechiness, &song.Acousticness, &song.Instrumentalness, &song.Liveness, &song.Valence, &song.Tempo, &song.Duration_ms, &song.Time_signature)
		if err != nil {
			log.Println(err)
		}
		songs = append(songs, song)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
	return songs
}
