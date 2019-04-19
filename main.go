package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const measureLoops = 3

func init() {
	gob.Register(map[string]interface{}{})
	gob.Register([]interface{}{})
	gob.Register(&List{})
}

func benchStdLibJsonStruct(images *List) {
	var data []byte
	var err error
	measure("stdlib json.Marshal(struct)", func() {
		data, err = json.Marshal(images)
		must(err)
	})

	measure("stdlib json.Unmarshal(struct)", func() {
		obj := List{}
		err = json.Unmarshal(data, &obj)
		must(err)
	})

	err = ioutil.WriteFile("images.json", data, os.ModePerm)
	must(err)

}

func benchStdLibJsonMap(tmp *List) {
	var data []byte
	var err error

	data, err = json.Marshal(tmp)
	images := make(map[string]interface{})
	err = json.Unmarshal(data, &images)
	must(err)

	measure("stdlib json.Marshal(map)", func() {
		_, err = json.Marshal(images)
		must(err)
	})

	measure("stdlib json.Unmarshal(map)", func() {
		obj := make(map[string]interface{})
		err = json.Unmarshal(data, &obj)
		must(err)
	})

}

func benchStdLibGOBMap(tmp *List) {
	var data []byte
	var err error

	data, err = json.Marshal(tmp)
	images := make(map[string]interface{})
	err = json.Unmarshal(data, &images)
	must(err)

	measure("stdlib gob.Encode(map)", func() {
		btmp := &bytes.Buffer{}
		enc := gob.NewEncoder(btmp)
		err = enc.Encode(images)
		must(err)
		data = btmp.Bytes()
	})

	measure("stdlib gob.Decode(map)", func() {
		dec := gob.NewDecoder(bytes.NewBuffer(data))
		obj := make(map[string]interface{})
		err = dec.Decode(&obj)
		must(err)
	})
}

func benchStdLibGOBStruct(images *List) {
	var data []byte
	var err error

	measure("stdlib gob.Encode(struct)", func() {
		btmp := &bytes.Buffer{}
		enc := gob.NewEncoder(btmp)
		err = enc.Encode(images)
		must(err)
		data = btmp.Bytes()
	})

	measure("stdlib gob.Decode(struct)", func() {
		dec := gob.NewDecoder(bytes.NewBuffer(data))
		obj := &List{}
		err = dec.Decode(&obj)
		must(err)
	})
}

func main() {
	images := GenerateImageMetaData(1000000)
	benchStdLibGOBMap(images)
	benchStdLibGOBStruct(images)
	benchStdLibJsonMap(images)
	benchStdLibJsonStruct(images)
}

func measure(what string, f func()) {
	start := time.Now()
	for i := 0; i < measureLoops; i++ {
		f()
	}
	fmt.Println(what, (time.Now().Sub(start) / measureLoops).String())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
