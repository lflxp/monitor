package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckFileLsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func ExecCommand(commands string) string {
	out, err := exec.Command("bash", "-c", commands).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Println(commands, string(out))
	return string(out)
}
