package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestLoginPost(t *testing.T) {
	type input struct {
		User string `json:"username"`
		Pass string `json:"password"`
	}

	srv := server{
		db: &memDB{
			portfolios: map[string]*Portfolio{
				"foo": &Portfolio{},
			},
		},
	}

	cases := []struct {
		name       string
		input      interface{}
		wantCode   int
		wantCookie bool
	}{
		{"no body", nil, 400, false},
		{"empty body", &input{}, 401, false},
		{"invalid body", false, 400, false},
		{"success", &input{User: "foo"}, 200, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var req *http.Request
			var err error
			if tc.input != nil {
				var body bytes.Buffer
				err := json.NewEncoder(&body).Encode(tc.input)
				if err != nil {
					t.Fatal(err)
				}
				req, err = http.NewRequest(http.MethodPost, "/login", &body)
			} else {
				req, err = http.NewRequest(http.MethodPost, "/login", nil)
			}

			if err != nil {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()
			srv.ServeHTTP(w, req)
			if w.Code != tc.wantCode {
				t.Errorf("wanted %v, got %v: %s", tc.wantCode, w.Code, w.Body)
			}

			haveCookie := false
			for _, c := range w.Result().Cookies() {
				if c.Name == AuthCookieName {
					haveCookie = true
					break
				}
			}

			if haveCookie != tc.wantCookie {
				t.Errorf("wanted cookie %v, got %v", tc.wantCookie, haveCookie)
			}
		})
	}
}

func TestLoginGet(t *testing.T) {
	srv := server{}
	req, err := http.NewRequest(http.MethodGet, "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("wanted %v, got %v: %s", http.StatusMethodNotAllowed, w.Code, w.Body)
	}

	for _, c := range w.Result().Cookies() {
		if c.Name == AuthCookieName {
			t.Error("auth cookie set on invalid method")
			break
		}
	}
}

func TestNoAuth(t *testing.T) {
	srv := server{}
	req, err := http.NewRequest(http.MethodGet, "/noexist", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Errorf("wanted %v, got %v: %s", http.StatusUnauthorized, w.Code, w.Body)
	}
}

func TestPortfolios(t *testing.T) {
	srv := server{
		db: &memDB{
			portfolios: map[string]*Portfolio{
				"": &Portfolio{},
			},
		},
	}
	req, err := http.NewRequest(http.MethodGet, "/portfolios", nil)
	if err != nil {
		t.Fatal(err)
	}
	authRequest(req, "")

	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	if w.Code != http.StatusForbidden {
		t.Errorf("wanted %v, got %v: %s", http.StatusForbidden, w.Code, w.Body)
	}
}

func TestPortfolio(t *testing.T) {
	testdata := map[string]*Portfolio{
		"pid":   &Portfolio{Name: "pid"},
		"wrong": &Portfolio{Name: "wrong"},
	}

	cases := []struct {
		name     string
		authid   string
		pid      string
		data     map[string]*Portfolio
		wantCode int
		wantData *Portfolio
	}{
		{"invalid pid", "invalid", "invalid", testdata, 404, nil},
		{"wrong pid", "wrong", "pid", testdata, 404, nil},
		{"success", "pid", "pid", testdata, 200, testdata["pid"]},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := server{db: &memDB{portfolios: tc.data}}
			req, err := http.NewRequest(http.MethodGet, "/portfolios/"+tc.pid, nil)
			if err != nil {
				t.Fatal(err)
			}
			authRequest(req, tc.authid)
			w := httptest.NewRecorder()
			srv.ServeHTTP(w, req)
			if w.Code != tc.wantCode {
				t.Errorf("wanted %v, got %v: %s", tc.wantCode, w.Code, w.Body)
			}

			body := w.Result().Body
			defer body.Close()

			got := &Portfolio{}
			err = json.NewDecoder(body).Decode(got)

			if tc.wantData == nil {
				if err == nil {
					t.Errorf("unexpected response: %v", got)
				}
			} else {
				if err != nil {
					t.Errorf("reading response: %v", err)
				}
				if !reflect.DeepEqual(got, tc.wantData) {
					t.Errorf("wanted %v, got %v", tc.wantData, got)
				}
			}
		})
	}
}

