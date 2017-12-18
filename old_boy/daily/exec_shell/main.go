package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)


func exec_shell(s string) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}

func main() {
	exec_shell("echo 'lipd' >  /root/asiainfo/lipd/golang/lipd.txt")
}