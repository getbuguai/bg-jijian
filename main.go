package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var downloadConf DownloadConfig

// flat 的值
var downType uint
var downFile string
var downStartPageNum uint
var downAll bool
var showLog bool
var showHelp bool

func init() {
	flag.UintVar(&downType, "t", 0, "下载的图片的类型[0 : 二次元, 1 : 人物, 2 : 精选 ]")
	flag.StringVar(&downFile, "o", "./images", "下载的图片保存路径")
	flag.UintVar(&downStartPageNum, "s", 1, "下载第几页的图片")
	flag.BoolVar(&downAll, "a", false, "是否下载后续的所有页面, 默认只下载一页")
	flag.BoolVar(&showLog, "v", false, "是否显示日志信息, 默认不显示")
	flag.BoolVar(&showHelp, "h", false, "帮助信息")

	flag.Usage = func() {
		fmt.Printf("\n 可执行文件路径: %s: \n", os.Args[0])
		flag.PrintDefaults()
		fmt.Printf("作者 : 这个程序不太乖 ,\n " +
			"GitHub: https://github.com/getbuguai \n " +
			"BiliBili: https://space.bilibili.com/278413353 \n ")
	}
}

func main() {
	flag.Parse()

	fmt.Printf("-t %d, -o %s, -s %d , -a %t, -v %t", downType, downFile, downStartPageNum, downAll, showLog)
	if len(os.Args) == 1 || showHelp {
		flag.Usage()
		return
	}

	downloadConf = NewDownloadConfig(downType, downFile, downStartPageNum, downAll)

	// 日志
	if showLog {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	ctx := context.Background()
	rt := &GetJsonReq{
		Target:  downloadConf.ImageType,
		PageNum: downloadConf.StartPageNum,
		EndNum:  downloadConf.EndPageNum,
	}

	err := rt.GetJson(ctx, DefaultDisposeImageJSON)
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Info().Msgf("All Page: %d,All Images: %d, Page Size:%d", rt.TempDate.Pages, rt.TempDate.Total, rt.TempDate.Size)

	for rt.HaveNextPage() {
		err = rt.GetJson(ctx, DefaultDisposeImageJSON)
		if err != nil {
			log.Fatal().Err(err)
		}
	}
}

// DefaultDisposeImageJSON 默认处理 GetJson
func DefaultDisposeImageJSON(body *ResultJSON) (ok bool, err error) {

	wg := sync.WaitGroup{}
	wg.Add(len(body.Result.Records))

	for _, v := range body.Result.Records {
		log.Debug().Msgf("URLName: %s", v.GetURLName())

		go func(img *imageMsg) {
			defer wg.Done()
			err = downloadConf.DownloadImage(img)
			if err != nil {
				log.Error().Err(err).Msgf("URLName: %s", img.ID)
			}
		}(v)

		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()

	return true, nil
}
