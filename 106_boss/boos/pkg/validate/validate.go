package validate

import "regexp"

func Regex(re *regexp.Regexp, value []byte) bool {
	return re.Match(value)
}
