package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	userName := "y-zumi"
	client := NewClient(userName)
	resp, err := client.Get()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Printf("%s", resp)
}

type HttpClient interface {
	Get() (string, error)
}

func NewClient(userName string) HttpClient {
	return &GithubGrass{
		userName: userName,
	}
}

type GithubGrass struct {
	userName string
}

func (g *GithubGrass) Get() (string, error) {
	res, err := http.Get(fmt.Sprintf("https://github.com/users/%s/contributions", g.userName))
	if err != nil {
		return "", err
	}
	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	return "", err
	// }
	// res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	ret, _ := doc.Find("svg").Parent().Html()

	return ret, nil
}
