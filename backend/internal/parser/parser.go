package parser

import (
	"errors"
	conf "nginx_debugger/abstractNginxConfig"
)

const (
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

func (l *Lexer) nextToken() string {
	currToken := ""
	end := 0

loop:
	for i, c := range l.tape {
		end = i

		switch c {
		case TokenSpace:
			continue
		case TokenBracketOpen:
		case TokenBracketClose:
		case TokenSemicolon:

			currToken = string(l.tape[i])
			break loop
		default:
			currToken += string(l.tape[i])
		}
	}

	l.tape = l.tape[:end]

	return currToken
}

func (l *Lexer) lex() []string {
	var tokens []string

	for len(l.tape) != 0 {
		token := l.nextToken()

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
		return nil, errors.New("expected opening brace")
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

	if p.popToken() == string(TokenSemicolon) {
		return nil, errors.New("expected token semicolon")
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
