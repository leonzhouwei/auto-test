package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"qiniu.com/pandora-auto-test/comparer/biz"
)

const (
	start = "=== comparing start"
	end = "=== comparing end"
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

	fmt.Println("config: ", config)

	// query
	compareAll(*config)

	// ok
	fmt.Println("ALL OK")
}

func compareAll(conf biz.Config) {
	for i, expected := range conf.ExpectedArray {
		actual := (conf.ActualArray)[i]
		compare(expected, actual)
	}
}

func compare(expectedAPI, actualAPI biz.ConfigAPI) {
	fmt.Println(start)
	fmt.Println("expected API:", expectedAPI)
	fmt.Println("  actual API:", actualAPI)
	fmt.Println()

	// query the expected
	expectedBytes, err := biz.Query(expectedAPI)
	if err != nil {
		log.Fatal(err)
	}
	expected := string(expectedBytes)

	// query the actual
	actualBytes, err := biz.Query(actualAPI)
	if err != nil {
		log.Fatal(err)
	}
	actual := string(actualBytes)

	// compare and stat
	if expected != actual {
		fmt.Printf(errMsgFmt, actual, expected)
		os.Exit(1)
	}
	fmt.Println("EQUAL")
	fmt.Println(end)
}
