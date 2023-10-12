package abstractNginxConfig

type AbstractNginxConfig struct {
	serverBlocks []ServerBlock
}

type ServerBlock struct {
	directives     []Directive[string]
	locationBlocks []LocationBlock
}

type LocationBlock struct {
	line          int
	matchModifier LocationMatchModifier
	locationMatch string
	directives    []Directive[string]
}

type LocationMatchModifier string

const (
	NoneMatchModifier          = LocationMatchModifier("")
	ExactMatchModifier         = LocationMatchModifier("=")
	CaseSensitiveMatchModifier = LocationMatchModifier("~")
	BestNonRegexMatchModifier  = LocationMatchModifier("^~")
)

type DirectiveKey string

const (
	DirectiveKeyListen     = DirectiveKey("listen")
	DirectiveKeyServerName = DirectiveKey("server_name")
	DirectiveKeyRoot       = DirectiveKey("root")
	DirectiveKeyProxyPass  = DirectiveKey("proxy_pass")
)

type Directive[T any] struct {
	line  int
	key   DirectiveKey
	value T
}
