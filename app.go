package main

import (
	"context"
	"database/sql"
	"fmt"

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
		`CREATE TABLE IF NOT EXISTS album (id INTEGER, title TEXT NOT NULL, artist TEXT NOT NULL, price REAL NOT NULL)`,
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
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func albumsByArtist(artist string) ([]*AlbumDbRow, error) {
	result, err := db.QueryContext(
		context.Background(),
		`INSERT INTO album (title, artist, price) VALUES (?,?,?);`, a.Title, a.Artist, a.Price,
	)
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil

	return nil, nil
}

func albumByID(id int) (*AlbumDbRow, error) {
	return nil, nil
}

func main() {
	fmt.Println("HELLO")
}
