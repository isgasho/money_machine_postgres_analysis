package main

//User object for DB
type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

type Day struct {
	ID        string
	DayOfWeek string
}

type News struct {
	ID        string
	DayID     string
	CreatedAt string
	NewsInfo  string
}

type Dow struct {
	ID        string
	DayID     string
	CreatedAt string
	DowInfo   string
}

type Stock struct {
	ID           string
	DayID        string
	CreatedAt    string
	Monitoring   bool
	Symbol       string
	Bid          string
	Ask          string
	Last         string
	Pchg         string
	Pcls         string
	Opn          string
	Vl           string
	Pvol         string
	Volatility12 string
	Wk52hi       string
	Wk52hidate   string
	Wk52lo       string
	Wk52lodate   string
	Hi           string
	Low          string
	PrAdp50      string
	PrAdp100     string
	Prchg        string
	Adp50        string
	Adp100       string
	Adv30        string
	Adv90        string
}

type TradeInfo struct {
	ID                string
	DayID             string
	CreatedAt         string
	BoughtPrice       string
	SoldPrice         string
	BuyCompletedTime  string
	SoldCompletedTime string
	TradeSuccessful   bool
	BuyPerformed      bool
	SoldPerformed     bool
}
