package parser

import (
	"errors"
	"fmt"
	conf "nginx_debugger/abstractNginxConfig"
)

const (
	TokenNewLine      = '\n'
	TokenSpace        = ' '
	TokenBracketOpen  = '{'
	TokenBracketClose = '}'
	TokenSemicolon    = ';'

	TokenServer   = "server"
	TokenLocation = "location"
)

type Lexer struct {
	tape string
}

func NewLexer(tape string) *Lexer {
	return &Lexer{
		tape: tape,
	}
}

func (l *Lexer) NextToken() string {

	currToken := ""

	for l.tape != "" {
		c := l.tape[0]
		switch c {
		case TokenNewLine:
			fallthrough
		case TokenSpace:
			if len(currToken) == 0 {
				l.tape = l.tape[1:]
			} else {
				return currToken
			}
		case TokenBracketOpen:
			fallthrough
		case TokenBracketClose:
			fallthrough
		case TokenSemicolon:
			if len(currToken) == 0 {
				l.tape = l.tape[1:]
				return string(c)
			} else {
				return currToken
			}
		default:
			currToken += string(c)
			l.tape = l.tape[1:]
		}
	}

	return currToken
}

func (l *Lexer) lex() []string {
	var tokens []string

	for len(l.tape) != 0 {
		token := l.NextToken()

		tokens = append(tokens, token)
	}

	return tokens
}

type Parser struct {
	tokens []string
}

func NewParser(tape string) *Parser {
	lexer := NewLexer(tape)

	return &Parser{
		tokens: lexer.lex(),
	}
}

func (p *Parser) peekToken() string {
	if len(p.tokens) == 0 {
		return "EOF"
	}

	return p.tokens[0]
}

func (p *Parser) popToken() string {
	if len(p.tokens) == 0 {
		return "EOF"
	}

	token := p.tokens[0]
	p.tokens = p.tokens[1:]

	return token
}

func (p *Parser) ParseServerBlock() (*conf.ServerBlock, error) {
	serverToken := p.popToken()
	if serverToken != TokenServer {
		return nil, errors.New("expected token server")
	}

	braceToken := p.popToken()
	if braceToken != string(TokenBracketOpen) {
		return nil, errors.New(fmt.Sprintf("expected opening brace, received %s", braceToken))
	}

	var locationBlocks []conf.LocationBlock
	var directives []conf.Directive
	for p.peekToken() != string(TokenBracketClose) {

		if p.peekToken() == TokenLocation {
			locationBlock, err := p.ParseLocationBlock()
			if err != nil {
				return nil, err
			}

			locationBlocks = append(locationBlocks, *locationBlock)
		} else {
			directive, err := p.ParseDirective()
			if err != nil {
				return nil, err
			}

			directives = append(directives, *directive)
		}
	}

	p.popToken()

	return &conf.ServerBlock{
		Directives:     directives,
		LocationBlocks: locationBlocks,
	}, nil
}

func (p *Parser) ParseDirective() (*conf.Directive, error) {
	key, err := conf.DirectiveFromToken(p.popToken())
	if err != nil {
		return nil, err
	}

	value := p.popToken()

	semicolonToken := p.popToken()
	if semicolonToken != string(TokenSemicolon) {
		return nil, errors.New(fmt.Sprintf("expected opening brace, received %s", semicolonToken))
	}

	return &conf.Directive{
		Key:   *key,
		Value: value,
	}, nil
}

func (p *Parser) ParseLocationBlock() (*conf.LocationBlock, error) {
	serverToken := p.popToken()
	if serverToken != TokenLocation {
		return nil, errors.New("expected token location")
	}

	braceToken := p.popToken()
	if braceToken != string(TokenBracketOpen) {
		return nil, errors.New("expected opening brace")
	}

	var directives []conf.Directive
	for p.peekToken() != string(TokenBracketClose) {
		directive, err := p.ParseDirective()
		if err != nil {
			return nil, err
		}

		directives = append(directives, *directive)
	}

	p.popToken()

	return &conf.LocationBlock{
		Directives: directives,
	}, nil
}

func (p *Parser) Parse() (*conf.AbstractNginxConfig, error) {
	var serverBlocks []conf.ServerBlock

	for {
		block, err := p.ParseServerBlock()
		if err != nil {
			return nil, err
		}

		serverBlocks = append(serverBlocks, *block)

		if p.peekToken() == "EOF" {
			break
		}
	}

	return &conf.AbstractNginxConfig{
		ServerBlocks: serverBlocks,
	}, nil
}
