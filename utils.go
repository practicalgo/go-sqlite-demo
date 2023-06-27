package main

import "fmt"

func insertTestData() error {
	testData := []Album{
		{"Giant Steps", "John Coltrane", 63.99},
		{"Jeru", "Gerry Mulligan", 17.99},
		{"Sarah Vaughan", "Sarah Vaughan", 34.98},
		{"Blue Train", "John Coltrane", 56.99},
	}

	for _, album := range testData {
		_, err := addAlbum(&album)
		if err != nil {
			return fmt.Errorf("error inserting row: %v", err)
		}
	}
	return nil

}
