package tags

import (
	"fmt"
	"strings"
)

var validTags = map[string]bool{
	"json": true,
	"yaml": true,
}

// formatTags formats Go struct tags.
func FormatTags(cmdlineTags string) (tagList []string, err error) {
	if cmdlineTags == "" {
		tagList = append(tagList, "json")
	} else {
		tagList = strings.Split(cmdlineTags, ",")
	}
	for _, tag := range tagList {
		if !validTags[tag] {
			return nil, fmt.Errorf("invalid tag: %s\n", tag)
		}
	}
	return
}
