package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"qiniu.com/pandora-auto-test/comparer/biz"

	"github.com/qiniu/log.v1"
)

const (
	sepLine   = "----------------------------------------------------------------------"
	errMsgFmt = "ERROR: GOT %v BUT EXPECTED WAS %v\n"
)

func main() {
	// load config
	bytes1, err1 := ioutil.ReadFile("conf/comparer.conf")
	if err1 != nil {
		log.Fatal(err1)
	}

	tomlStr := string(bytes1)
	config, err2 := biz.NewConfig(tomlStr)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(sepLine)
	fmt.Println(config)
	log.SetOutputLevel(config.DebugLevel)

	// query
	fmt.Println(sepLine)
	compareAll(*config)

	// ok
	fmt.Println(sepLine)
	fmt.Println("OK")
}

func compareAll(conf biz.Config) {
	for i, expected := range conf.ExpectedArray {
		actual := (conf.ActualArray)[i]
		compare(expected, actual)
	}
}

func compare(expectedAPI, actualAPI biz.ConfigAPI) {
	fmt.Println("comparing", expectedAPI, "to", actualAPI, "...")

	// query the expected
	expectedBytes, err := biz.Query(expectedAPI)
	if err != nil {
		log.Fatal(err)
	}
	expected := string(expectedBytes)
	log.Debug("expected:", expected)

	// query the actual
	actualBytes, err := biz.Query(actualAPI)
	if err != nil {
		log.Fatal(err)
	}
	actual := string(actualBytes)
	log.Debug("actual:", actual)

	// compare and stat
	if expected != actual {
		errMsg := fmt.Sprintf(errMsgFmt, actual, expected)
		err := errors.New(errMsg)
		log.Fatal(err)
	}
	fmt.Println("EQUAL")
}
