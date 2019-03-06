package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"
)

const (
	// AuthCookieName is the cookie that holds the login token.
	AuthCookieName = "ADALPHA-TEST-AUTH"
)

type server struct {
	port int
	db   Repository
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = shiftPath(req.URL.Path)

	if head == "login" {
		s.Login(w, req)
		return
	}

	if !s.Auth(w, req) {
		http.Error(w, "call /login first", http.StatusUnauthorized)
		return
	}

	switch head {
	case "portfolios":
		s.Portfolios(w, req)
		return
	case "funds":
		s.Funds(w, req)
		return
	}

	http.NotFound(w, req)
}

func (s *server) Login(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	form := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if req.Body == nil {
		http.Error(w, "missing request body", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(req.Body).Decode(&form)
	defer req.Body.Close()

	if err != nil {
		http.Error(w, "request should be json data with username and password fields", http.StatusBadRequest)
		return
	}

	if p, err := s.db.PortfolioByName(form.Username); err != nil || p == nil {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}

	// Clearly this is a terrible mechanism for authentication, not for production use.
	http.SetCookie(w, &http.Cookie{
		Name:    AuthCookieName,
		Value:   form.Username,
		Expires: time.Now().UTC().Add(time.Hour),
	})
}

func (s *server) Auth(w http.ResponseWriter, req *http.Request) bool {
	// Again, this is not real security!
	c, err := req.Cookie(AuthCookieName)
	if err != nil {
		return false
	}
	return c != nil
}

func (s *server) Portfolios(w http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = shiftPath(req.URL.Path)

	if head != "" {
		s.Portfolio(w, req, head)
		return
	}

	if req.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.Error(w, "portfolio listing denied", http.StatusForbidden)
}

func (s *server) Portfolio(w http.ResponseWriter, req *http.Request, portfolioName string) {
	authedName, err := authPortfolioName(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if authedName != portfolioName {
		http.Error(w, portfolioName+" not found", http.StatusNotFound)
		return
	}

	p, err := s.db.PortfolioByName(portfolioName)
	if err != nil {
		http.Error(w, fmt.Sprintf("loading portfolio: %v", err), http.StatusInternalServerError)
		return
	}
	if p == nil {
		http.Error(w, portfolioName+" not found", http.StatusNotFound)
		return
	}

	var head string
	head, req.URL.Path = shiftPath(req.URL.Path)

	if head != "" {
		if head == "accounts" {
			s.Accounts(w, req, p)
			return
		}
		http.NotFound(w, req)
		return
	}

	if req.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sendJSON(w, p)
}

func (s *server) Accounts(w http.ResponseWriter, req *http.Request, p *Portfolio) {
	var head string
	head, req.URL.Path = shiftPath(req.URL.Path)

	if head != "" {
		s.Account(w, req, p, head)
		return
	}

	if req.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sendJSON(w, p.Accounts)
}

func (s *server) Account(w http.ResponseWriter, req *http.Request, p *Portfolio, accName string) {
	for _, acc := range p.Accounts {
		if acc.Name == accName {
			var head string
			head, req.URL.Path = shiftPath(req.URL.Path)

			if head != "" {
				if head == "trades" {
					s.Trades(w, req, p.Name, acc)
					return
				}
				http.NotFound(w, req)
				return
			}

			if req.Method != http.MethodGet {
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
				return
			}

			sendJSON(w, acc)
			return
		}
	}

	http.Error(w, accName+" not found", http.StatusNotFound)
}

func (s *server) Trades(w http.ResponseWriter, req *http.Request, portfolioName string, account *Account) {
	if req.Method == http.MethodPost {
		s.PlaceTrade(w, req, portfolioName, account.Name)
		return
	}

	if req.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sendJSON(w, account.Trades)
}

func (s *server) PlaceTrade(w http.ResponseWriter, req *http.Request, portfolioName, accountName string) {
	if req.Body == nil {
		http.Error(w, "missing request body", http.StatusBadRequest)
		return
	}

	trade := &Trade{}
	err := json.NewDecoder(req.Body).Decode(&trade)
	defer req.Body.Close()

	if err != nil {
		http.Error(w, "unable to parse request", http.StatusBadRequest)
		return
	}

	err = s.db.SaveTrade(portfolioName, accountName, trade)
	if err != nil {
		http.Error(w, fmt.Sprintf("saving trade: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *server) Funds(w http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = shiftPath(req.URL.Path)

	if head != "" {
		s.Fund(w, req, head)
		return
	}

	if req.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	funds, err := s.db.Funds()
	if err != nil {
		http.Error(w, fmt.Sprintf("loading funds: %v", err), http.StatusInternalServerError)
		return
	}

	type summary struct {
		ISIN  string `json:"isin"`
		Name  string `json:"name"`
		Price int    `json:"price"`
	}

	var response []summary
	for _, f := range funds {
		response = append(response, summary{ISIN: f.ISIN, Name: f.Name, Price: f.CurrentPrice})
	}

	sendJSON(w, response)
}

func (s *server) Fund(w http.ResponseWriter, req *http.Request, name string) {
	if req.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	f, err := s.db.FundByName(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("loading fund %s: %v", name, err), http.StatusInternalServerError)
		return
	}

	if f == nil {
		http.Error(w, name+" not found", http.StatusNotFound)
		return
	}

	sendJSON(w, f)
}

func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

func authPortfolioName(req *http.Request) (string, error) {
	c, err := req.Cookie(AuthCookieName)
	if err != nil {
		return "", fmt.Errorf("getting portfolio name: %v", err)
	}
	return c.Value, nil
}

func sendJSON(w http.ResponseWriter, p interface{}) {
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		http.Error(w, "marshalling json response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}
