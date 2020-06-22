package templ

import (
	"strings"
)

// 模板模式
type MessageRetriever interface {
	Message() string
}

type Templates interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

type Template struct{}

func (t *Template) first() string {
	return "hello"
}

func (t *Template) third() string {
	return "template"
}

func (t *Template) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}
