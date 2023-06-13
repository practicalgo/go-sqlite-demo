package main

import (
	"path"
	"testing"
)

func TestInsertAndQueryData(t *testing.T) {

	dbPath := path.Join(t.TempDir(), "app.db")
	err := initDatabase(dbPath)
	if err != nil {
		t.Fatalf("error initializing database: ", err)
	}

	testData := []Album{
		{"Blue Train", "John Coltrane", 56.99},
		{"Giant Steps", "John Coltrane", 63.99},
		{"Jeru", "Gerry Mulligan", 17.99},
		{"Sarah Vaughan", "Sarah Vaughan", 34.98},
	}

	for _, album := range testData {
		rowCnt, err := insertData(album)
		if rowCnt != 1 {
			t.Fatalf("expected inserted row count: 1, got: %d", rowCnt)
		}
		if err != nil {
			t.Fatalf("error inserting row: %v", album)
		}
	}

	for idx, album := range testData {
		gotAlbum, err := queryData(album.Title)
		if err != nil {
			t.Fatalf("error inserting row: %v", album)
		}

		if gotAlbum.ID != idx {
			t.Errorf("id doesn't match. expected: %s, got: %s", idx, gotAlbum.ID)
		}

		if gotAlbum.Artist != album.Artist {
			t.Errorf("artist doesn't match. expected: %s, got: %s", album.Artist, gotAlbum.Artist)
		}

		if gotAlbum.Price != album.Price {
			t.Errorf("price doesn't match. expected: %s, got: %s", album.Price, gotAlbum.Price)
		}
	}
}
