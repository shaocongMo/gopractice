package main

import "testing"

func TestImageCompress(t *testing.T) {
	imageCompress("404-bg.jpg", 1900, 870)
	imageCompress("nn", 1900, 870)
	imageCompress("tt.png", 1900, 870)
}
