package browse

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
	"github.com/spf13/cobra"
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

func Execute() {
	resp, err := http.Get("https://www.nytimes.com/2020/09/14/technology/china-bytedance-tiktok-sale.html")
	check(err)
	tokenizer := html.NewTokenizer(resp.Body)
	depth := 0
	for {
		tagType := tokenizer.Next()
		fmt.Println(depth)
		switch tagType {
		case html.ErrorToken:
			resp.Body.Close()
			return
		case html.StartTagToken, html.EndTagToken:
			if tagType == html.StartTagToken {
				if depth != 0 {
					depth++
				}
			}
			if tagType == html.EndTagToken {
				if depth != 0 {
					depth--
					if depth == 0 {
						fmt.Println("Done")
					}
				}
			}
			break
		case html.TextToken:
			if depth == 0 {
				depth = 1
			}
			break
		}
	}
}
