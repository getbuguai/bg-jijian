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
	ImageType     string
	StartPageNum  uint
	EndPageNum    uint
}

func NewDownloadConfig(downType uint,
	downFile string, downSatrtPageNum uint,
	downAll bool) DownloadConfig {

	res := DownloadConfig{
		FileDirectory: downFile,
		ImageType:     TargetAnime,
		StartPageNum:  downSatrtPageNum,
		EndPageNum:    downSatrtPageNum,
	}
	if downType == 1 {
		res.ImageType = TargetPeople
	}
	if downType == 2 {
		res.ImageType = TargetIndex
	}

	if downAll {
		res.EndPageNum = 0
	}
	return res

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

	// content-length : 2573746 文件的大小
	//fileSize, err := strconv.ParseUint(res.Header.Get("content-length"), 10, 10)
	//if err != nil {
	//	return err
	//}
	//if fileSize == 0 {
	//	return fmt.Errorf("file size is zero ... ")
	//}

	img.BodyByte, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// if fileSize == uint64(len(img.BodyByte)) {
	// 	return fmt.Errorf("file size not equil imgBytes ... ")
	// }

	filename := filepath.Join(conf.FileDirectory, img.GetFileName())
	if len(filename) == 0 {
		return fmt.Errorf("filename length is 0")
	}

	if len(img.BodyByte) == 0 {
		return fmt.Errorf("data length is 0")
	}

	return ioutil.WriteFile(filename, img.BodyByte, 0644)
}
