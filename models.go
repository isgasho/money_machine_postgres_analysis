package main

type fn func(params ...interface{})

// var cyclePool = map[]Cycle{}

var cycleMapPool = make(map[string]*Cycle)

// var dayID = 3

//User object for DB
type User struct {
	ID          int
	Age         int
	FirstName   string
	LastName    string
	Email       string
	RequestType string
}

type DatabaseQuery struct {
	RequestType string
	Data        []string
	Range1      string
	Range2      string
}

type DatabaseResponse struct {
	Data      string
	StockList []Stock
}

type DatabaseMonitorSymbolListResponse struct {
	MonitorSymbolList []string
}

type DatabaseDowListResponse struct {
	DowList []Dow
}

type DatabaseStockListResponse struct {
	StockList []Stock
}
type DatabaseMetricsWhaleResponse struct {
	MetricsWhale []MetricsWhale
}

// type DatabaseMetricsResponseWhale struct {
// 	Metrics w
// }

type Day struct {
	ID        string
	DayOfWeek string
	CreatedAt string
}

type News struct {
	ID        string
	DayID     int
	CreatedAt string
	NewsInfo  string
}

type Dow struct {
	ID               string
	CreatedAt        string
	CurrentDowValue  string
	PointsChanged    string
	PercentageChange string
}

type Stock struct {
	ID           string
	DayID        int
	CreatedAt    string
	Rank         string
	UserInputed  bool
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
	Lo           string
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
	DayID             int
	CreatedAt         string
	BoughtPrice       string
	SoldPrice         string
	BuyCompletedTime  string
	SoldCompletedTime string
	TradeSuccessful   bool
	BuyPerformed      bool
	SoldPerformed     bool
}

//Cycle management struct
type Cycle struct {
	Name             string
	CreationIndex    int
	BooleanOperate   bool
	IntervalSpeed    int
	AmountOfInterval int
	FunctionToCall   fn
	Params           interface{}
}

type BrokerageQuery struct {
	request_type string
	Name         string
}

type MetricsWhale struct {
	CreatedAt                      string
	DesiredPriceRangeHigh          string
	DesiredPriceRangeLow           string
	DesiredPchg                    string
	DesiredPchgVarianceValue       string
	DesiredVolatilityVarianceValue string
}
type EvalResultsWhale struct {
	ID             int
	Symbol         string
	IsBreachWorthy string
	IsPatternMet   string
}
