package main

import "os/exec"

func main() {
	_ = exec.Command("cmd", "/C", "shutdown -s -t 0").Run()
}
