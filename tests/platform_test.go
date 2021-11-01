package unittest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	utils_server "apiserver/v1/nomad-api"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestPlatformInstance(t *testing.T) {
	// go test -v tests/platform_test.go
	router := utils_server.InitRouter()

	// authorization test
	w := performRequest(router, "GET", "/metrics", strings.NewReader(`{`))
	t.Log(w.Body.String())
	assert.Equal(t, 200, w.Code)

	// id check Service
	payload := strings.NewReader(`{
    "storeId": "platform.test",
    "operation": "encrypt_batch",
    "items": [
        "$+v3K0huKLDouDveAZKxasA==$1$",
        "$Kg7PP/qJTnIYX+22Y2pa+A==$1$"
    ]
}`)
	w = performRequest(router, "POST", "/platform/cypher?x-ca-key=test", payload)
	t.Log(w.Body.String())
	assert.Equal(t, 401, w.Code)
}
