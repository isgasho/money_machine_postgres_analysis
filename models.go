package main

type fn func(params ...interface{})

// var cyclePool = map[]Cycle{}

var cycleMapPool = make(map[string]*Cycle)

var globalTSPCollectionStatementCache = []string{}

var globalCacheStockWisemenTopTier = []string{}

var intervalTradeMonitorDelimiter = 0

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
type DatabaseMetricsWisemenResponse struct {
	MetricsWisemen []MetricsWisemen
}

type DatabaseOrderInformationWisemenResponse struct {
	OrderInformationWisemen []OrderInformationWisemen
}
type DatabaseAltIntervalBuyWisemen struct {
	AltIntervalBuyWisemen []AltIntervalBuyWisemen
}

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
	ID              string
	CreatedAt       string
	CurrentDowValue string
}

type Stock struct {
	ID                                    string
	DayID                                 int
	CreatedAt                             string
	Rank                                  string
	UserInputed                           bool
	Monitoring                            bool
	Symbol                                string
	Bid                                   string
	Ask                                   string
	Last                                  string
	Pchg                                  string
	Pcls                                  string
	Opn                                   string
	Vl                                    string
	TimeCreated                           string
	Pvol                                  string
	Volatility12                          string
	Wk52hi                                string
	Wk52hidate                            string
	Wk52lo                                string
	Wk52lodate                            string
	Hi                                    string
	Lo                                    string
	PrAdp50                               string
	PrAdp100                              string
	Prchg                                 string
	Adp50                                 string
	Adp100                                string
	Adv30                                 string
	Adv90                                 string
	IsCurrentPriceHigherThanPreviousClose string
}

type TradeBoughtEvaluation struct {
	IsBought    bool
	HoldingList []Holding
}
type Holding struct {
	Symbol        string
	PurchasePrice string
	Qty           string
}
type HoldingWisemen struct {
	CreatedAt   string
	Symbol      string
	Price       string
	Qty         string
	OrderStatus string
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
type AccountBrokerage struct {
	Total          string
	CashAvailable  string
	UnsettledFunds string
}

//down_day_evaluation
type DownDayEvaluation struct {
	CreatedAt    string
	IsDownDay    string
	Dow          string
	PreviousDow  string
	GreatestPchg string
}

//cash_day_evaluation
type CashDayEvaluation struct {
	CreatedAt            string
	IsUnsettledFunds     string
	AmountUnsettledFunds string
}

// CREATE TABLE alt_interval_buy_wisemen
// (
//    id SERIAL PRIMARY KEY,
//    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
//    symbol VARCHAR,
//    is_alt_interval_operation VARCHAR
// );

type AltIntervalBuyWisemen struct {
	CreatedAt              string
	Symbol                 string
	IsAltIntervalOperation string
	ReasonCancelation      string
}

type MetricsWhale struct {
	CreatedAt                      string
	DesiredPriceRangeHigh          string
	DesiredPriceRangeLow           string
	DesiredPchg                    string
	DesiredPchgVarianceValue       string
	DesiredVolatilityVarianceValue string
}
type MetricsWisemen struct {
	CreatedAt                      string
	DesiredPriceRangeHigh          string
	DesiredPriceRangeLow           string
	PriceHighPchgAlgoDecision      string
	PriceLowPchgAlgoDecision       string
	PriceHighPchgTrade             string
	PriceLowPchgTrade              string
	DesiredPchgVarianceValue       string
	DesiredVolatilityVarianceValue string
	SellTime                       string
}
type OrderInformationWisemen struct {
	CreatedAt string
	IsBought  string
	Symbol    string
}

type EvalResultsWhale struct {
	ID             int
	Symbol         string
	IsBreachWorthy string
	IsPatternMet   string
}

type DayTrackingRecord struct {
	CreatedAt              string
	Symbol                 string
	DayOfWeekCreated       string
	DayOfWeekDayIteration  string
	LastDayOfWeekDayUpdate string
	AmountOfTrades         string
	IsWeekPassed           string
}

type ContainerOrders struct {
	ListOrders []Order
}
type Order struct {
	Symbol      string
	SVI         string
	OrderStatus string
	Qty         string
}
type ContainerHolding struct {
	ListHolding []HoldingWisemen
}

type InformationAtTrade struct {
	CreatedAt string
	Symbol    string
	TypeTrade string
	Side      string
	Qty       string
	Year      string
	Month     string
	Day       string
	Hour      string
	Minute    string
	Dow       string
	Bid       string
	Ask       string
	Last      string
}

type TradeConditionalMetrics struct {
	CreatedAt    string
	Symbol       string
	TimeStart    string
	TimeEnd      string
	PriceDropout string
}

type BuyStatusWisemen struct {
	CreatedAt  string
	Symbol     string
	IsHoldings bool
	QtyBought  string
}

type ContainerNumberRange struct {
	ListNumberRange []WebscrapeNumberRange
}
type WebscrapeNumberRange struct {
	NumberRange []string
	StringValue string
}
type WisemenMatchClosestToDelimiter struct {
	SplitStringValue      int
	DistanceFromDelimiter int
}

type AlgorithmEvaluationForDay struct {
	CreatedAt     string
	Name          string
	Symbol        string
	TimeStart     string
	TimeEnd       string
	IsCompleted   string
	IsProfitable  string
	BalanceBefore string
	BalanceAfter  string
}

type HistoryValue struct {
	Symbol            string
	Date              string
	Side              string
	Qty               string
	Price             string
	IntervalInList    string
	IsCalculationTrue string
}

type HistoryValueDay struct {
	ListHistoryValue []HistoryValue
}

type HistoryDayListContainer struct {
	HistoryValueDayList []HistoryValueDay
}

type TransactionHistory struct {
	Symbol           string
	Outcome          string
	HistoryValueList []HistoryValue
}

type TradeResultStore struct {
	CreatedAt               string
	AlgorithmUsed           string
	Result                  string
	BoughtPrice             string
	SellPrice               string
	ChangeAmount            string
	StockSymbol             string
	Qty                     string
	TimeStart               string
	TimeEnd                 string
	TimeTradeBuy            string
	TimeTradeSell           string
	HighestPricePointForDay string
	TimeHighestPricePoint   string
	LowestPricePointForDay  string
	TimeLowestPricePoint    string
	Dow1                    string
	Dow2                    string
	Dow3                    string
	Dow4                    string
	Dow5                    string
	Dow6                    string
}

type RecordSystemMonthContainer struct {
	RecordSystemMonthList []RecordSystemMonth
}

type RecordSystemMonth struct {
	IntMonthOfYear  int
	IntNumberOfDays int
}

type MarketOpenAnalysis struct {
	CreatedAt      string
	IsMarketClosed string
}

type DBResponseContainer struct {
	ListStringFromDB []StringResponse
}
type StringResponse struct {
	ListString []string
}
type AccountBalance struct {
	Balance string
}

type ShortDayAnalysis struct {
	IsShortDay string
}
type ShortCalendarDay struct {
	DateOfOccurance string
}

type EndOfDayDow struct {
	EndOfDayDowValue string
}

type DayReviewStatement struct {
	DowValues           string
	IsMarketClosed      string
	IsDownDay           string
	IsCashDay           string
	PrevEndDow          string
	WisemenSymbolValues string
}

type TSPCollectionStatement struct {
	DataCache string
}
