package main

import (
	"fmt"
	"gateway-client/utils"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

const GATEWAY_URL = "http://gateway"

func main() {
	res := execShellFile("/root/new.sh")
	devs := strings.Split(res, " ")
	t := time.Now()
	curTime := t.Format("2006-01-02 15:04:05")
	for idx, dev := range devs {
		fmt.Printf("idx %d, val %s", idx, dev)
		msg := utils.BuildMsg(
			100,
			"http://baidu.com",
			"有人连接啦",
			"router-notify",
			dev,
			curTime,
			"",
		)
		utils.SendToGateway(GATEWAY_URL, msg)
	}

}

func execShellFile(f string) (result string) {
	cmd := exec.Command("/bin/sh", f)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: ", err.Error())
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		return
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println("ReadAll stderr: ", err.Error())
		return
	}

	if len(bytesErr) != 0 {
		fmt.Printf("stderr is not nil: %s", bytesErr)
		return
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: ", err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait: ", err.Error())
		return
	}

	fmt.Printf("stdout: %s", bytes)
	result = string(bytes)
	return
}
