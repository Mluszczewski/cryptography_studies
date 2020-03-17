package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func prepareText(text string) string {
	cmd := exec.Command("bash", "prepareText.sh")
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}
