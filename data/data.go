package data

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io"
	"log"
	"math/rand"
	"os"
)

var (
	filePath	string	 // 毒鸡汤文件路径
	list		[]string // 毒鸡汤语句列表
	watcher		*fsnotify.Watcher
)

// 初始化
func Init(path string) (err error) {
	filePath = path
	if list, err = FileReadLine(filePath); err != nil {
		return err
	}
	log.Println(fmt.Sprintf("毒鸡汤总条数: %d", len(list)))

	// 监听文件变更
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	go listener()
	watcher.Add(filePath)
	return nil
}

// 获取一条毒鸡汤语句
func GetLine() string {
	n := rand.Intn(len(list))
	return list[n]
}

func listener() {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				lines, err := FileReadLine(filePath)
				if err == nil {
					list = lines
				} else {
					log.Println("文件读取失败: " + err.Error())
				}

				msg := fmt.Sprintf("文件发生变更: %s 总条数: %d", event.Name, len(list))
				log.Println(msg)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("文件变更监听失败: ", err)
		}
	}
}

// 文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func FileReadLine(path string) ([]string, error) {
	if ! FileExists(path) {
		return nil, errors.New("文件不存在")
	}

	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()


	buff := new(bytes.Buffer)
	lineList := make([]string, 0)
	r := bufio.NewReader(fp)
	for ; ; {
		lineByte, isPrefix, err := r.ReadLine()
		if err != nil || io.EOF == err {
			break
		}

		buff.Write(lineByte)
		if isPrefix {
			continue
		}

		lineList = append(lineList, buff.String())
		buff.Reset()
	}

	return lineList, nil
}