package main

import (
	"log"
	"os/exec"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var arg1 string="sqs"
	var arg2 string="horseking"
	var arg3 string="sqs"
	var out bytes.Buffer
	cmd := exec.Command("/bin/sh","-c","/root/demo/startGoClient.sh "+arg1+" "+arg2+" "+arg3)
	cmd.Stdout = &out
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
	fmt.Println(strings.Split(strings.Split(out.String(),"\n")[1],"\"")[3])
}
