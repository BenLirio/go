package browse

import (
	"fmt"
	"os"
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

type Document struct {
	Title		string
	Text		[]string
}

func decompose(n *html.Node) string {
	text := ""
	if n.Type == html.TextNode {
		text = n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text = fmt.Sprintf("%s%s", text, decompose(c))
	}
	return text
}

func Execute() {
	var document Document
	file, err := os.Open("index.html")
	check(err)
	doc, err := html.Parse(file)
	check(err)
	var f func(*html.Node) bool
	f = func(n *html.Node) bool {
		if n.Data == "title" {
			document.Title = n.FirstChild.Data
		}
		if n.Type == html.TextNode {
			if len(strings.Trim(n.Data, "\n\t")) > 0 {
				document.Text = append(document.Text, decompose(n.Parent))
				return false
			}
		}
		more := true
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Data != "script" && c.Data != "style" {
				if more == true {
					more = f(c)
				}
			}
		}
		return true
	}
	f(doc)
	y, err := yaml.Marshal(document)
	check(err)
	fmt.Println(string(y))
}
