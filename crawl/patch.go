package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	errPath    = "./errors.txt"
	outputPath = "./output.txt"
	searchUrl  = "http://www.baidu.com/s?wd="
)

func Baidu(target string) {
	targetUrl := searchUrl + target
	res, err := http.Get(targetUrl)
	if err != nil {
		panic(err)
	}
	outputFile := OpenOutputFile(outputPath)
	defer func() {
		if newerr := outputFile.Close(); newerr != nil {
			panic(err)
		}
	}()
	WriteResult(outputFile, res)
}

func OpenErrorFile(path string) *os.File {
	file, err := os.OpenFile(errPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0)
	if err != nil && os.IsNotExist(err) {
		newfile, newerr := os.Create(errPath)
		if newerr != nil {
			panic(err)
		}
		return newfile
	} else if err != nil {
		panic(err)
	}
	return file
}

func OpenOutputFile(path string) *os.File {
	file, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0)
	if err != nil && os.IsNotExist(err) {
		newfile, newerr := os.Create(outputPath)
		if newerr != nil {
			panic(err)
		}
		return newfile
	} else if err != nil {
		panic(err)
	}
	return file
}

func WriteResult(file *os.File, res *http.Response) {
	io.Copy(file, res.Body)
	res.Body.Close()
}

func main() {
	var target string
	fmt.Print("请输入你想百度的内容：")
	fmt.Scanf("%s", &target)
	Baidu(target)
}
