package tools

import (
	"io"
)

type ElementHandler func(el *ElementData)

type Parser struct {
	handlers []ElementHandler
}

func (p *Parser) OnElement(h ElementHandler) {
	p.handlers = append(p.handlers, h)
}

func (p *Parser) Parse(r io.Reader) chan Result {
	return Parse(r)
}
