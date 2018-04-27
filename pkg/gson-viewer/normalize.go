package gsonviewer

import (
	"regexp"
	"unsafe"
)

var regexForMatchingArray = regexp.MustCompile(`(\[\d\])`)

// Normalize returns normalized input text
// ex) list.[0].name => list.0.name
func NormalizeInputText(in *string) string {
	result := regexForMatchingArray.ReplaceAllFunc(*(*[]byte)(unsafe.Pointer(in)), func(matched []byte) []byte {
		return matched[1:][:len(matched)-2]
	})

	return string(result)
}
