package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

const (
	downloadImageURL = "https://w.wallhaven.cc/full/q6/wallhaven-"
)

// DownloadConfig 下载的配置文件
type DownloadConfig struct {
	FileDirectory string
}

// DownloadImage 下载图片
func (conf *DownloadConfig) DownloadImage(img *imageMsg) error {
	if conf == nil {
		return fmt.Errorf("file directory is nil")
	}

	imgName := img.GetFileName()
	if len(imgName) == 0 {
		return nil
	}

	log.Log().Msg(fmt.Sprintf("%s%s", downloadImageURL, imgName))

	res, err := http.Get(fmt.Sprintf("%s%s", downloadImageURL, imgName))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	img.BodyByte, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return SaveFile(filepath.Join(conf.FileDirectory, imgName), img.BodyByte)
}
