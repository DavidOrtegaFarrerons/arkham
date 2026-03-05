package git

import (
	"fmt"
	"regexp"
)

func (g *Git) Format(placeholderValues map[string]string) string {
	templateRegex := regexp.MustCompile(`\{(\w+)}`)
	result := templateRegex.ReplaceAllStringFunc(g.cfg.CommitTemplate, func(match string) string {
		matchWithoutBrackets := match[1 : len(match)-1]
		v, ok := placeholderValues[matchWithoutBrackets]
		if !ok {
			fmt.Printf("There was no match found for the placeholder %s", match) //This should be logged somewhere better
			return match
		}

		return v
	})

	return result
}
