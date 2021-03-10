package main

import (
	"log"
	"myProject/test/config"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	port := config.Config.Port
	http.HandleFunc("/notes/", handler)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

/*
 *  用于处理文件的请求
 */
func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var ret string
	var err error
	if strings.HasPrefix(path, "/notes") {
		path = strings.Replace(path, "/notes", "vnotebook", 1)
		ret, err = getFile(path)
	} else {
		ret = "HelloWorld"
	}
	_, err = w.Write([]byte(ret))
	if err != nil {
		log.Fatal(err.Error())
	}
}

/**
 *	请求相应文件或者目录
 */
func getFile(path string) (string, error) {
	var err error
	var content string
	var expire time.Duration

	local, err := findLocal(path)
	if err == nil && len(local) != 0 {
		return local, nil
	}

	if strings.HasSuffix(path, ".md") {
		// 请求一个markdown文件
		content, err = returnMd(path)
		expire = markdownExpireTime
	} else if strings.Index(path, ".") != -1 {
		// 请求一个图片（base64编码）
		content, err = returnBase64Image(path)
		expire = imageExpireTime
	} else {
		content, err = returnDir(path)
		expire = directoryExpireTime
	}

	if err != nil {
		saveLocal(path, content, expire)
	}

	return content, err
}
