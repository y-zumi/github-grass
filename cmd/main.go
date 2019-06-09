package main

import (
	"fmt"

	"github.com/y-zumi/github-grass-cli/gitgrass"
)

func main() {
	userName := "y-zumi"
	client := gitgrass.NewClient(userName)
	resp, err := client.Get()
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}

	fmt.Printf("%s", resp)
}
