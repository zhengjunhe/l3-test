package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func writeToFile(fileName string, content interface{}) {
	jbytes, err := json.MarshalIndent(content, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fileName, jbytes, 0666)
	if err != nil {
		fmt.Println("Failed to write to file:", fileName)
	}
	fmt.Println("tx is written to file: ", fileName)
}

func writeCoinStatics(fileName string, content TransferStatics2DB) {
	err := ioutil.WriteFile(fileName, []byte(content.Name), 0666)
	if err != nil {
		fmt.Println("Failed to write to file:", fileName)
	}

	for hash, res := range content.TxResult {
		err = ioutil.WriteFile(fileName, []byte(hash), 0666)
		if err != nil {
			fmt.Println("Failed to write to file:", fileName)
		}
		jbytes, err := json.MarshalIndent(&res, "", "\t")
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(fileName, jbytes, 0666)
		if err != nil {
			fmt.Println("Failed to write to file:", fileName)
		}
	}

	jbytes, err := json.MarshalIndent(content.Blocks, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fileName, jbytes, 0666)
	if err != nil {
		fmt.Println("Failed to write to file:", fileName)
	}

	jbytes, err = json.MarshalIndent(content.SuccessCnt, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fileName, jbytes, 0666)
	if err != nil {
		fmt.Println("Failed to write to file:", fileName)
	}
	fmt.Println("tx is written to file: ", fileName)
}
