// 调用坚果云 webDav 来获取某个路径或者文件的内容
package main

import (
	"encoding/json"
	"log"
	"myProject/test/config"
	"studio-b12/gowebdav"
)

const url = "https://dav.jianguoyun.com/dav/"

var user string
var password string
var client *gowebdav.Client

type File struct {
	Name string
	Dir  bool
}

/**
 * 返回一个markdown文件的内容
 */
func returnMd(path string) (string, error) {

	bytes, err := client.Read(path)
	if err != nil {
		return "", err
	}
	text := string(bytes)
	saveLocal(path, text, markdownExpireTime)
	return text, nil
}

/**
 * 返回一个图片的base64格式
 */
func returnBase64Image(path string) (string, error) {
	return "", nil
}

/**
 * 返回目录内的所有文件的文件名、以及是否是文件夹
 */
func returnDir(path string) (string, error) {
	dirs, err := client.ReadDir(path)
	if err != nil {
		return "", err
	}
	points := make([]File, 0, len(dirs))
	for _, dir := range dirs {

		// 不显示图片文件夹
		if !(dir.Name() == "images") {
			points = append(points, File{dir.Name(), dir.IsDir()})
		}
	}

	bytes, err := json.Marshal(points)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func init() {
	user = config.Config.User
	password = config.Config.Password
	client = gowebdav.NewClient(url, user, password)
	err := client.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
}
