package main

import "fmt"

var config *RuntimeConfig

func init() {
	c := getConfig()
	config = c

	fmt.Printf("Configuration is %#v", config)
}

func main() {

}
