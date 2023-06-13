package parse

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func Parse(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return parseString(string(content))
}

func parseString(content string) (string, error) {
	matcher0, err := regexp.Compile(`\$\$\$\$`)
	if err != nil {
		return "", err
	}
	output0 := matcher0.ReplaceAllString(content, "")

	matcher1, err := regexp.Compile(`(?U)\${2}[^\$]+\${2}`)
	if err != nil {
		return "", err
	}
	output1 := matcher1.ReplaceAllStringFunc(output0, func(s string) string {
		s = strings.TrimFunc(s, func(r rune) bool {
			return r == '$'
		})
		s = url.QueryEscape(s)
		return fmt.Sprintf(`![](https://www.zhihu.com/equation?tex=\bbox[white]{%s})`, s)
	})
	matcher2, err := regexp.Compile(`(?U)\${1}[^\$]+\${1}`)
	if err != nil {
		return "", err
	}
	output2 := matcher2.ReplaceAllStringFunc(output1, func(s string) string {
		s = strings.TrimFunc(s, func(r rune) bool {
			return r == '$'
		})
		s = url.QueryEscape(s)
		return fmt.Sprintf(`![](https://www.zhihu.com/equation?tex=\bbox[white]{%s})`, s)
	})
	return output2, nil
}
