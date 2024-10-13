package main

import (
	"encoding/base64"
	"fmt"

	"go_hw/task2/client"
	"go_hw/task2/server"
)

func main() {

	s := server.NewServer()
	go s.Start()

	client := client.NewClient("http://localhost:8080")

	version, err := client.GetVersion()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Received version: %s\n", version)

	str := "милота"
	encodedStr := base64.StdEncoding.EncodeToString([]byte(str))
	dstring, err := client.PostDecode(encodedStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dstring)

	status, err := client.GetHardOp()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Operation completed successfully. Status:%d\n", status)

}
