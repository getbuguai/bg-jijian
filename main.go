package main

import (
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

var downloadConf DownloadConfig

func init() {
	downloadConf = DownloadConfig{FileDirectory: "./images"}
}

func main() {
	ctx := context.Background()
	rt := &GetJsonReq{
		// Target:  TargetAnime,
		Target:  TargetPeople,
		PageNum: 1,
	}
	err := rt.GetJson(ctx, DefaultDisposeImageJSON)
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Log().Msgf("Page: %d", rt.PageNum)

	for rt.HaveNextPage() {
		err = rt.GetJson(ctx, DefaultDisposeImageJSON)
		if err != nil {
			log.Fatal().Err(err)
		}
	}
}

// DefaultDisposeImageJSON 默认处理 GetJson
func DefaultDisposeImageJSON(body *ResultJSON) error {

	wg := sync.WaitGroup{}
	wg.Add(len(body.Result.Records))

	for _, v := range body.Result.Records {
		log.Log().Msgf("URLName: %s", v.GetURLName())

		go func(img *imageMsg) {
			defer wg.Done()
			downloadConf.DownloadImage(img)
		}(v)

		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()

	return nil
}
