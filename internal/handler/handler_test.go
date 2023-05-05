package handler

import (
	"github.com/dobb2/zenTotem/internal/entity"
	"github.com/dobb2/zenTotem/internal/logging"
	"github.com/dobb2/zenTotem/internal/storage/mocks"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func testRequest(t *testing.T, ts *httptest.Server, path, method string, body io.Reader) (int, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	defer resp.Body.Close()
	return resp.StatusCode, string(respBody)
}

func TestUserHandler_IncrementVal(t *testing.T) {
	type want struct {
		code int
		json string
	}
	tests := []struct {
		name   string
		url    string
		json   string
		method string
		want   want
	}{
		{
			name:   "positive redis test #1",
			url:    "/redis/incr",
			json:   `{"key":"age","value":19}`,
			method: "POST",
			want: want{
				code: http.StatusOK,
				json: `{"value":19}`,
			},
		},
		{
			name:   "positive redis test #2",
			url:    "/redis/incr",
			json:   `{"key":"age","value":7}`,
			method: "POST",
			want: want{
				code: http.StatusOK,
				json: `{"value":26}`,
			},
		},
		{
			name:   "negative redis test #3",
			url:    "/redis/incr",
			json:   `{"key":"HeapInuse","value":"23"}`,
			method: "POST",
			want: want{
				code: http.StatusBadRequest,
			},
		},
		{
			name:   "negative redis test #4",
			url:    "/redis/incr",
			json:   `{"key":"HeapInuse","value":23.33}`,
			method: "POST",
			want: want{
				code: http.StatusBadRequest,
			},
		},
		{
			name:   "negative redis test #5",
			url:    "/redis/incr",
			json:   `{"v":"HeapInuse","vz":23.33}`,
			method: "POST",
			want: want{
				code: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := mocks.NewCacher(t)
			storage := mocks.NewStorer(t)
			logger := logging.CreateLogger()

			cache.
				On("Increment", entity.Element{Key: "age", Value: 19}).
				Return(entity.Element{Value: 19}, nil).
				Maybe()

			cache.
				On("Increment", entity.Element{Key: "age", Value: 7}).
				Return(entity.Element{Value: 26}, nil).
				Maybe()

			a := New(storage, cache, logger)

			r := func(u UserHandler) chi.Router {
				r := chi.NewRouter()
				r.Post("/redis/incr", u.IncrementVal)
				return r
			}(a)
			ts := httptest.NewServer(r)
			defer ts.Close()

			statusCode, _ := testRequest(t, ts, tt.url, tt.method, strings.NewReader(tt.json))
			assert.Equal(t, tt.want.code, statusCode)
		})
	}
}

func TestUserHandler_CreateUser(t *testing.T) {
	type want struct {
		code int
		json string
	}
	tests := []struct {
		name   string
		url    string
		json   string
		method string
		want   want
	}{
		{
			name:   "positive postgres test #1",
			url:    "/postgres/users",
			json:   `{"name": "Alex","age": 21}`,
			method: "POST",
			want: want{
				code: http.StatusOK,
				json: `{"id": 1}`,
			},
		},
		{
			name:   "positive postgres test #2",
			url:    "/postgres/users",
			json:   `{"name": "Dima","age": 23}`,
			method: "POST",
			want: want{
				code: http.StatusOK,
				json: `{"id":2}`,
			},
		},

		{
			name:   "negative postgres test #3",
			url:    "/postgres/users",
			json:   `{"name":"HeapInuse","agef":"23.33"}`,
			method: "POST",
			want: want{
				code: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := mocks.NewCacher(t)
			storage := mocks.NewStorer(t)
			logger := logging.CreateLogger()

			storage.
				On("Create", entity.User{Name: "Alex", Age: 21}).
				Return(entity.User{Id: 1}, nil).
				Maybe()

			storage.
				On("Create", entity.User{Name: "Dima", Age: 23}).
				Return(entity.User{Id: 2}, nil).
				Maybe()

			a := New(storage, cache, logger)

			r := func(u UserHandler) chi.Router {
				r := chi.NewRouter()
				r.Post("/postgres/users", u.CreateUser)
				return r
			}(a)
			ts := httptest.NewServer(r)
			defer ts.Close()

			statusCode, _ := testRequest(t, ts, tt.url, tt.method, strings.NewReader(tt.json))
			assert.Equal(t, tt.want.code, statusCode)
		})
	}
}

func TestUserHandler_PostSign(t *testing.T) {
	type want struct {
		code int
		json string
	}
	tests := []struct {
		name   string
		url    string
		json   string
		method string
		want   want
	}{
		{
			name:   "positive sign test #1",
			url:    "/sign/hmacsha512",
			json:   `{"text":"test","key": "test123"}`,
			method: "POST",
			want: want{
				code: http.StatusOK,
				json: `{"hex": "9e1cdea55a6add8dc6688fbabfd6dd28b1b7896fa39aa36a0bef8f5e6c06c680"}`,
			},
		},
		{
			name:   "positive sign test #2",
			url:    "/sign/hmacsha512",
			json:   `{"text":"test2","key": "1"}`,
			method: "POST",
			want: want{
				code: http.StatusOK,
				json: `{"hex":"807c2bced808c51dfeba34788bfacbcffcd827fd394a96473e0999b5d2f803e5"}`,
			},
		},

		{
			name:   "negative sign test #3",
			url:    "/sign/hmacsha512",
			json:   `{"name":"HeapInuse","agef":"23.33"}`,
			method: "POST",
			want: want{
				code: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := mocks.NewCacher(t)
			storage := mocks.NewStorer(t)
			logger := logging.CreateLogger()

			a := New(storage, cache, logger)

			r := func(u UserHandler) chi.Router {
				r := chi.NewRouter()
				r.Post("/sign/hmacsha512", u.PostSign)
				return r
			}(a)
			ts := httptest.NewServer(r)
			defer ts.Close()

			statusCode, _ := testRequest(t, ts, tt.url, tt.method, strings.NewReader(tt.json))
			assert.Equal(t, tt.want.code, statusCode)
		})
	}
}
