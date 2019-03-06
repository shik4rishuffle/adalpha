package main

import (
	"fmt"
	"time"
)

// Repository provides access to a persistence implementation.
type Repository interface {
	PortfolioByName(name string) (*Portfolio, error)
	Funds() ([]*Fund, error)
	FundByName(name string) (*Fund, error)
	SaveTrade(portfolio, account string, trade *Trade) error
}

type memDB struct {
	funds      map[string]*Fund
	portfolios map[string]*Portfolio
}

func (db *memDB) Funds() ([]*Fund, error) {
	var funds []*Fund
	for _, f := range db.funds {
		funds = append(funds, f)
	}
	return funds, nil
}

func (db *memDB) FundByName(name string) (*Fund, error) {
	f, found := db.funds[name]
	if !found {
		return nil, nil
	}
	return f, nil
}

func (db *memDB) PortfolioByName(name string) (*Portfolio, error) {
	p, found := db.portfolios[name]
	if !found {
		return nil, nil
	}
	return p, nil
}

func (db *memDB) SaveTrade(portfolio, account string, trade *Trade) error {
	p, found := db.portfolios[portfolio]

	if !found {
		return fmt.Errorf("portfolio not found")
	}

	for _, acc := range p.Accounts {
		if acc.Name == account {
			acc.Trades = append(acc.Trades, trade)
			return nil
		}
	}

	return fmt.Errorf("account not found")
}

func testData() *memDB {
	return &memDB{
		funds: map[string]*Fund{
			"IE00B52L4369": &Fund{ISIN: "IE00B52L4369", Name: "BlackRock Institutional Cash Series Sterling Liquidity Agency Inc", CurrentPrice: 100, History: historic["IE00B52L4369"]},
			"GB00BQ1YHQ70": &Fund{ISIN: "GB00BQ1YHQ70", Name: "Threadneedle UK Property Authorised Investment Net GBP 1 Acc", CurrentPrice: 116, History: historic["GB00BQ1YHQ70"]},
			"GB00B3X7QG63": &Fund{ISIN: "GB00B3X7QG63", Name: "Vanguard FTSE U.K. All Share Index Unit Trust Accumulation", CurrentPrice: 19830, History: historic["GB00B3X7QG63"]},
			"GB00BG0QP828": &Fund{ISIN: "GB00BG0QP828", Name: "Legal & General Japan Index Trust C Class Accumulation", CurrentPrice: 6187, History: historic["GB00BG0QP828"]},
			"GB00BPN5P238": &Fund{ISIN: "GB00BPN5P238", Name: "Vanguard US Equity Index Institutional Plus GBP Accumulation", CurrentPrice: 18379, History: historic["GB00BPN5P238"]},
			"IE00B1S74Q32": &Fund{ISIN: "IE00B1S74Q32", Name: "Vanguard U.K. Investment Grade Bond Index Fund GBP Accumulation", CurrentPrice: 9550, History: historic["IE00B1S74Q32"]},
		},
		portfolios: map[string]*Portfolio{
			"ADA123456789": &Portfolio{
				Name: "ADA123456789",
				Accounts: []*Account{
					&Account{
						Name: "ADA123456789-ISA",
						Holdings: map[string]Decimal{
							"IE00B52L4369": NewDecimal(44000, 0),
							"GB00BQ1YHQ70": NewDecimal(3793103448, -5),
							"GB00B3X7QG63": NewDecimal(1109430156, -7),
							"GB00BG0QP828": NewDecimal(1777921448, -7),
							"GB00BPN5P238": NewDecimal(1795527504, -7),
							"IE00B1S74Q32": NewDecimal(6910994764, -7),
						},
					},
				},
			},
		},
	}
}

func parseISO(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		panic(fmt.Sprintf("bad test date %s: %v", s, err))
	}
	return t
}

