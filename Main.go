package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	battery := "BAT0"
	path := "/sys/class/power_supply"

	path = path + "/" + battery

	fmt.Printf("%v%% %v\n",
		capacity(path),
		status(path))
}

func capacity(path string) string {
	file := "capacity"
	return read(path + "/" + file)
}

func status(path string) string {
	file := "status"
	return read(path + "/" + file)
}

func read(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSuffix(string(data), "\n")
}
