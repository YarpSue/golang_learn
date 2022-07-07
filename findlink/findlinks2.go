package findlink

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findLinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// 这块能够用匿名函数进行实现的 在links中有
func findLinks(url string) ([]string, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != http.StatusOK {
		r.Body.Close()
		return nil, fmt.Errorf("getting %s : %s", url, r.Status)
	}
	doc, err := html.Parse(r.Body)
	r.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Parsing %s as HTML: %v\n", url, err)
	}

	return visit(nil, doc), nil
}


// n 其实是执行类型的一个指针
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr{
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
