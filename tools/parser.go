package tools

import (
	"context"
	"io"
)

type ElementHandler func(el *ElementData)

type Parser struct {
	handlers []ElementHandler
}

func (p *Parser) OnElement(h ElementHandler) {
	p.handlers = append(p.handlers, h)
}

func (p *Parser) Parse(ctx context.Context, r io.Reader) chan Result {
	return Parse(ctx, r)
}
