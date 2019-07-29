package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
	"strconv"

	"github.com/nfnt/resize"
)

func imageCompress(path string, width, height uint) {
	fmt.Println("start image compress:", path)
	filePathWithoutSuffix, imageType, err := isPictureFormat(path)
	if err != nil {
		log.Println("Err: ", err)
		return
	}
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}
	var img image.Image

	if imageType == "jpeg" || imageType == "jpg" {
		img, err = jpeg.Decode(file)
	} else if imageType == "png" {
		img, err = png.Decode(file)
	}
	if err != nil {
		log.Println(err)
		return
	}

	// 压缩
	newImg := resize.Resize(width, height, img, resize.Lanczos3)

	// 创建压缩图片
	out, err := os.Create(filePathWithoutSuffix + "_resize" + "." + imageType)
	if err != nil {
		log.Println(err)
		return
	}
	defer out.Close()
	if imageType == "jpeg" || imageType == "jpg" {
		jpeg.Encode(out, newImg, nil)
	} else {
		png.Encode(out, newImg)
	}
}

func isPictureFormat(filePath string) (string, string, error) {
	splits := strings.Split(filePath, ".")
	if len(splits) <= 1 {
		return "", "", errors.New("without suffix")
	}
	suffix := splits[1]
	if suffix == "jpg" || suffix == "jpeg" || suffix == "png" {
		return splits[0], splits[1], nil
	}
	return "", "", errors.New("It's not a jpg, jpeg or png picture")
}

func main() {
	argsLen := len(os.Args)
	filePath := ""
	var width uint64 = 1900
	var height uint64 = 870
	if argsLen >= 2 {
		filePath = os.Args[1]
	}
	if argsLen >= 4 {
		var err error
		width, err = strconv.ParseUint(os.Args[2], 10, 64)
		if err != nil {
			fmt.Println("Args Width Format Error")
			return
		}
		height, err = strconv.ParseUint(os.Args[3], 10, 64)
		if err != nil {
			fmt.Println("Args Height Format Error")
			return
		}
	}
	imageCompress(filePath, uint(width), uint(height))
}