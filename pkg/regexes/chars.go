package regexes

import "regexp"

var NotNumbers = regexp.MustCompile(`[^0-9]+`)