var historic = map[string][]FundHistory{
	"IE00B52L4369": []FundHistory{
		{Date: parseISO("2019-03-01"), Price: 100},
		{Date: parseISO("2019-02-28"), Price: 100},
		{Date: parseISO("2019-02-27"), Price: 100},
		{Date: parseISO("2019-02-26"), Price: 100},
		{Date: parseISO("2019-02-25"), Price: 100},
		{Date: parseISO("2019-02-24"), Price: 100},
		{Date: parseISO("2019-02-23"), Price: 100},
		{Date: parseISO("2019-02-22"), Price: 100},
		{Date: parseISO("2019-02-21"), Price: 100},
		{Date: parseISO("2019-02-20"), Price: 100},
		{Date: parseISO("2019-02-19"), Price: 100},
		{Date: parseISO("2019-02-18"), Price: 100},
		{Date: parseISO("2019-02-17"), Price: 100},
		{Date: parseISO("2019-02-16"), Price: 100},
		{Date: parseISO("2019-02-15"), Price: 100},
		{Date: parseISO("2019-02-14"), Price: 100},
		{Date: parseISO("2019-02-13"), Price: 100},
		{Date: parseISO("2019-02-12"), Price: 100},
		{Date: parseISO("2019-02-11"), Price: 100},
		{Date: parseISO("2019-02-10"), Price: 100},
		{Date: parseISO("2019-02-09"), Price: 100},
		{Date: parseISO("2019-02-08"), Price: 100},
		{Date: parseISO("2019-02-07"), Price: 100},
		{Date: parseISO("2019-02-06"), Price: 100},
		{Date: parseISO("2019-02-05"), Price: 100},
		{Date: parseISO("2019-02-04"), Price: 100},
		{Date: parseISO("2019-02-03"), Price: 100},
		{Date: parseISO("2019-02-02"), Price: 100},
		{Date: parseISO("2019-02-01"), Price: 100},
		{Date: parseISO("2019-01-31"), Price: 100},
		{Date: parseISO("2019-01-30"), Price: 100},
	},
	"GB00BQ1YHQ70": []FundHistory{
		{Date: parseISO("2019-03-01"), Price: 116},
		{Date: parseISO("2019-02-28"), Price: 116},
		{Date: parseISO("2019-02-27"), Price: 116},
		{Date: parseISO("2019-02-26"), Price: 116},
		{Date: parseISO("2019-02-25"), Price: 116},
		{Date: parseISO("2019-02-22"), Price: 116},
		{Date: parseISO("2019-02-21"), Price: 116},
		{Date: parseISO("2019-02-20"), Price: 116},
		{Date: parseISO("2019-02-19"), Price: 116},
		{Date: parseISO("2019-02-18"), Price: 116},
		{Date: parseISO("2019-02-15"), Price: 116},
		{Date: parseISO("2019-02-14"), Price: 116},
		{Date: parseISO("2019-02-13"), Price: 116},
		{Date: parseISO("2019-02-12"), Price: 116},
		{Date: parseISO("2019-02-11"), Price: 116},
		{Date: parseISO("2019-02-08"), Price: 116},
		{Date: parseISO("2019-02-07"), Price: 116},
		{Date: parseISO("2019-02-06"), Price: 116},
		{Date: parseISO("2019-02-05"), Price: 116},
		{Date: parseISO("2019-02-04"), Price: 116},
		{Date: parseISO("2019-02-01"), Price: 116},
		{Date: parseISO("2019-01-31"), Price: 116},
		{Date: parseISO("2019-01-30"), Price: 116},
	},
	"GB00B3X7QG63": []FundHistory{
		{Date: parseISO("2019-03-01"), Price: 19830},
		{Date: parseISO("2019-02-28"), Price: 19717},
		{Date: parseISO("2019-02-27"), Price: 19663},
		{Date: parseISO("2019-02-26"), Price: 19873},
		{Date: parseISO("2019-02-25"), Price: 19944},
		{Date: parseISO("2019-02-22"), Price: 19843},
		{Date: parseISO("2019-02-21"), Price: 19903},
		{Date: parseISO("2019-02-20"), Price: 19990},
		{Date: parseISO("2019-02-19"), Price: 19765},
		{Date: parseISO("2019-02-18"), Price: 19959},
		{Date: parseISO("2019-02-15"), Price: 19973},
		{Date: parseISO("2019-02-14"), Price: 19777},
		{Date: parseISO("2019-02-13"), Price: 19730},
		{Date: parseISO("2019-02-12"), Price: 19667},
		{Date: parseISO("2019-02-11"), Price: 19566},
		{Date: parseISO("2019-02-08"), Price: 19497},
		{Date: parseISO("2019-02-07"), Price: 19573},
		{Date: parseISO("2019-02-06"), Price: 19702},
		{Date: parseISO("2019-02-05"), Price: 19794},
		{Date: parseISO("2019-02-04"), Price: 19447},
		{Date: parseISO("2019-02-01"), Price: 19409},
		{Date: parseISO("2019-01-31"), Price: 19277},
		{Date: parseISO("2019-01-30"), Price: 19140},
	},
	"GB00BG0QP828": []FundHistory{
		{Date: parseISO("2019-03-01"), Price: 6187},
		{Date: parseISO("2019-02-28"), Price: 6170},
		{Date: parseISO("2019-02-27"), Price: 6223},
		{Date: parseISO("2019-02-26"), Price: 6252},
		{Date: parseISO("2019-02-25"), Price: 6335},
		{Date: parseISO("2019-02-22"), Price: 6320},
		{Date: parseISO("2019-02-21"), Price: 6319},
		{Date: parseISO("2019-02-20"), Price: 6330},
		{Date: parseISO("2019-02-19"), Price: 6356},
		{Date: parseISO("2019-02-18"), Price: 6345},
		{Date: parseISO("2019-02-15"), Price: 6296},
		{Date: parseISO("2019-02-14"), Price: 6317},
		{Date: parseISO("2019-02-13"), Price: 6287},
		{Date: parseISO("2019-02-12"), Price: 6250},
		{Date: parseISO("2019-02-11"), Price: 6126},
		{Date: parseISO("2019-02-08"), Price: 6121},
		{Date: parseISO("2019-02-07"), Price: 6285},
		{Date: parseISO("2019-02-06"), Price: 6303},
		{Date: parseISO("2019-02-05"), Price: 6262},
		{Date: parseISO("2019-02-04"), Price: 6240},
		{Date: parseISO("2019-02-01"), Price: 6238},
		{Date: parseISO("2019-01-31"), Price: 6233},
		{Date: parseISO("2019-01-30"), Price: 6141},
	},
	"GB00BPN5P238": []FundHistory{
		{Date: parseISO("2019-02-28"), Price: 18379},
		{Date: parseISO("2019-02-27"), Price: 18382},
		{Date: parseISO("2019-02-26"), Price: 18509},
		{Date: parseISO("2019-02-25"), Price: 18783},
		{Date: parseISO("2019-02-22"), Price: 18774},
		{Date: parseISO("2019-02-21"), Price: 18631},
		{Date: parseISO("2019-02-20"), Price: 18701},
		{Date: parseISO("2019-02-19"), Price: 18711},
		{Date: parseISO("2019-02-18"), Price: 18831},
		{Date: parseISO("2019-02-15"), Price: 18948},
		{Date: parseISO("2019-02-14"), Price: 18837},
		{Date: parseISO("2019-02-13"), Price: 18723},
		{Date: parseISO("2019-02-12"), Price: 18664},
		{Date: parseISO("2019-02-11"), Price: 18455},
		{Date: parseISO("2019-02-08"), Price: 18313},
		{Date: parseISO("2019-02-07"), Price: 18254},
		{Date: parseISO("2019-02-06"), Price: 18428},
		{Date: parseISO("2019-02-05"), Price: 18501},
		{Date: parseISO("2019-02-04"), Price: 18218},
		{Date: parseISO("2019-02-01"), Price: 18074},
		{Date: parseISO("2019-01-31"), Price: 17960},
		{Date: parseISO("2019-01-30"), Price: 17918},
	},
	"IE00B1S74Q32": []FundHistory{
		{Date: parseISO("2019-03-01"), Price: 9550},
		{Date: parseISO("2019-02-28"), Price: 9538},
		{Date: parseISO("2019-02-27"), Price: 9484},
		{Date: parseISO("2019-02-26"), Price: 9577},
		{Date: parseISO("2019-02-25"), Price: 9591},
		{Date: parseISO("2019-02-22"), Price: 9600},
		{Date: parseISO("2019-02-21"), Price: 9571},
		{Date: parseISO("2019-02-20"), Price: 9583},
		{Date: parseISO("2019-02-19"), Price: 9529},
		{Date: parseISO("2019-02-18"), Price: 9588},
		{Date: parseISO("2019-02-15"), Price: 9583},
		{Date: parseISO("2019-02-14"), Price: 9598},
		{Date: parseISO("2019-02-13"), Price: 9521},
		{Date: parseISO("2019-02-12"), Price: 9572},
		{Date: parseISO("2019-02-11"), Price: 9579},
		{Date: parseISO("2019-02-08"), Price: 9590},
		{Date: parseISO("2019-02-07"), Price: 9575},
		{Date: parseISO("2019-02-06"), Price: 9556},
		{Date: parseISO("2019-02-05"), Price: 9545},
		{Date: parseISO("2019-02-04"), Price: 9508},
		{Date: parseISO("2019-02-01"), Price: 9529},
		{Date: parseISO("2019-01-31"), Price: 9539},
		{Date: parseISO("2019-01-30"), Price: 9454},
	},
}
