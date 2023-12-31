package main

import (
	"reflect"
	"testing"
)

func TestInsertAndQueryData(t *testing.T) {
	err := initDatabase(":memory:")
	if err != nil {
		t.Fatal("error initializing database: ", err)
	}

	insertErr := insertTestData()
	if insertErr != nil {
		t.Fatal(insertErr)
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
