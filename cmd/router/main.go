package main

import (
	"fmt"
	"gateway-client/utils"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

const (
	GATEWAY_URL = "http://gateway/u"
	CACHE_FILE  = "./router.cache"
)

func main() {
	res := execShellFile("./new.sh")
	devs := strings.Split(res, " ")
	t := time.Now().UTC()
	t = t.Add(8 * time.Hour)
	curTime := t.Format("2006-01-02 15:04:05")
	tmpMap := make(map[interface{}]interface{})
	oldMap, _ := utils.LoadFromFile(CACHE_FILE)
	for _, dev := range devs {
		kvs := strings.Split(dev, "#")
		if oldMap[kvs[1]] != kvs[0] {
			// new
			msg := utils.BuildMsg(
				100,
				"http://router",
				"有人连接啦",
				"router-notify",
				kvs[0],
				curTime,
				kvs[1],
			)
			utils.SendToGateway(GATEWAY_URL, msg)
		}
		tmpMap[kvs[1]] = kvs[0]
	}
	utils.SaveToFile(CACHE_FILE, tmpMap)

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
