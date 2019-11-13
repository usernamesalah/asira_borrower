package tests

import (
	"asira_borrower/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestAgentLogin(t *testing.T) {
	RebuildData()

	api := router.NewRouter()

	server := httptest.NewServer(api)

	defer server.Close()

	e := httpexpect.New(t, server.URL)

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Basic "+clientBasicToken)
	})

	obj := auth.GET("/clientauth").
		Expect().
		Status(http.StatusOK).JSON().Object()

	admintoken := obj.Value("token").String().Raw()

	auth = e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+admintoken)
	})

	payload := map[string]interface{}{
		"key":      "081234567890",
		"password": "password",
	}

	// test valid response
	obj = auth.POST("/client/agent_login").WithJSON(payload).
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.Keys().Contains("expires_in", "token")

	// test invalid empty body
	obj = auth.POST("/client/agent_login").WithJSON(map[string]interface{}{}).
		Expect().
		Status(http.StatusBadRequest).JSON().Object()

	// test invalid client token
	auth = e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer wrongtoken")
	})
	obj = auth.POST("/client/agent_login").WithJSON(map[string]interface{}{}).
		Expect().
		Status(http.StatusUnauthorized).JSON().Object()
}

func TestAgentGetProfile(t *testing.T) {
	RebuildData()

	api := router.NewRouter()

	server := httptest.NewServer(api)

	defer server.Close()

	e := httpexpect.New(t, server.URL)

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Basic "+clientBasicToken)
	})

	borrowertoken := getAgentLoginToken(e, auth, "1")

	auth = e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+borrowertoken)
	})

	// valid response
	obj := auth.GET("/agent/profile").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.ContainsKey("username").ValueEqual("username", "agentK")

	// wrong token
	auth = e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer wrongtoken")
	})
	auth.GET("/agent/profile").WithJSON(map[string]interface{}{}).
		Expect().
		Status(http.StatusUnauthorized).JSON().Object()
}
