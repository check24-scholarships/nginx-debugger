package parser

import (
	"github.com/stretchr/testify/assert"
	"nginx_debugger/abstractNginxConfig"
	"testing"
)

type nextTokenTestCase struct {
	name          string
	tape          string
	expected      string
	remainingTape string
}

var nextTokenTestCases = []*nextTokenTestCase{
	&nextTokenTestCase{"2 words", "abc def", "abc", " def"},
	&nextTokenTestCase{"3 words", "abc def ghi", "abc", " def ghi"},
	&nextTokenTestCase{"2 words with space in beginning", " abc def", "abc", " def"},
	&nextTokenTestCase{"2 words with brace", " { def", "{", " def"},
	&nextTokenTestCase{"tokens without spaces", "abc{def", "abc", "{def"},
	&nextTokenTestCase{"tokens without spaces", "abc\ndef", "abc", "\ndef"},
	&nextTokenTestCase{"server block", "server {}", "server", " {}"},
	&nextTokenTestCase{"server block", " {}", "{", "}"},
}

func TestLexer_NextToken(t *testing.T) {
	for _, test := range nextTokenTestCases {
		t.Run(test.name, func(t *testing.T) {
			lexer := NewLexer(test.tape)
			actual := lexer.NextToken()

			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.remainingTape, lexer.tape)
		})
	}
}

func TestParser_Parse_Empty_ServerBlock(t *testing.T) {
	parser := NewParser("server {}")
	config, err := parser.Parse()

	assert.Nil(t, err)
	assert.Len(t, config.ServerBlocks, 1)
}

func TestParser_Parse_Directive(t *testing.T) {
	parser := NewParser("server {listen 10; proxy_pass lol;}")
	config, err := parser.Parse()

	assert.Nil(t, err)
	assert.Len(t, config.ServerBlocks, 1)
	assert.Len(t, config.ServerBlocks[0].Directives, 2)
	assert.Equal(t, abstractNginxConfig.DirectiveKeyListen, config.ServerBlocks[0].Directives[0].Key)
	assert.Equal(t, "10", config.ServerBlocks[0].Directives[0].Value)
	assert.Equal(t, abstractNginxConfig.DirectiveKeyProxyPass, config.ServerBlocks[0].Directives[1].Key)
	assert.Equal(t, "lol", config.ServerBlocks[0].Directives[1].Value)
}

func TestParser_Parse_Location_Block(t *testing.T) {
	parser := NewParser("server { location / { proxy_pass http://127.0.0.1:8080; } }")
	config, err := parser.Parse()

	assert.Nil(t, err)
	assert.Len(t, config.ServerBlocks, 1)
	assert.Len(t, config.ServerBlocks[0].LocationBlocks, 1)
	assert.Equal(t, "/", config.ServerBlocks[0].LocationBlocks[0].LocationMatch)
	assert.Equal(t, abstractNginxConfig.NoneMatchModifier, config.ServerBlocks[0].LocationBlocks[0].MatchModifier)
}
