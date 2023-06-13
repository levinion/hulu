package main

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func main() {
	content, err := os.ReadFile("test/test.md")
	if err != nil {
		panic(err)
	}
	matcher, err := regexp.Compile(`(?U)\${2}[^\$]+\${2}`)
	if err != nil {
		panic(err)
	}
	output := matcher.ReplaceAllStringFunc(string(content), func(s string) string {
		s = strings.TrimFunc(s, func(r rune) bool {
			return r == '$'
		})
		s = url.QueryEscape(s)
		return fmt.Sprintf(`![](https://www.zhihu.com/equation?tex=\bbox[white]{%s})`, s)
	})
	matcher2, err := regexp.Compile(`(?U)\${1}[^\$]+\${1}`)
	if err != nil {
		panic(err)
	}
	output2 := matcher2.ReplaceAllStringFunc(output, func(s string) string {
		s = strings.TrimFunc(s, func(r rune) bool {
			return r == '$'
		})
		s = url.QueryEscape(s)
		return fmt.Sprintf(`![](https://www.zhihu.com/equation?tex=\bbox[white]{%s})`, s)
	})
	fmt.Println(output2)
}
