package h

import (
	"regexp"
	"sort"
	"strings"
)

var idfinder *regexp.Regexp
var classfinder *regexp.Regexp
var tagmatcher *regexp.Regexp

func init() {
	idfinder = regexp.MustCompile("#\\w+")
	classfinder = regexp.MustCompile("\\.\\w+")
	tagmatcher = regexp.MustCompile("^\\w+")
}

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
	// the actual tagname
	tagName := tagmatcher.FindString(h.TagName)
	if tagName == "" {
		tagName = "div"
	}

	// all the classes
	classes := []string{h.Attrs["class"]}
	for _, class := range classfinder.FindAllString(h.TagName, -1) {
		classes = append(classes, class[1:])
	}
	className := strings.TrimSpace(strings.Join(classes, " "))
	if className != "" {
		h.Attrs["class"] = className
	}

	// the last id
	ids := idfinder.FindAllString(h.TagName, -1)
	if len(ids) > 0 {
		if _, ok := h.Attrs["id"]; !ok {
			h.Attrs["id"] = ids[len(ids)-1][1:]
		}
	}

	// render the children
	var innerHTML string
	if h.Children != nil {
		innerHTML = h.Children.Render()
	}

	return "<" + tagName + h.Attrs.ToString() + ">" + innerHTML + "</" + tagName + ">"
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

type A map[string]string

func (attrs A) ToString() (content string) {
	var keys sort.StringSlice
	for key, _ := range attrs {
		keys = append(keys, key)
	}
	sort.Sort(keys)

	for _, key := range keys {
		content += " " + key + "='" + attrs[key] + "'"
	}
	return content
}
