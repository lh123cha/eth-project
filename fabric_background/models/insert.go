package models

import (
	"log"
	"os/exec"
)

func Insert (id string,name string)  {
	log.Printf("/root/demo/startGoClient.sh \""+id+"\" \""+name+"\"")
	cmd := exec.Command("/bin/sh","-c","/root/demo/startGoClient.sh \""+id+"\" \""+name+"\"")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)

}
