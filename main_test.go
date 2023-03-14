package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthz(t *testing.T) {
	router := setupRouter()

	type args struct {
		sub, obj, act string
		code          int
	}
	var tests = []struct {
		name string
		args args
	}{
		// p, alice, /dataset1/*, GET
		{
			name: "alice GET dataset1/resource1",
			args: args{"alice", "/dataset1/resource1", "GET", http.StatusOK},
		},
		{
			name: "alice GET dataset1/resource2",
			args: args{"alice", "/dataset1/resource2", "GET", http.StatusOK},
		},
		{
			name: "bob GET dataset1/resource1",
			args: args{"bob", "/dataset1/resource1", "GET", http.StatusForbidden},
		},

		// p, alice, /dataset1/resource1, POST
		{
			name: "alice POST dataset1/resource1",
			args: args{"alice", "/dataset1/resource1", "POST", http.StatusOK},
		},
		{
			name: "bob POST dataset1/resource1",
			args: args{"bob", "/dataset1/resource1", "POST", http.StatusForbidden},
		},

		// p, bob, /dataset2/resource1, *
		{
			name: "bob GET dataset2/resource1",
			args: args{"bob", "/dataset2/resource1", "GET", http.StatusOK},
		},
		{
			name: "bob POST dataset2/resource1",
			args: args{"bob", "/dataset2/resource1", "POST", http.StatusOK},
		},
		{
			name: "bob PUT dataset2/resource1",
			args: args{"bob", "/dataset2/resource1", "PUT", http.StatusOK},
		},

		// p, dataset1_admin, /dataset1/*, *
		// g, cathy, dataset1_admin
		{
			name: "cathy GET dataset1/resource1",
			args: args{"cathy", "/dataset1/resource1", "GET", http.StatusOK},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			a := tt.args
			req, _ := http.NewRequest(a.act, a.obj, nil)
			req.SetBasicAuth(a.sub, "")

			router.ServeHTTP(w, req)
			if w.Code == http.StatusForbidden {
				t.Log(a.sub, a.act, a.obj, http.StatusText(w.Code))
			} else {
				t.Log(w.Body)
			}
			assert.Equal(t, a.code, w.Code)
		})
	}

}
