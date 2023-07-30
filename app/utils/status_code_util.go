package utils

import "github.com/Firhan384/restfull-api-golang/app/schemas"

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	if text, ok := schemas.StatusTextMap[code]; ok {
		return text
	}

	return ""
}
