package main

import (
	"bytes"
	"os/exec"
	"log"
)

func State() string {
	cmd := exec.Command("firewall-cmd", "--state")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Printf("error on state request: %v", err)
	}
	return out.String()
}
