---
title: Using SQLite from Go
date: 2023-06-29
categories:
-  articles
---

In the book, we used MySQL server as a way to store relational data from our applications. In this post, we will see how to use work with SQLite from Go. 

The table names and operations are intentionally chosen to losely match that of the official Go Project tutorial, [Accessing a relational database](https://go.dev/doc/tutorial/database-access) which uses MySQL as the database server.

Let's get started!

- [Prerequisites](#prerequisites)
- [Overview](#overview)
- [Initialize your module](#initialize-your-module)
- [Initializing the database](#initializing-the-database)
- [Types for Album](#types-for-album)
- [Inserting data](#inserting-data)
- [Querying multiple rows](#querying-multiple-rows)
- [Query for a single row](#query-for-a-single-row)
- [Using sqlite3 to interact with the database](#using-sqlite3-to-interact-with-the-database)
- [Using in-memory databases for testing](#using-in-memory-databases-for-testing)
- [Conclusion](#conclusion)
- [Learn more](#learn-more)

## Prerequisites

- An installation of Go
- A tool to edit your code. Any text editor you have will work fine
- A command terminal to run the `go` commands from
- (optional) a `sqlite3` install, see [here](https://sqlite.org/cli.html) for instructions.


## Overview

[SQLite](https://www.sqlite.org/about.html) is file-based database engine.

Thus, we do not have a separate server process for it.

We will create the database, create a table, insert rows and query them back
using:

- [database/sql](https://pkg.go.dev/database/sql#section-documentation)
- [pkg.go.dev/modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite): I choose
  this package because it's a cgo-free Go implementation of SQLite. Thus, installing
  it doesn't require you to have a working C compiler, to start with. What's interesting
  is this bit of text from the [sqlite](https://sqlite.org/index.html) website:
  > SQLite is a C-language library ...

  Thus, `modernc.org/sqlite` is a Go implementation of SQLite.

## Initialize your module

- Create a directory for the go module
- Create a go module inside it using `go mod init github.com/practicalgo/go-sqlite-demo`
- Fetch the sqlite module, `go get modernc.org/sqlite`

## Initializing the database

To initialize the database, i.e. to create the file which will be used for the database,
import the two libraries:

- `database/sql`
- unnamed import of `modernc.org/sqlite` to register the `sqlite` SQL [driver](https://pkg.go.dev/database/sql/driver)
- call the `sql.Open()` function with specifying `sqlite` as the driver (the first argument) and the file
  path as the second argument

Once we have initialized the database, we can create a table, `album` using the
[`ExecContext()`](https://pkg.go.dev/database/sql#DB.ExecContext) method:

This is a code snippet from `app.go` which encapsulates the initialization in a function, `initDatabase()`:

```go

import (
	"database/sql"
	_ "modernc.org/sqlite"
)


var db *sql.DB

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

## Types for Album

We have defined two types for each album. The first, `Album` is to represent an album we are going to add
to the table:

```go
type Album struct {
	Title  string
	Artist string
	Price  float32
}
```

The second type, `AlbumDbRow` is used to represent an album that we *retrieve* from the database:

```go
type AlbumDbRow struct {
	ID int
	Album
}
```

An album retrieved from the database, in addition to all the fields of `Album` will have an additional field, `ID` representing their identifier (a row number) in the album. Thus, we embed the `Album` 
struct and define `ID` as an additional field. 

## Inserting data

Once we have created the table, we can insert data using the `INSERT` SQL statement. 

To execute the SQL statement, we will once again use the `db.ExecContext()` method.

The `addAlbum()` function in `app.go` shows how we can do so:

```go
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
```

## Querying multiple rows

To query data, we will execute the `SELECT` statement. 

As we expect to retrieve multiple rows, we use the [`db.QueryContext()`](https://pkg.go.dev/database/sql#DB.QueryContext) method. 

The `albumsByArtist()` function shows an example:

```go
func albumsByArtist(artist string) ([]AlbumDbRow, error) {

        // this slice will contain all the albums retrieved

	var albums []AlbumDbRow
	rows, err := db.QueryContext(
		context.Background(),
		`SELECT * FROM album WHERE artist=?;`, artist,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// we iterate over each row retrieved, scanning each row
	// into an object of type album, successively appending each
	// scanned album into a slice, albums[]
	
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


```
SQLITE_DB_PATH=app.db ./go-sqlite-demo
2023/06/27 13:13:54 recreating table: album
database initialized..
test data inserted..
querying test data by album ID..
{1 {Giant Steps John Coltrane 63.99}}
{2 {Jeru Gerry Mulligan 17.99}}
{3 {Sarah Vaughan Sarah Vaughan 34.98}}
{4 {Blue Train John Coltrane 56.99}}
```

## Using sqlite3 to interact with the database 

```sql
$ sqlite3 app.db
SQLite version 3.42.0 2023-05-16 12:36:15
Enter ".help" for usage hints.
sqlite> .tables
album
sqlite> .database
main: /home/echorand/work/github.com/practicalgo/go-sqlite-demo/app.db r/w
sqlite> .schema album
CREATE TABLE album (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        title TEXT NOT NULL,
                        artist TEXT NOT NULL,
                        price REAL NOT NULL
                );
sqlite> select * from album
   ...> ;
1|Giant Steps|John Coltrane|63.9900016784668
2|Jeru|Gerry Mulligan|17.9899997711182
3|Sarah Vaughan|Sarah Vaughan|34.9799995422363
4|Blue Train|John Coltrane|56.9900016784668

```

## Using in-memory databases for testing

## Conclusion

## Learn more

- [Embedding struct in structs](https://eli.thegreenplace.net/2020/embedding-in-go-part-1-structs-in-structs/)



