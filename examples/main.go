package main

import (
	"fmt"
	"strings"

	link "github.com/harchit/linkparser"
)

var html = `
<html>
<body>
	<h1>Hello!</h1>
	<a href="/other-page">A link to anothe page"</a>
	<a href="/page-two">link to page two"</a>
</body>
</html>`

func main() {
	r := strings.NewReader(html)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}