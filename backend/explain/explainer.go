package explain

import (
	"fmt"
	"nginx_debugger/abstractNginxConfig"
	"strings"
)

func ExplainNginxConfig(abstractConfig abstractNginxConfig.AbstractNginxConfig) Explanation {
	explanation := make(map[int]string)

	for _, serverBlock := range abstractConfig.ServerBlocks {
		for _, directive := range serverBlock.Directives {
			explanation[directive.Line] = explainDirective(directive)
		}

		for locationBlockIndex, locationBlock := range serverBlock.LocationBlocks {
			explanation[locationBlock.Line] = explainLocationMatch(locationBlockIndex, serverBlock.LocationBlocks)

			for _, directive := range serverBlock.Directives {
				explanation[directive.Line] = explainDirective(directive)
			}
		}
	}

	return explanation
}

func explainDirective(directive abstractNginxConfig.Directive) string {
	switch directive.Key {
	case abstractNginxConfig.DirectiveKeyListen:
		return fmt.Sprintf("listens for http connections on port %s", directive.Value)

	case abstractNginxConfig.DirectiveKeyServerName:
		serverNames := strings.Split(directive.Value, " ")

		var filteredServerNames []string
		for _, serverName := range serverNames {
			if strings.TrimSpace(serverName) != "" {
				filteredServerNames = append(filteredServerNames, fmt.Sprintf("'%s'", serverName))
			}
		}

		return fmt.Sprintf("host header has to be %s", strings.Join(filteredServerNames, " or "))

	case abstractNginxConfig.DirectiveKeyProxyPass:
		return fmt.Sprintf("passes matching requests to the proxied server at '%s'", directive.Value)

	case abstractNginxConfig.DirectiveKeyRoot:
		return fmt.Sprintf("specifies the root directory that will be used to search for a file, which is '%s'", directive.Value)
	}

	return "unknown, please consult documentation ;)"
}

func explainLocationMatch(locationBlockIndex int, allLocationBlocks []abstractNginxConfig.LocationBlock) string {
	locationBlock := allLocationBlocks[locationBlockIndex]

	switch locationBlock.MatchModifier {
	case abstractNginxConfig.NoneMatchModifier:
		return fmt.Sprintf("will be choosen if '%s' is the longest prefix match of the URI and no other location "+
			"block with a modifier matches", locationBlock.LocationMatch)

	case abstractNginxConfig.ExactMatchModifier:
		return explainExactLocationMatch(locationBlockIndex, allLocationBlocks)

	case abstractNginxConfig.CaseSensitiveRegexMatchModifier:
		return explainCaseSensitiveRegexLocationMatch(locationBlockIndex, allLocationBlocks)

	case abstractNginxConfig.CaseInsensitiveRegexMatchModifier:
		return fmt.Sprintf("will be choosen if URI matches the case-insensitive regex '%s'", locationBlock.LocationMatch)

	case abstractNginxConfig.BestNonRegexMatchModifier:
		return explainBestNonRegexLocationMatch(locationBlockIndex, allLocationBlocks)
	}

	return ""
}

func explainExactLocationMatch(locationBlockIndex int, allLocationBlocks []abstractNginxConfig.LocationBlock) string {
	locationBlock := allLocationBlocks[locationBlockIndex]

	for i := 0; i < locationBlockIndex; i++ {
		otherLocationBlock := allLocationBlocks[i]

		if otherLocationBlock.MatchModifier == abstractNginxConfig.ExactMatchModifier &&
			otherLocationBlock.LocationMatch == locationBlock.LocationMatch {
			return fmt.Sprintf("this location block will never be choosen, because location block in line %d will "+
				"always have a higher precedence", otherLocationBlock.Line)
		}
	}

	return fmt.Sprintf("will be choosen if URI is exactly '%s'", locationBlock.LocationMatch)
}

func explainBestNonRegexLocationMatch(locationBlockIndex int, allLocationBlocks []abstractNginxConfig.LocationBlock) string {
	locationBlock := allLocationBlocks[locationBlockIndex]

	var exactMatchesWithPrefix []string

	for _, otherLocationBlock := range allLocationBlocks {
		if otherLocationBlock.MatchModifier == abstractNginxConfig.ExactMatchModifier &&
			strings.HasPrefix(otherLocationBlock.LocationMatch, locationBlock.LocationMatch) {
			exactMatchesWithPrefix = append(exactMatchesWithPrefix, fmt.Sprintf("'%s'", otherLocationBlock.LocationMatch))
		}
	}

	if len(exactMatchesWithPrefix) > 0 {
		return fmt.Sprintf("will be choosen if URI start with '%s' and URI is not exactly one of: %s",
			strings.Join(exactMatchesWithPrefix, ","))
	}

	return fmt.Sprintf("will be choosen if URI start with '%s'", locationBlock.LocationMatch)
}

func explainCaseSensitiveRegexLocationMatch(locationBlockIndex int, allLocationBlocks []abstractNginxConfig.LocationBlock) string {
	locationBlock := allLocationBlocks[locationBlockIndex]

	var exactMatchesWithPrefix []string
	var bestNonRegexLocationMatches []string

	for _, otherLocationBlock := range allLocationBlocks {
		if otherLocationBlock.MatchModifier == abstractNginxConfig.ExactMatchModifier &&
			strings.HasPrefix(otherLocationBlock.LocationMatch, locationBlock.LocationMatch) {
			exactMatchesWithPrefix = append(exactMatchesWithPrefix, fmt.Sprintf("'%s'", otherLocationBlock.LocationMatch))
		}

		if otherLocationBlock.MatchModifier == abstractNginxConfig.BestNonRegexMatchModifier {
			bestNonRegexLocationMatches = append(bestNonRegexLocationMatches, fmt.Sprintf("'%s'", otherLocationBlock.LocationMatch))
		}
	}

	if len(exactMatchesWithPrefix) == 0 && len(bestNonRegexLocationMatches) == 0 {
		return fmt.Sprintf("will be choosen if URI matches the case-sensitive regex '%s'", locationBlock.LocationMatch)
	}

	//if len(exactMatchesWithPrefix) > 0 && len(bestNonRegexLocationMatches) == 0 {
	//	return fmt.Sprintf("will be choosen if URI start with '%s' and URI is not exactly one of: %s",
	//		strings.Join(exactMatchesWithPrefix, ","))
	//}

	// TODO
	return fmt.Sprintf("will be choosen if URI start with '%s'", locationBlock.LocationMatch)
}
