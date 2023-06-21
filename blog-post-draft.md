## Prerequisites

- An installation of Go
- A tool to edit your code. Any text editor you have will work fine
- A command terminal. Go works well using any terminal on Linux and Mac, and on PowerShell or cmd in Windows


## Overview

[SQLite](https://www.sqlite.org/about.html) is file-based database engine.

Thus, we do not have a separate server process for it.

We will create the database, create a table, insert rows and query them back
using:

- [database/sql](https://pkg.go.dev/database/sql#section-documentation)
- [pkg.go.dev/modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite): I choose
  this package because it's a cgo-free Go implementation of SQLite


## Initialize your module

- Create a directory for the go module
- Create a go module inside it using `go mod init`
- Fetch the sqlite module, `go get modernc.org/sqlite`

## Initializing the database

```
import (
	"database/sql"
	_ "modernc.org/sqlite"
)


func initDatabase(dbPath string) error {
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
...
```


## Create table

```
func initDatabase(dbPath string) error {
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	_, err = db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS album (
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
```




## Inserting data 

```func addAlbum(a *Album) (int64, error) {
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
```

## Querying multiple rows

```
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
```

## Query for a single row 



```
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
```


## Using sqlite3 to interact with the database 

## Conclusion


