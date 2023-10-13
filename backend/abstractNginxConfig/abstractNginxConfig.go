package abstractNginxConfig

import "errors"

type AbstractNginxConfig struct {
	ServerBlocks []ServerBlock
}

type ServerBlock struct {
	Directives     []Directive
	LocationBlocks []LocationBlock
}

type LocationBlock struct {
	Line          int
	MatchModifier LocationMatchModifier
	LocationMatch string
	Directives    []Directive
}

type LocationMatchModifier string

func LocationMatchModifierFromToken(token string) (*LocationMatchModifier, error) {
	for _, m := range LocationMatchModifiers {
		if string(m) == token {
			return &m, nil
		}
	}

	return nil, errors.New("location match modifier not found")
}

const (
	NoneMatchModifier                 = LocationMatchModifier("")
	ExactMatchModifier                = LocationMatchModifier("=")
	CaseSensitiveRegexMatchModifier   = LocationMatchModifier("~")
	CaseInsensitiveRegexMatchModifier = LocationMatchModifier("~*")
	BestNonRegexMatchModifier         = LocationMatchModifier("^~")
)

var LocationMatchModifiers = []LocationMatchModifier{
	NoneMatchModifier,
	ExactMatchModifier,
	CaseSensitiveRegexMatchModifier,
	CaseInsensitiveRegexMatchModifier,
	BestNonRegexMatchModifier,
}

type DirectiveKey string

func DirectiveFromToken(token string) (*DirectiveKey, error) {
	for _, d := range DirectiveKeys {
		if string(d) == token {
			return &d, nil
		}
	}

	return nil, errors.New("directive key not found")
}

const (
	DirectiveKeyListen     = DirectiveKey("listen")
	DirectiveKeyServerName = DirectiveKey("server_name")
	DirectiveKeyRoot       = DirectiveKey("root")
	DirectiveKeyProxyPass  = DirectiveKey("proxy_pass")
)

var DirectiveKeys = []DirectiveKey{
	DirectiveKeyListen,
	DirectiveKeyServerName,
	DirectiveKeyRoot,
	DirectiveKeyProxyPass,
}

type Directive struct {
	Line  int
	Key   DirectiveKey
	Value string
}
