package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

type Album struct {
	Title  string
	Artist string
	Price  float32
}

type AlbumDbRow struct {
	ID int
	Album
}

var db *sql.DB

func initDatabase(dbPath string) error {
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	_, err = db.ExecContext(
		context.Background(),
		`DROP TABLE IF EXISTS album;
		 CREATE TABLE album (
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			title TEXT NOT NULL, 
			artist TEXT NOT NULL, 
			price REAL NOT NULL
		)`,
	)
	if err != nil {
		return err
	}
	return nil
}

func addAlbum(a *Album) (int64, error) {
	result, err := db.ExecContext(
		context.Background(),
		`INSERT INTO album (title, artist, price) VALUES (?,?,?);`, a.Title, a.Artist, a.Price,
	)
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func albumsByArtist(artist string) ([]AlbumDbRow, error) {

	var albums []AlbumDbRow
	rows, err := db.QueryContext(
		context.Background(),
		`SELECT * FROM album WHERE artist=?;`, artist,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var album AlbumDbRow
		if err := rows.Scan(
			&album.ID, &album.Title, &album.Artist, &album.Price,
		); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, err
}

func albumByID(id int) (AlbumDbRow, error) {
	var album AlbumDbRow
	row := db.QueryRowContext(
		context.Background(),
		`SELECT * FROM album WHERE id=?`, id,
	)
	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		return album, err
	}
	return album, nil
}

func main() {

	dbPath := os.Getenv("SQLITE_DB_PATH")
	if len(dbPath) == 0 {
		log.Fatal("specify the SQLITE_DB_PATH environment variable")
	}
	err := initDatabase(dbPath)
	if err != nil {
		log.Fatal("error initializing DB connection: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("error initializing DB connection: ping error: ", err)
	}
	fmt.Println("database initialized..")
	err = insertTestData()
	if err != nil {
		log.Fatal("error inserting test data: ", err)
	}
	fmt.Println("test data inserted..")

	fmt.Println("querying test data by album ID..")
	// query back each record with IDs 1 - 4
	for i := 1; i <= 4; i++ {
		album, err := albumByID(i)
		if err != nil {
			fmt.Printf("error querying album ID: %d, %s\n", i, err)
		} else {
			fmt.Printf("%v\n", album)
		}
	}
}
