package h

import "strings"

func Element(tagname string, attrs A, children []H) HElement {
	return HElement{tagname, attrs, children}
}

func Text(text string) HText {
	return HText{text}
}

type H interface {
	Render() string
}

type HElement struct {
	TagName  string
	Attrs    A
	Children []H
}

type HText struct {
	Text string
}

func (h HElement) Render() string {
	attrs := ""
	if h.Attrs != nil {
		for key, value := range h.Attrs {
			attrs += " " + strings.TrimSpace(key) + "='" + strings.TrimSpace(value) + "'"
		}
	}

	content := ""
	if h.Children != nil {
		for _, child := range h.Children {
			content += child.Render()
		}
	}

	return "<" + h.TagName + attrs + ">" + content + "</" + h.TagName + ">"
}

func (h HText) Render() string {
	return h.Text
}

type A map[string]string
