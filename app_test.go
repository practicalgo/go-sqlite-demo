package main

import (
	"path"
	"reflect"
	"testing"
)

func TestInsertAndQueryData(t *testing.T) {

	t.Cleanup(func() {
		db.Close()
	})

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

	for idx, album := range testData {
		lastInsertId, err := addAlbum(&album)
		if lastInsertId != int64(idx+1) {
			t.Fatalf("expected inserted ID: %d, got: %d", idx+1, lastInsertId)
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

	if len(gotAlbums) != len(expectedAlbums) {
		t.Fatalf("expected to get: %d albums as result, got: %d", len(expectedAlbums), len(gotAlbums))
	}

	gotAlbumMap := make(map[int]AlbumDbRow)
	for _, a := range gotAlbums {
		gotAlbumMap[a.ID] = a
	}

	for _, a := range expectedAlbums {
		if _, ok := gotAlbumMap[a.ID]; !ok {
			t.Error("expected to get album with ID:", a.ID)
		}
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