func TestAccounts(t *testing.T) {
	testdata := map[string]*Portfolio{
		"pid": &Portfolio{
			Name: "pid",
			Accounts: []*Account{
				&Account{Name: "acc"},
			},
		},
		"wrong": &Portfolio{
			Name:     "wrong",
			Accounts: []*Account{},
		},
	}

	cases := []struct {
		name     string
		authid   string
		pid      string
		data     map[string]*Portfolio
		wantCode int
		wantData []*Account
	}{
		{"invalid pid", "invalid", "invalid", testdata, 404, nil},
		{"wrong pid", "wrong", "pid", testdata, 404, nil},
		{"success", "pid", "pid", testdata, 200, testdata["pid"].Accounts},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := server{db: &memDB{portfolios: tc.data}}
			req, err := http.NewRequest(http.MethodGet, "/portfolios/"+tc.pid+"/accounts", nil)
			if err != nil {
				t.Fatal(err)
			}
			authRequest(req, tc.authid)
			w := httptest.NewRecorder()
			srv.ServeHTTP(w, req)
			if w.Code != tc.wantCode {
				t.Errorf("wanted %v, got %v: %s", tc.wantCode, w.Code, w.Body)
			}

			body := w.Result().Body
			defer body.Close()

			var got []*Account
			err = json.NewDecoder(body).Decode(&got)

			if tc.wantData == nil {
				if err == nil {
					t.Errorf("unexpected response: %v", got)
				}
			} else {
				if err != nil {
					t.Errorf("reading response: %v", err)
				}
				if !reflect.DeepEqual(got, tc.wantData) {
					t.Errorf("wanted %v, got %v", tc.wantData, got)
				}
			}
		})
	}
}

func TestAccount(t *testing.T) {
	testdata := map[string]*Portfolio{
		"pid": &Portfolio{
			Name: "pid",
			Accounts: []*Account{
				&Account{Name: "acc"},
			},
		},
		"wrong": &Portfolio{
			Name:     "wrong",
			Accounts: []*Account{},
		},
	}

	cases := []struct {
		name     string
		authid   string
		pid      string
		acc      string
		data     map[string]*Portfolio
		wantCode int
		wantData *Account
	}{
		{"invalid pid", "invalid", "invalid", "acc", testdata, 404, nil},
		{"wrong pid", "wrong", "pid", "acc", testdata, 404, nil},
		{"invalid acc", "pid", "pid", "invalid", testdata, 404, nil},
		{"success", "pid", "pid", "acc", testdata, 200, testdata["pid"].Accounts[0]},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := server{db: &memDB{portfolios: tc.data}}
			req, err := http.NewRequest(http.MethodGet, "/portfolios/"+tc.pid+"/accounts/"+tc.acc, nil)
			if err != nil {
				t.Fatal(err)
			}
			authRequest(req, tc.authid)
			w := httptest.NewRecorder()
			srv.ServeHTTP(w, req)
			if w.Code != tc.wantCode {
				t.Errorf("wanted %v, got %v: %s", tc.wantCode, w.Code, w.Body)
			}

			body := w.Result().Body
			defer body.Close()

			got := &Account{}
			err = json.NewDecoder(body).Decode(got)

			if tc.wantData == nil {
				if err == nil {
					t.Errorf("unexpected response: %v", got)
				}
			} else {
				if err != nil {
					t.Errorf("reading response: %v", err)
				}
				if !reflect.DeepEqual(got, tc.wantData) {
					t.Errorf("wanted %v, got %v", tc.wantData, got)
				}
			}
		})
	}
}

func TestTrades(t *testing.T) {
	testdata := map[string]*Portfolio{
		"pid": &Portfolio{
			Name: "pid",
			Accounts: []*Account{
				&Account{Name: "acc", Trades: []*Trade{
					&Trade{
						Status:   TradePending,
						FundName: "GBP",
						Amount:   100,
					},
				}},
			},
		},
		"wrong": &Portfolio{
			Name:     "wrong",
			Accounts: []*Account{},
		},
	}

	cases := []struct {
		name     string
		authid   string
		pid      string
		acc      string
		data     map[string]*Portfolio
		wantCode int
		wantData []*Trade
	}{
		{"invalid pid", "invalid", "invalid", "acc", testdata, 404, nil},
		{"wrong pid", "wrong", "pid", "acc", testdata, 404, nil},
		{"invalid acc", "pid", "pid", "invalid", testdata, 404, nil},
		{"success", "pid", "pid", "acc", testdata, 200, testdata["pid"].Accounts[0].Trades},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := server{db: &memDB{portfolios: tc.data}}
			req, err := http.NewRequest(http.MethodGet, "/portfolios/"+tc.pid+"/accounts/"+tc.acc+"/trades", nil)
			if err != nil {
				t.Fatal(err)
			}
			authRequest(req, tc.authid)
			w := httptest.NewRecorder()
			srv.ServeHTTP(w, req)
			if w.Code != tc.wantCode {
				t.Errorf("wanted %v, got %v: %s", tc.wantCode, w.Code, w.Body)
			}

			body := w.Result().Body
			defer body.Close()

			var got []*Trade
			err = json.NewDecoder(body).Decode(&got)

			if tc.wantData == nil {
				if err == nil {
					t.Errorf("unexpected response: %v", got)
				}
			} else {
				if err != nil {
					t.Errorf("reading response: %v", err)
				}
				if !reflect.DeepEqual(got, tc.wantData) {
					t.Errorf("wanted %v, got %v", tc.wantData, got)
				}
			}
		})
	}
}

