package autoimport

import (
	"fmt"
	"unicode/utf8"
)

// encodePkgModPath encodes the path as used by the Go compiler. See the source:
// https://github.com/golang/go/blob/master/src/cmd/go/internal/sumweb/encode.go#L88-L114
func encodePkgModPath(s string) (encoding string, err error) {
	haveUpper := false

	for _, r := range s {
		if r == '!' || r >= utf8.RuneSelf {
			// This should be disallowed by CheckPath, but diagnose anyway.
			// The correctness of the encoding loop below depends on it.
			return "", fmt.Errorf("internal error: inconsistency in encodePkgModPath")
		}

		if 'A' <= r && r <= 'Z' {
			haveUpper = true
		}
	}

	if !haveUpper {
		return s, nil
	}

	var buf []byte

	for _, r := range s {
		if 'A' <= r && r <= 'Z' {
			buf = append(buf, '!', byte(r+'a'-'A'))
		} else {
			buf = append(buf, byte(r))
		}
	}

	return string(buf), nil
}
