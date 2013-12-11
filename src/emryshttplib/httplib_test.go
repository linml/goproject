package emryshttplib

import (
	"fmt"
	"io/ioutil"
)

func T() {
	fmt.Print("test function")
	resp, err := Get("http://beego.me").Debug(true).getResponse()
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print("hello\n\n")
	}
	fmt.Print(data)

	//fmt.Print(data)
}
