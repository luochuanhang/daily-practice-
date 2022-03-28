package regex

import "regexp"

// Phone regex.
var Phone = regexp.MustCompile(`^1(\d{10})$`)
