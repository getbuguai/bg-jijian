package main

import "testing"

func Test_imageMsg_GetGrouping(t *testing.T) {
	img := &imageMsg{
		ID: "392zpd",
	}
	got := img.GetGrouping()
	t.Log(got)

}
