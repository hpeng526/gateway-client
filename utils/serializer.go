package utils

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
)

func SaveToFile(filePath string, res map[interface{}]interface{}) (err error) {
	tempFile, err := os.Create(filePath)
	defer tempFile.Close()
	var datas bytes.Buffer
	enc := gob.NewEncoder(&datas)
	err = enc.Encode(res)
	tempFile.Write(datas.Bytes())
	return
}

func LoadFromFile(filePath string) (res map[interface{}]interface{}, err error) {
	tempFile, err := os.Open(filePath)
	defer tempFile.Close()
	datas, err := ioutil.ReadAll(tempFile)
	dec := gob.NewDecoder(bytes.NewBuffer(datas))
	dec.Decode(&res)
	return
}