func TestPlaceTrade(t *testing.T) {
	testdata := map[string]*Portfolio{
		"pid": &Portfolio{
			Name: "pid",
			Accounts: []*Account{
				&Account{Name: "acc"},
			},
		},
		"wrong": &Portfolio{
			Name:     "wrong",
			Accounts: []*Account{},
		},
	}

	cases := []struct {
		name     string
		authid   string
		pid      string
		acc      string
		data     map[string]*Portfolio
		input    interface{}
		wantCode int
	}{
		{"invalid pid", "invalid", "invalid", "acc", testdata, &Trade{}, 404},
		{"wrong pid", "wrong", "pid", "acc", testdata, &Trade{}, 404},
		{"invalid acc", "pid", "pid", "invalid", testdata, &Trade{}, 404},
		{"no body", "pid", "pid", "acc", testdata, nil, 400},
		{"invalid bodu", "pid", "pid", "acc", testdata, "foo", 400},
		{"success", "pid", "pid", "acc", testdata, &Trade{}, 201},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			url := "/portfolios/" + tc.pid + "/accounts/" + tc.acc + "/trades"
			srv := server{db: &memDB{portfolios: tc.data}}
			var req *http.Request
			var err error
			if tc.input != nil {
				var body bytes.Buffer
				err := json.NewEncoder(&body).Encode(tc.input)
				if err != nil {
					t.Fatal(err)
				}
				req, err = http.NewRequest(http.MethodPost, url, &body)
			} else {
				req, err = http.NewRequest(http.MethodPost, url, nil)
			}

			if err != nil {
				t.Fatal(err)
			}
			authRequest(req, tc.authid)
			w := httptest.NewRecorder()
			srv.ServeHTTP(w, req)
			if w.Code != tc.wantCode {
				t.Errorf("wanted %v, got %v: %s", tc.wantCode, w.Code, w.Body)
			}
		})
	}
}

func TestFunds(t *testing.T) {
	funds := map[string]*Fund{
		"f": &Fund{Name: "f"},
	}
	var want []*Fund
	for _, f := range funds {
		want = append(want, f)
	}

	srv := server{db: &memDB{funds: funds}}
	req, err := http.NewRequest(http.MethodGet, "/funds", nil)
	if err != nil {
		t.Fatal(err)
	}
	authRequest(req, "")

	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("wanted %v, got %v: %s", http.StatusOK, w.Code, w.Body)
	}

	var got []*Fund
	err = json.NewDecoder(w.Result().Body).Decode(&got)
	if err != nil {
		t.Errorf("reading response: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestFund(t *testing.T) {
	testdata := map[string]*Fund{"f": &Fund{Name: "f"}}

	cases := []struct {
		name     string
		fid      string
		data     map[string]*Fund
		wantCode int
		wantData *Fund
	}{
		{"wrong fid", "wrong", testdata, 404, nil},
		{"success", "f", testdata, 200, testdata["f"]},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := server{db: &memDB{funds: tc.data}}
			req, err := http.NewRequest(http.MethodGet, "/funds/"+tc.fid, nil)
			if err != nil {
				t.Fatal(err)
			}
			authRequest(req, "")
			w := httptest.NewRecorder()
			srv.ServeHTTP(w, req)
			if w.Code != tc.wantCode {
				t.Errorf("wanted %v, got %v: %s", tc.wantCode, w.Code, w.Body)
			}

			body := w.Result().Body
			defer body.Close()

			got := &Fund{}
			err = json.NewDecoder(body).Decode(got)

			if tc.wantData == nil {
				if err == nil {
					t.Errorf("unexpected response: %v", got)
				}
			} else {
				if err != nil {
					t.Errorf("reading response: %v", err)
				}
				if !reflect.DeepEqual(got, tc.wantData) {
					t.Errorf("wanted %v, got %v", tc.wantData, got)
				}
			}
		})
	}
}

func authRequest(req *http.Request, pid string) {
	req.AddCookie(&http.Cookie{
		Name:    AuthCookieName,
		Value:   pid,
		Expires: time.Now().UTC().Add(time.Hour),
	})
}
