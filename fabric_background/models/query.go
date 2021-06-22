package models

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Query (key string) string {
	//log.Printf("/root/demo/startGoClient.sh \""+id+"\" \""+name+"\"")
	//cmd := exec.Command("/bin/sh","-c","/root/demo/startGoClient.sh \""+id+"\" \""+name+"\"")
	//log.Printf("Running command and waiting for it to finish...")
	//err := cmd.Run()
	//log.Printf("Command finished with error: %v", err)

	var arg1 string="query"
	var arg2 string="process"
	var out bytes.Buffer
	cmd := exec.Command("/bin/sh","-c","/root/demo/startGoClient.sh "+arg1+" "+arg2+" "+key)
	cmd.Stdout = &out
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
	fmt.Printf(out.String())
	return strings.Split(strings.Split(out.String(),"\n")[1],"\"")[3]

}