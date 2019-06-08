package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	svg "github.com/ajstarks/svgo"
)

func main() {
	// http.Handle("/circle", http.HandlerFunc(circle))
	// err := http.ListenAndServe(":2003", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe:", err)
	// }
	userName := "y-zumi"
	client := NewClient(userName)
	resp, err := client.Get()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Printf("%s", resp)
}

func circle(w http.ResponseWriter, req *http.Request) {
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(250, 250, 125, "fill:none;stroke:black")
	s.End()
	// width := 500
	// height := 500
	// canvas := svg.New(os.Stdout)
	// canvas.Start(width, height)
	// canvas.Circle(width/2, height/2, 100)
	// canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
	// canvas.End()
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

	ret, _ := doc.Find("svg").Html()

	return ret, nil
}
