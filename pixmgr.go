package main

import (
	"fmt"
	"io/ioutil"
	"sync"
)

var FileDitto map[string]int
var mu sync.Mutex
var wg sync.WaitGroup

func main()  {
	dirPath := "testdata"
	Start(dirPath)
}

func Start(dirPath string) {
	FileDitto = make(map[string]int)

	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			checkDir(fmt.Sprintf("%s/%s", dirPath, fileInfo.Name()))
		} else {
			fmt.Println(fmt.Sprintf("VN_INFO:%s is not dir", fileInfo.Name()))
		}
	}

	wg.Wait()
	for key, value := range FileDitto {
		fmt.Println(key, "=", value)
	}
}

func checkDir(dirPath string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		//fmt.Println(dirPath)
		fileInfos, err := ioutil.ReadDir(dirPath)
		if err != nil {
			fmt.Println(fmt.Sprintf("VN_ERROR:for dir=%s, error:%v", dirPath, err))
			return
		}
		for _,fileInfo := range  fileInfos {
			if fileInfo.IsDir() {
				checkDir(fmt.Sprintf("%s/%s", dirPath, fileInfo.Name()))
			} else {
				//key := fmt.Sprintf("%s", fileInfo.Name())
				key := fmt.Sprintf("%s_%d", fileInfo.Name(), fileInfo.Size())
				pixPlus(key)
				//fmt.Println(fileInfo.Name(), ", ", fileInfo.Size(), ",", fileInfo.ModTime().Unix())
			}
		}
	}()
}

func pixPlus(key string) {
	mu.Lock()
	defer mu.Unlock()
	FileDitto[key] = FileDitto[key] + 1
}