package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
	"io"
	"time"
)

func main() {
	path := "F:/sort/input"
	timeBegin := time.Now()
	fileList, e := ioutil.ReadDir(path)
	if e != nil{
		fmt.Println("error")
		return
	}
	for _, v := range fileList{
		fileName := path + "/" + v.Name()
		file, err := os.Open(fileName)
		if err != nil{
			fmt.Println("error")
			return
		}

		rd := bufio.NewReader(file)
		for {
			line, err := rd.ReadString('\n')
			if err != nil || io.EOF == err{
				break
			}
			fmt.Println(line)
		}
		defer file.Close()

	}
	timeEnd := time.Now()

	fmt.Println(timeBegin)
	fmt.Println(timeEnd)
}