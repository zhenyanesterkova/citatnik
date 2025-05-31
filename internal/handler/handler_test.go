package handler

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/zhenyanesterkova/citatnik/internal/app/apperrors"
	"github.com/zhenyanesterkova/citatnik/internal/app/quote"
)

type testServer struct {
	s *httptest.Server
}

func (ts *testServer) testRequest(t *testing.T, method, path string, reqBody string) *http.Response {
	t.Helper()

	var buff bytes.Buffer
	_, _ = buff.WriteString(reqBody)

	req, _ := http.NewRequest(method, ts.s.URL+path, &buff)

	resp, _ := ts.s.Client().Do(req)

	return resp
}

func TestRouter(t *testing.T) {
	store := &testMemory{}

	router := mux.NewRouter()

	rHandler := NewRepositorieHandler(store)
	rHandler.InitRouter(router)

	server := httptest.NewServer(router)

	testServer := testServer{
		s: server,
	}

	tests := []struct {
		name            string
		method          string
		url             string
		reqBody         string
		wantRespBody    string
		wantStatus      int
		wantErr         bool
		lenForGetRandom bool
	}{
		{
			name:       "Ping_success",
			method:     http.MethodGet,
			url:        "/ping",
			wantStatus: http.StatusOK,
		},
		{
			name:         "Ping_error",
			method:       http.MethodGet,
			url:          "/ping",
			wantStatus:   http.StatusInternalServerError,
			wantErr:      true,
			wantRespBody: TextServerError + "\n",
		},
		{
			name:         "Get_all_success",
			method:       http.MethodGet,
			url:          "/quotes",
			wantRespBody: "[{\"id\":1,\"author\":\"test\",\"quote\":\"text from test\",\"created_at\":\"2013-02-03T00:00:00Z\",\"updated_at\":\"2013-02-03T00:00:00Z\"}]\n",
			wantStatus:   http.StatusOK,
		},
		{
			name:         "Get_author_success",
			method:       http.MethodGet,
			url:          "/quotes?author=test",
			wantRespBody: "[{\"id\":1,\"author\":\"test\",\"quote\":\"text from test\",\"created_at\":\"2013-02-03T00:00:00Z\",\"updated_at\":\"2013-02-03T00:00:00Z\"}]\n",
			wantStatus:   http.StatusOK,
		},
		{
			name:            "Get_random_success",
			method:          http.MethodGet,
			url:             "/quotes/random",
			wantRespBody:    "{\"id\":1,\"author\":\"test\",\"quote\":\"text from test\",\"created_at\":\"2013-02-03T00:00:00Z\",\"updated_at\":\"2013-02-03T00:00:00Z\"}\n",
			wantStatus:      http.StatusOK,
			lenForGetRandom: true,
		},
		{
			name:       "Add_success",
			method:     http.MethodPost,
			url:        "/quotes",
			reqBody:    "{\"author\":\"test\", \"text\":\"text from test\"}",
			wantStatus: http.StatusOK,
		},
		{
			name:         "Add_err_decode",
			method:       http.MethodPost,
			url:          "/quotes",
			reqBody:      "{\"author:\"test\", \"text\":\"text from test\"}",
			wantStatus:   http.StatusInternalServerError,
			wantRespBody: TextServerError + "\n",
		},
		{
			name:         "Add_err",
			method:       http.MethodPost,
			url:          "/quotes",
			reqBody:      "{\"author\":\"test\", \"text\":\"text from test\"}",
			wantStatus:   http.StatusInternalServerError,
			wantRespBody: TextServerError + "\n",
			wantErr:      true,
		},
		{
			name:       "Delete_success",
			method:     http.MethodDelete,
			url:        "/quotes/1",
			wantStatus: http.StatusOK,
		},
		{
			name:         "Delete_parse_id_err",
			method:       http.MethodDelete,
			url:          "/quotes/ttt",
			wantStatus:   http.StatusBadRequest,
			wantRespBody: TextDeleteBadID + "\n",
		},
		{
			name:         "Delete_parse_id_err",
			method:       http.MethodDelete,
			url:          "/quotes/3",
			wantStatus:   http.StatusBadRequest,
			wantRespBody: TextDeleteNotFound + "\n",
			wantErr:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store.err = test.wantErr
			store.len = test.lenForGetRandom

			resp := testServer.testRequest(t, test.method, test.url, test.reqBody)
			respBody, _ := io.ReadAll(resp.Body)
			body := string(respBody)

			_ = resp.Body.Close()

			if resp.StatusCode != test.wantStatus {
				t.Errorf("StatusCode = %v, want %v", resp.StatusCode, test.wantStatus)
			}

			if body != test.wantRespBody {
				t.Errorf("Body = %v, want %v", body, test.wantRespBody)
			}
		})
	}
}

var testTime, _ = time.Parse("2006-Jan-02", "2013-Feb-03")

type testMemory struct {
	err bool
	len bool
}

func (m *testMemory) Add(quote *quote.Quote) error {
	if m.err {
		return errors.New("failed add quote to memory: ID 3 already exists")
	}
	return nil
}

func (m *testMemory) GetAll() []*quote.Quote {
	return []*quote.Quote{
		&quote.Quote{
			ID:        1,
			Author:    "test",
			Text:      "text from test",
			CreatedAt: testTime,
			UpdatedAt: testTime,
		},
	}
}

func (m *testMemory) GetRandom() *quote.Quote {
	if m.len {
		return &quote.Quote{
			ID:        1,
			Author:    "test",
			Text:      "text from test",
			CreatedAt: testTime,
			UpdatedAt: testTime,
		}
	}
	return nil
}

func (m *testMemory) GetByAuthor(author string) []*quote.Quote {
	return []*quote.Quote{
		&quote.Quote{
			ID:        1,
			Author:    "test",
			Text:      "text from test",
			CreatedAt: testTime,
			UpdatedAt: testTime,
		},
	}
}

func (m *testMemory) Delete(id uint64) error {
	if m.err {
		return fmt.Errorf("failed delete quote with ID 3: %w", apperrors.ErrDeleteNotFound)
	}

	return nil
}

func (m *testMemory) Ping() error {
	if m.err {
		return errors.New("test err")
	}
	return nil
}
