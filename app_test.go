package main

import (
	"path"
	"reflect"
	"testing"
)

func TestInsertAndQueryData(t *testing.T) {

	dbPath := path.Join(t.TempDir(), "app.db")
	err := initDatabase(dbPath)
	if err != nil {
		t.Fatal("error initializing database: ", err)
	}

	testData := []Album{
		{"Giant Steps", "John Coltrane", 63.99},
		{"Jeru", "Gerry Mulligan", 17.99},
		{"Sarah Vaughan", "Sarah Vaughan", 34.98},
		{"Blue Train", "John Coltrane", 56.99},
	}

	for _, album := range testData {
		rowCnt, err := addAlbum(&album)
		if rowCnt != 1 {
			t.Fatalf("expected inserted row count: 1, got: %d", rowCnt)
		}
		if err != nil {
			t.Fatalf("error inserting row: %v", err)
		}
	}

	expectedAlbums := []AlbumDbRow{
		{4, Album{"Blue Train", "John Coltrane", 56.99}},
		{1, Album{"Giant Steps", "John Coltrane", 63.99}},
	}
	gotAlbums, err := albumsByArtist("John Coltrane")
	if err != nil {
		t.Fatalf("error querying data: %v", err)
	}

	if !reflect.DeepEqual(gotAlbums, expectedAlbums) {
		t.Fatalf("expected: %#v, got: %#v", expectedAlbums, gotAlbums)
	}

	expectedAlbum := AlbumDbRow{
		3, Album{"Sarah Vaughan", "Sarah Vaughan", 34.98},
	}
	gotAlbum, err := albumByID(expectedAlbum.ID)
	if err != nil {
		t.Fatal("expected non-nil error, got:", err)
	}
	if !reflect.DeepEqual(expectedAlbum, gotAlbum) {
		t.Fatalf("expected: %#v, got: %#v", expectedAlbum, gotAlbum)
	}
}
