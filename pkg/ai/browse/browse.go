package browse

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var Cmd = &cobra.Command{
	Use: "browse",
	Short: "Print out a webpage",
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var url string

func init() {
	Cmd.Flags().StringVarP(&url, "url", "u", "", "Url to request")
	Cmd.MarkFlagRequired("url")
}

type Document struct {
	Title		string
	Text		[]string
}

func decompose(n *html.Node) string {
	text := ""
	if n.Type == html.TextNode {
		text = strings.Trim(n.Data, "\n\t ")
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text = fmt.Sprintf("%s%s", text, decompose(c))
	}
	return text
}

func Execute() {
	var document Document
	resp, err := http.Get(url)
	check(err)
	doc, err := html.Parse(resp.Body)
	check(err)
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Data == "title" {
			document.Title = n.FirstChild.Data
		} else {
			if n.Type == html.TextNode {
				if len(strings.Trim(n.Data, "\n\t ")) > 0 {
					document.Text = append(document.Text, decompose(n))
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Data != "script" && c.Data != "style" {
				f(c)
			}
		}
	}
	f(doc)
	y, err := yaml.Marshal(document)
	check(err)
	fmt.Println(string(y))
}
