package main

import (
	"context"
	"testing"
)

func Test_imageMsg_GetGrouping(t *testing.T) {
	img := &imageMsg{
		ID: "392zpd",
	}
	got := img.GetGrouping()
	t.Log(got)

}

func TestGetJsonReq_GetJson(t *testing.T) {

	r := &GetJsonReq{
		Target:  TargetAnime,
		PageNum: 2,
	}
	if err := r.GetJson(context.Background(), func(j *ResultJSON) error {
		t.Log(j.Result.Current, j.Result.Records[5].ID)
		return nil
	}); err != nil {
		t.Errorf("GetJson() error = %v", err)
	}

}
