package main

import (
	"fmt"

	"github.com/y-zumi/github-grass-cli/pkg/client/github"
)

func main() {
	// TODO: - input arg from terminal
	userName := "y-zumi"
	client := github.NewClient(userName)
	resp, err := client.Get()
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}

	fmt.Printf("%s", resp)
}
