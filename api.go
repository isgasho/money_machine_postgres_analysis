package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type JSONObject struct {
	ID        int
	Age       int
	FirstName string
	// LastName  string
	// Email     string
}

type BrokerageQuery struct {
	RequestType string
	Name        string
}

func coolPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the coolPage!")
	fmt.Println("Endpoint Hit: coolPage")
}

func stockQuery(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var brokerageQuery BrokerageQuery
	err := decoder.Decode(&brokerageQuery)
	if err != nil {
		panic(err)
		fmt.Println("Stock error1")
	}
	fmt.Println(brokerageQuery.RequestType)

}

// func create(w http.ResponseWriter, r *http.Request) {
// 	create(11)
// }

// func read(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: delete")
// 	deleteEntry(11)
// 	fmt.Println("Endpoint Hit end: delete")
// }

// func update(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: delete")
// 	deleteEntry(11)
// 	fmt.Println("Endpoint Hit end: delete")
// }

// func delete(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: delete")
// 	deleteEntry(11)
// 	fmt.Println("Endpoint Hit end: delete")
// }

func handleRequests() {
	// http.HandleFunc("/create", create)
	// http.HandleFunc("/read", read)
	// http.HandleFunc("/update", update)
	// http.HandleFunc("/delete", delete)

	http.HandleFunc("/stockQuery", stockQuery)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	// post()
	// get()
	// handleRequests()

	//wrap in timer
	queryResponseStringParsed := `<?xml version=\"1.0\" encoding=\"UTF-8\"?><response id=\"126565f9-ee57-4117-aff4-dcbf19f4d673\"><elapsedtime>0</elapsedtime>
		<quotes>
		<quotetype>Real Time -  market data real time, National Best Bid and Offer</quotetype>
		<quote>
			<adp_100>196.7560</adp_100>
			<adp_200>185.8540</adp_200>
			<adp_50>199.2840</adp_50>
			<adv_21>18696278</adv_21>
			<adv_30>19827715</adv_30>
			<adv_90>26735159</adv_90>
			<ask>207.08</ask>
			<ask_time>17:16</ask_time>
			<asksz>5</asksz>
			<basis>na</basis>
			<beta>1.071</beta>
			<bid>206.95</bid>
			<bid_time>17:15</bid_time>
			<bidsz>1</bidsz>
			<bidtick>d</bidtick>
			<chg>1.6500</chg>
			<chg_sign>d</chg_sign>
			<chg_t>na</chg_t>
			<cl>208.67</cl>
			<contract_size>na</contract_size>
			<cusip>na</cusip>
			<date>2019-07-25</date>
			<datetime>2019-07-25T15:59:00-04:00</datetime>
			<days_to_expiration>na</days_to_expiration>
			<div>0.77</div>
			<divexdate>20190510</divexdate>
			<divfreq>Q</divfreq>
			<divpaydt>20190516</divpaydt>
			<dollar_value>2880148662.79</dollar_value>
			<eps>11.9</eps>
			<exch>NASD</exch>
			<exch_desc>NASDAQ</exch_desc>
			<hi>209.24</hi>
			<iad>3.08</iad>
			<idelta>na</idelta>
			<igamma>na</igamma>
			<imp_volatility>na</imp_volatility>
			<incr_vl>1261226</incr_vl>
			<irho>na</irho>
			<issue_desc>na</issue_desc>
			<itheta>na</itheta>
			<ivega>na</ivega>
			<last>207.02</last>
			<lo>206.73</lo>
			<name>APPLE INC</name>
			<op_delivery>na</op_delivery>
			<op_flag>1</op_flag>
			<op_style>na</op_style>
			<op_subclass>na</op_subclass>
			<openinterest>na</openinterest>
			<opn>208.89</opn>
			<opt_val>na</opt_val>
			<pchg>0.79</pchg>
			<pchg_sign>na</pchg_sign>
			<pcls>208.67</pcls>
			<pe>17.5353</pe>
			<phi>209.15</phi>
			<plo>207.17</plo>
			<popn>207.67</popn>
			<pr_adp_100>196.6192</pr_adp_100>
			<pr_adp_200>185.6875</pr_adp_200>
			<pr_adp_50>198.4545</pr_adp_50>
			<pr_date>2019-07-24</pr_date>
			<pr_openinterest>na</pr_openinterest>
			<prbook>8.89</prbook>
			<prchg>1.00</prchg>
			<prem_mult>na</prem_mult>
			<put_call>na</put_call>
			<pvol>0</pvol>
			<qcond>0</qcond>
			<rootsymbol>na</rootsymbol>
			<secclass>0</secclass>
			<sesn>na</sesn>
			<sho>4601075000</sho>
			<strikeprice>na</strikeprice>
			<symbol>AAPL</symbol>
			<tcond>29</tcond>
			<timestamp>1564089376</timestamp>
			<tr_num>123348</tr_num>
			<tradetick>e</tradetick>
			<trend>na</trend>
			<under_cusip>na</under_cusip>
			<undersymbol>na</undersymbol>
			<vl>13862806</vl>
			<volatility12>0.3137</volatility12>
			<vwap>207.76</vwap>
			<wk52hi>233.47</wk52hi>
			<wk52hidate>20181003</wk52hidate>
			<wk52lo>142.00</wk52lo>
			<wk52lodate>20190103</wk52lodate>
			<xdate>na</xdate>
			<xday>na</xday>
			<xmonth>na</xmonth>
			<xyear>na</xyear>
			<yield>1.47601</yield>
		</quote>
		</quotes><error>Success</error></response>`

	var stock Stock
	stock = parseQuery(queryResponseStringParsed)
	fmt.Println(stock.Symbol, stock.Bid, stock.Ask, stock.Last, stock.Pchg, stock.Pcls)
}
