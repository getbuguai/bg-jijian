package main

import "testing"

func TestDownloadConfig_DownloadImage(t *testing.T) {
	downloadConf := DownloadConfig{FileDirectory: "./images"}

	if err := downloadConf.DownloadImage(&imageMsg{
		Type: "j",
		ID:   "q6oxqq",
		X:    0,
		Y:    0,
	}); err != nil {
		t.Errorf("DownloadImage() error = %v", err)
	}
}
