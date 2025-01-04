package http

import (
	"strings"
	"testing"
)

func TestPathToRegex(t *testing.T) {
	regex, params := pathToRegex("/files/:file/:id/finish")
	wantRegex := "^/files/([^/]+)/([^/]+)/finish$"
	wantParams := []string{"file", "id"}

	if regex != wantRegex {
		t.Errorf("got %s, want %s", regex, wantRegex)
	}

	if strings.Join(params, ",") != strings.Join(wantParams, ",") {
		t.Errorf("got %s, want %s", params, wantParams)
	}
}
