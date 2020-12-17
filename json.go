package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/golang-commonmark/markdown"
)

type Item struct {
	Status int16  `json:"status"`
	Type   string `json:"type"`
	Desc   string `json:"desc"`
}

type Category struct {
	Name string `json:"name"`
	Desc string `json:"desc"`

	Items []*Item `json:"items"`
}

func PrintJSON(list []*Category) {
	jsonBytes, err := json.MarshalIndent(list, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
}

func dumpJSON(v interface{}) {
	jsonBytes, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
}

func Parse(r io.Reader) ([]*Category, error) {
	source, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	inHeading := false
	inBlock := false
	inLine := false
	level := 0

	list := make([]*Category, 0)
	current := (*Category)(nil)

	//Parse the markdown
	md := markdown.New()
	tokens := md.Parse(source)
	for _, t := range tokens {
		level = t.Level()
		switch t.(type) {
		case *markdown.HeadingOpen:
			inHeading = true
		case *markdown.HeadingClose:
			inHeading = false
		case *markdown.BlockquoteOpen:
			inBlock = true
		case *markdown.BlockquoteClose:
			inBlock = false
		case *markdown.ParagraphOpen:
			inLine = true
		case *markdown.ParagraphClose:
			inLine = false
		case *markdown.Inline:
			switch {
			case level == 1 && inHeading && !inBlock && !inLine: //category
				line := t.(*markdown.Inline)
				category := line.Content
				if current != nil {
					list = append(list, current)
				}
				current = &Category{Name: category}
			case level == 2 && !inHeading && inBlock && inLine: //description
				line := t.(*markdown.Inline)
				desc := line.Content
				if current != nil {
					current.Desc = desc
				}
			case level == 1 && !inHeading && !inBlock && inLine: //item
				itemToken := t.(*markdown.Inline)

				texts := make([]string, 0)
				for _, it := range itemToken.Children {
					switch it.(type) {
					case *markdown.Text:
						text := it.(*markdown.Text)
						content := strings.Trim(text.Content, " ")
						if len(content) > 0 {
							texts = append(texts, content)
						}
					case *markdown.CodeInline:
						code := it.(*markdown.CodeInline)
						content := strings.Trim(code.Content, " ")
						if len(content) > 0 {
							texts = append(texts, content)
						}
					}
				}
				if len(texts) != 3 {
					fmt.Println("inline", len(texts), texts)
					dumpJSON(itemToken)
					debug.PrintStack()
					return nil, errors.New("error markdown file")
				}

				status, err := strconv.ParseInt(texts[0], 10, 64)
				if err != nil {
					panic(err)
				}
				typeString := texts[1]
				desc := texts[2]
				if current != nil {
					current.Items = append(current.Items, &Item{
						Status: int16(status),
						Type:   typeString,
						Desc:   desc,
					})
				}
			}
		}
		// fmt.Println(t.Level(), reflect.TypeOf(t), inHeading, inBlock, inLine)
	}

	if current != nil {
		list = append(list, current)
	}
	return list, nil
}
