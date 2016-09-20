package h

import "strings"

type A map[string]string

func Element(tagname string, attrs A, children H) HElement {
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
	Children H
}

func (h HElement) Render() string {
	attrs := ""
	if h.Attrs != nil {
		for key, value := range h.Attrs {
			attrs += " " + strings.TrimSpace(key) + "='" + strings.TrimSpace(value) + "'"
		}
	}

	var innerHTML string
	if h.Children != nil {
		innerHTML = h.Children.Render()
	}

	return "<" + h.TagName + attrs + ">" + innerHTML + "</" + h.TagName + ">"
}

type HText struct {
	Text string
}

func (h HText) Render() string {
	return h.Text
}

type HH []H

func (hh HH) Render() string {
	content := ""
	for _, child := range hh {
		content += child.Render()
	}
	return content
}
