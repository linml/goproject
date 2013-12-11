package emryshttplib

import (
	"fmt"
	"io/ioutil"
)

func T() {
	fmt.Print("test function")
	resp, err := Get("http://www.baidu.com").Debug(true).getResponse()
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print("hello\n\n")
	}
	fmt.Print(string(data))

	//fmt.Print(data)
}
