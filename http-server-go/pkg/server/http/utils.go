package http

import (
	"fmt"
	"os"
	"strings"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func GetFilePath(headers map[string]string, fileName string) string {
	return fmt.Sprintf("%s%s", headers["X-Directory"], fileName)
}

func pathToRegex(path string) (string, []string) {
	segments := strings.Split(path, "/")

	var regexParts []string
	var params []string

	for _, seg := range segments {
		if seg == "" {
			continue
		}

		if strings.HasPrefix(seg, ":") {
			paramName := strings.TrimPrefix(seg, ":")
			params = append(params, paramName)
			regexParts = append(regexParts, `([^/]+)`)
		} else {
			regexParts = append(regexParts, seg)
		}
	}

	regex := "^/" + strings.Join(regexParts, "/") + "$"
	return regex, params
}

func parseAcceptEncoding(acceptEncoding string) []string {
	enc := strings.TrimSpace(acceptEncoding)
	if enc == "" {
		return []string{}
	}

	parts := strings.Split(enc, ",")
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}

	return parts
}
