package gitgrass

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

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

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	ret, _ := doc.Find("svg").Parent().Html()

	return ret, nil
}
