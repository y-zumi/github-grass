package github

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type GithubClient struct {
	userName string
}

func NewClient(userName string) *GithubClient {
	return &GithubClient{
		userName: userName,
	}
}

func (g *GithubClient) Get() (string, error) {
	res, err := http.Get(fmt.Sprintf("https://github.com/users/%s/contributions", g.userName))
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	ret, err := doc.Find("svg").Parent().Html()
	if err != nil {
		return "", err
	}

	return ret, nil
}
