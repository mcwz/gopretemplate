package parseTemplate

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Parser struct {
	Template        *os.File
	TemplateContent string
	Tags            []Tag
	TemplateResult  string
}

func New(fileName string) (*Parser, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	parse := &Parser{Template: file}

	// 读取模板内容
	err = parse.readlines()
	if err != nil {
		return nil, err
	}
	return parse, nil
}

func (p *Parser) readlines() error {
	defer p.Template.Close()
	bs := make([]byte, 1024)

	var templateString string

	for {
		n, err := p.Template.Read(bs)

		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		templateString += string(bs)
	}
	p.TemplateContent = templateString

	return nil
}

func (p *Parser) parseTagLines() []string {
	reg := regexp.MustCompile(`#{2}\w+\b\s(([a-zA-Z]+=(\w+|,)+)\s?)+`)
	finded := reg.FindAllString(string(p.TemplateContent), -1)
	p.TemplateResult = reg.ReplaceAllString(p.TemplateContent, "")
	return finded
}

func (p *Parser) parseTags(findStrings []string) {
	for _, tagString := range findStrings {
		aTagString := strings.Replace(tagString, "##", "", -1)
		TagName := aTagString[0:strings.Index(aTagString, " ")]
		aTagString = aTagString[strings.Index(aTagString, " ")+1:]

		tag := &Tag{TagName: TagName}
		tag.TagAttr = make(map[string]string)

		tagAttrString := strings.Split(aTagString, " ")
		for _, aTag := range tagAttrString {
			keyValue := strings.Split(aTag, "=")
			if len(keyValue) == 2 {
				tag.TagAttr[keyValue[0]] = keyValue[1]
			} else {
				fmt.Println("why ", keyValue)
			}
		}
		p.Tags = append(p.Tags, *tag)
	}
}

func (p *Parser) Parse() {
	tagLines := p.parseTagLines()
	p.parseTags(tagLines)
}
