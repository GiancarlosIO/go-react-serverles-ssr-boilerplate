package utils

import (
	"path"
	"strings"
)

// ShiftPath given a url-like string it returns the head and tail of it
// Example: given: /users/detail/123, output: user, /detail/123
// Example: given: /detail/123, output: detail, /123
// Example: given: /123, output: 123, ""
// It's useful to pass the rest of the url to the "child" controllers
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
