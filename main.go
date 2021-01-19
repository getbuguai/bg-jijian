package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"time"
)

var downloadConf DownloadConfig

func init() {
	downloadConf = DownloadConfig{FileDirectory: "./images"}
}

func main() {
	ctx := context.Background()
	rt := &GetJsonReq{
		Target:  TargetAnime,
		PageNum: 1,
	}
	err:=rt.GetJson(ctx, DefaultDisposeImageJSON)
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Log().Msgf("Page: %d", rt.PageNum)

	if rt.HaveNextPage() {
		err = rt.GetJson(ctx, DefaultDisposeImageJSON)
		if err != nil {
			log.Fatal().Err(err)
		}
	}
}

// DefaultDisposeImageJSON 默认处理 GetJson
func DefaultDisposeImageJSON(body *ResultJSON) error {

	for _, v := range body.Result.Records {
		log.Log().Msgf("FileNam: %s", v.GetFileName())

		go downloadConf.DownloadImage(v)
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}
