package git

import (
	"fmt"
	"regexp"
)

func (g *Git) Parse(branchName string) map[string]string {
	pattern := g.cfg.BranchPattern
	templateRegex := regexp.MustCompile(`\{(\w+)}`)
	keys := templateRegex.FindAllStringSubmatch(pattern, -1)

	indexes := templateRegex.FindAllStringSubmatchIndex(pattern, len(keys))

	rgxp := ""
	for i := range keys {
		name := keys[i][1]
		separatorPosition := indexes[i][1]
		if separatorPosition >= len(pattern) {
			rgxp += fmt.Sprintf("(?P<%s>.+)", name)
			continue
		}

		separator := pattern[separatorPosition]
		rgxp += fmt.Sprintf("(?P<%s>[^%c]+)", name, separator)
		rgxp += string(separator)

	}

	finalRegex := regexp.MustCompile(rgxp)

	values := finalRegex.FindAllStringSubmatch(branchName, -1)

	result := make(map[string]string)

	for i, _ := range keys {

		result[keys[i][1]] = values[0][i+1]
	}

	return result
}
