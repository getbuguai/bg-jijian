package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

const (
	downloadImageURL = "https://w.wallhaven.cc/full/%s/wallhaven-%s"
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

	imgURLName := img.GetURLName()
	if len(imgURLName) == 0 {
		log.Error().Msg(fmt.Sprintf("DownloadImage GetURLName %v", img))
		return nil
	}
	group := img.GetGrouping()
	if len(group) == 0 {
		log.Error().Msg(fmt.Sprintf("DownloadImage GetGrouping %v", img))
		return nil
	}

	log.Log().Msg(fmt.Sprintf(downloadImageURL, group, imgURLName))

	res, err := http.Get(fmt.Sprintf(downloadImageURL, group, imgURLName))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	img.BodyByte, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return SaveFile(filepath.Join(conf.FileDirectory, img.GetFileName()), img.BodyByte)
}
