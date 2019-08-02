package main

import "fmt"

func topStockPull() {
}

// type fn func(params ...interface{})

func queryMultiStockPull(params ...interface{}) {

	// listVal := reflect.ValueOf(params[0])
	// var listSymbolsInterface interface{} = listVal.Index(0).Interface()

	// listSymbols := listSymbolsInterface.([]string)

	// symbol1 := listSymbols[0]
	// symbol2 := listSymbols[1]
	// symbol3 := listSymbols[2]

	// // Query construction.
	// // query_multi_stock
	// json := `{
	// 	"request_type": "query_multi_stock",
	// 	"data": [
	// 	`
	// json += "\"" + symbol1 + "\","
	// json += "\"" + symbol2 + "\","
	// json += "\"" + symbol3 + "\""
	// json = json + `]}`

	// url := "http://localhost:3000/api/brokerage"
	// // response := ""
	// response := post(url, json)
	// fmt.Println(response)

	response := `<?xml version="1.0" encoding="UTF-8"?><response id="2d8f6fce-8ac2-459c-b3cd-6d32775d7792"><elapsedtime>0</elapsedtime>
	<quotes>
	  <quotetype>Real Time -  market data real time, National Best Bid and Offer</quotetype>
	  <quote>
		<adp_100>197.3094</adp_100>
		<adp_200>187.9423</adp_200>
		<adp_50>202.1685</adp_50>
		<adv_21>22456709</adv_21>
		<adv_30>21245400</adv_30>
		<adv_90>26286361</adv_90>
		<ask>208.4000</ask>
		<ask_time>16:35</ask_time>
		<asksz>33</asksz>
		<basis>na</basis>
		<beta>1.0714</beta>
		<bid>208.3000</bid>
		<bid_time>16:33</bid_time>
		<bidsz>1</bidsz>
		<bidtick>u</bidtick>
		<chg>4.6100</chg>
		<chg_sign>d</chg_sign>
		<chg_t>na</chg_t>
		<cl>213.0400</cl>
		<contract_size>na</contract_size>
		<cusip>na</cusip>
		<date>2019-08-01</date>
		<datetime>2019-08-01T16:00:00-04:00</datetime>
		<days_to_expiration>na</days_to_expiration>
		<div>0.77</div>
		<divexdate>20190809</divexdate>
		<divfreq>Q</divfreq>
		<divpaydt>20190815</divpaydt>
		<dollar_value>11407167379.2735</dollar_value>
		<eps>11.74</eps>
		<exch>NASD</exch>
		<exch_desc>NASDAQ</exch_desc>
		<hi>218.0300</hi>
		<iad>3.08</iad>
		<idelta>na</idelta>
		<igamma>na</igamma>
		<imp_volatility>na</imp_volatility>
		<incr_vl>1482512</incr_vl>
		<irho>na</irho>
		<issue_desc>na</issue_desc>
		<itheta>na</itheta>
		<ivega>na</ivega>
		<last>208.4300</last>
		<lo>206.7435</lo>
		<name>APPLE INC</name>
		<op_delivery>na</op_delivery>
		<op_flag>1</op_flag>
		<op_style>na</op_style>
		<op_subclass>na</op_subclass>
		<openinterest>na</openinterest>
		<opn>213.9000</opn>
		<opt_val>na</opt_val>
		<pchg>2.16</pchg>
		<pchg_sign>na</pchg_sign>
		<pcls>213.0400</pcls>
		<pe>18.1465</pe>
		<phi>221.3700</phi>
		<plo>211.3000</plo>
		<popn>216.4200</popn>
		<pr_adp_100>197.1878</pr_adp_100>
		<pr_adp_200>187.7591</pr_adp_200>
		<pr_adp_50>201.6477</pr_adp_50>
		<pr_date>2019-07-31</pr_date>
		<pr_openinterest>na</pr_openinterest>
		<prbook>9.9441</prbook>
		<prchg>-3.3800</prchg>
		<prem_mult>na</prem_mult>
		<put_call>na</put_call>
		<pvol>0</pvol>
		<qcond>0</qcond>
		<rootsymbol>na</rootsymbol>
		<secclass>0</secclass>
		<sesn>na</sesn>
		<sho>4601075000</sho>
		<strikeprice>na</strikeprice>
		<symbol>Test1</symbol>
		<tcond>29</tcond>
		<timestamp>1564691755</timestamp>
		<tr_num>450273</tr_num>
		<tradetick>e</tradetick>
		<trend>na</trend>
		<under_cusip>na</under_cusip>
		<undersymbol>na</undersymbol>
		<vl>53755343</vl>
		<volatility12>0.314</volatility12>
		<vwap>212.2052</vwap>
		<wk52hi>233.47</wk52hi>
		<wk52hidate>20181003</wk52hidate>
		<wk52lo>142.00</wk52lo>
		<wk52lodate>20190103</wk52lodate>
		<xdate>na</xdate>
		<xday>na</xday>
		<xmonth>na</xmonth>
		<xyear>na</xyear>
		<yield>1.44574</yield>
	  </quote>
	  <quote>
		<adp_100>30.0418</adp_100>
		<adp_200>27.1697</adp_200>
		<adp_50>31.8860</adp_50>
		<adv_21>55136491</adv_21>
		<adv_30>52062046</adv_30>
		<adv_90>67201879</adv_90>
		<ask>29.90</ask>
		<ask_time>16:35</ask_time>
		<asksz>11</asksz>
		<basis>na</basis>
		<beta>3.2202</beta>
		<bid>29.87</bid>
		<bid_time>16:35</bid_time>
		<bidsz>5</bidsz>
		<bidtick>u</bidtick>
		<chg>0.5900</chg>
		<chg_sign>d</chg_sign>
		<chg_t>na</chg_t>
		<cl>30.45</cl>
		<contract_size>na</contract_size>
		<cusip>na</cusip>
		<date>2019-08-01</date>
		<datetime>2019-08-01T15:59:00-04:00</datetime>
		<days_to_expiration>na</days_to_expiration>
		<div>0</div>
		<divexdate>na</divexdate>
		<divfreq>N</divfreq>
		<divpaydt>na</divpaydt>
		<dollar_value>2433485305.67</dollar_value>
		<eps>0.17</eps>
		<exch>NASD</exch>
		<exch_desc>NASDAQ</exch_desc>
		<hi>31.48</hi>
		<iad>na</iad>
		<idelta>na</idelta>
		<igamma>na</igamma>
		<imp_volatility>na</imp_volatility>
		<incr_vl>1162664</incr_vl>
		<irho>na</irho>
		<issue_desc>na</issue_desc>
		<itheta>na</itheta>
		<ivega>na</ivega>
		<last>29.86</last>
		<lo>29.10</lo>
		<name>ADVANCED MICRO DEVICES INC</name>
		<op_delivery>na</op_delivery>
		<op_flag>1</op_flag>
		<op_style>na</op_style>
		<op_subclass>na</op_subclass>
		<openinterest>na</openinterest>
		<opn>30.50</opn>
		<opt_val>na</opt_val>
		<pchg>1.94</pchg>
		<pchg_sign>na</pchg_sign>
		<pcls>30.45</pcls>
		<pe>179.1177</pe>
		<phi>32.30</phi>
		<plo>30.30</plo>
		<popn>32.08</popn>
		<pr_adp_100>30.0094</pr_adp_100>
		<pr_adp_200>27.1457</pr_adp_200>
		<pr_adp_50>31.9420</pr_adp_50>
		<pr_date>2019-07-31</pr_date>
		<pr_openinterest>na</pr_openinterest>
		<prbook>25.52</prbook>
		<prchg>-1.63</prchg>
		<prem_mult>na</prem_mult>
		<put_call>na</put_call>
		<pvol>0</pvol>
		<qcond>0</qcond>
		<rootsymbol>na</rootsymbol>
		<secclass>0</secclass>
		<sesn>na</sesn>
		<sho>1081601000</sho>
		<strikeprice>na</strikeprice>
		<symbol>Test2</symbol>
		<tcond>29</tcond>
		<timestamp>1564691755</timestamp>
		<tr_num>237906</tr_num>
		<tradetick>e</tradetick>
		<trend>na</trend>
		<under_cusip>na</under_cusip>
		<undersymbol>na</undersymbol>
		<vl>79915913</vl>
		<volatility12>0.676</volatility12>
		<vwap>30.45</vwap>
		<wk52hi>34.86</wk52hi>
		<wk52hidate>20190716</wk52hidate>
		<wk52lo>16.03</wk52lo>
		<wk52lodate>20181226</wk52lodate>
		<xdate>na</xdate>
		<xday>na</xday>
		<xmonth>na</xmonth>
		<xyear>na</xyear>
		<yield>na</yield>
	  </quote>
	  <quote>
		<adp_100>10.1125</adp_100>
		<adp_200>9.8569</adp_200>
		<adp_50>10.3560</adp_50>
		<adv_21>51899793</adv_21>
		<adv_30>46842706</adv_30>
		<adv_90>49743560</adv_90>
		<ask>10.090</ask>
		<ask_time>16:32</ask_time>
		<asksz>203</asksz>
		<basis>na</basis>
		<beta>1.0311</beta>
		<bid>10.080</bid>
		<bid_time>16:32</bid_time>
		<bidsz>13</bidsz>
		<bidtick>u</bidtick>
		<chg>0.3700</chg>
		<chg_sign>d</chg_sign>
		<chg_t>na</chg_t>
		<cl>10.450</cl>
		<contract_size>na</contract_size>
		<cusip>na</cusip>
		<date>2019-08-01</date>
		<datetime>2019-08-01T00:00:00-04:00</datetime>
		<days_to_expiration>na</days_to_expiration>
		<div>0.01</div>
		<divexdate>20190628</divexdate>
		<divfreq>Q</divfreq>
		<divpaydt>20190725</divpaydt>
		<dollar_value>875480114.014</dollar_value>
		<eps>0</eps>
		<exch>NYSE</exch>
		<exch_desc>New York Stock Exchange</exch_desc>
		<hi>10.485</hi>
		<iad>0.04</iad>
		<idelta>na</idelta>
		<igamma>na</igamma>
		<imp_volatility>na</imp_volatility>
		<incr_vl>0</incr_vl>
		<irho>na</irho>
		<issue_desc>na</issue_desc>
		<itheta>na</itheta>
		<ivega>na</ivega>
		<last>10.080</last>
		<lo>9.980</lo>
		<name>GENERAL ELECTRIC CO</name>
		<op_delivery>na</op_delivery>
		<op_flag>1</op_flag>
		<op_style>na</op_style>
		<op_subclass>na</op_subclass>
		<openinterest>na</openinterest>
		<opn>10.370</opn>
		<opt_val>na</opt_val>
		<pchg>3.54</pchg>
		<pchg_sign>na</pchg_sign>
		<pcls>10.450</pcls>
		<pe>0</pe>
		<phi>10.790</phi>
		<plo>10.050</plo>
		<popn>10.760</popn>
		<pr_adp_100>10.0964</pr_adp_100>
		<pr_adp_200>9.8526</pr_adp_200>
		<pr_adp_50>10.3465</pr_adp_50>
		<pr_date>2019-07-31</pr_date>
		<pr_openinterest>na</pr_openinterest>
		<prbook>2.847</prbook>
		<prchg>-0.310</prchg>
		<prem_mult>na</prem_mult>
		<put_call>na</put_call>
		<pvol>0</pvol>
		<qcond>0</qcond>
		<rootsymbol>na</rootsymbol>
		<secclass>0</secclass>
		<sesn>na</sesn>
		<sho>8727072000</sho>
		<strikeprice>na</strikeprice>
		<symbol>Test3</symbol>
		<tcond>29</tcond>
		<timestamp>1564691755</timestamp>
		<tr_num>146647</tr_num>
		<tradetick>e</tradetick>
		<trend>na</trend>
		<under_cusip>na</under_cusip>
		<undersymbol>na</undersymbol>
		<vl>86147063</vl>
		<volatility12>0.4318</volatility12>
		<vwap>10.162</vwap>
		<wk52hi>13.2460</wk52hi>
		<wk52hidate>20181009</wk52hidate>
		<wk52lo>6.4019</wk52lo>
		<wk52lodate>20181211</wk52lodate>
		<xdate>na</xdate>
		<xday>na</xday>
		<xmonth>na</xmonth>
		<xyear>na</xyear>
		<yield>0.38278</yield>
	  </quote>
	</quotes><error>Success</error></response>`

	stockList := parseStockSetQuery(response)

	for i, v := range stockList {
		// fmt.Println(v.Prchg)
		insertStock(v)
		i++
	}

}

func queryStoreQueryMonitored(params ...interface{}) {

	response := `<?xml version="1.0" encoding="UTF-8"?><response id="2d8f6fce-8ac2-459c-b3cd-6d32775d7792"><elapsedtime>0</elapsedtime>
	<quotes>
	  <quotetype>Real Time -  market data real time, National Best Bid and Offer</quotetype>
	  <quote>
		<adp_100>197.3094</adp_100>
		<adp_200>187.9423</adp_200>
		<adp_50>202.1685</adp_50>
		<adv_21>22456709</adv_21>
		<adv_30>21245400</adv_30>
		<adv_90>26286361</adv_90>
		<ask>208.4000</ask>
		<ask_time>16:35</ask_time>
		<asksz>33</asksz>
		<basis>na</basis>
		<beta>1.0714</beta>
		<bid>208.3000</bid>
		<bid_time>16:33</bid_time>
		<bidsz>1</bidsz>
		<bidtick>u</bidtick>
		<chg>4.6100</chg>
		<chg_sign>d</chg_sign>
		<chg_t>na</chg_t>
		<cl>213.0400</cl>
		<contract_size>na</contract_size>
		<cusip>na</cusip>
		<date>2019-08-01</date>
		<datetime>2019-08-01T16:00:00-04:00</datetime>
		<days_to_expiration>na</days_to_expiration>
		<div>0.77</div>
		<divexdate>20190809</divexdate>
		<divfreq>Q</divfreq>
		<divpaydt>20190815</divpaydt>
		<dollar_value>11407167379.2735</dollar_value>
		<eps>11.74</eps>
		<exch>NASD</exch>
		<exch_desc>NASDAQ</exch_desc>
		<hi>218.0300</hi>
		<iad>3.08</iad>
		<idelta>na</idelta>
		<igamma>na</igamma>
		<imp_volatility>na</imp_volatility>
		<incr_vl>1482512</incr_vl>
		<irho>na</irho>
		<issue_desc>na</issue_desc>
		<itheta>na</itheta>
		<ivega>na</ivega>
		<last>208.4300</last>
		<lo>206.7435</lo>
		<name>APPLE INC</name>
		<op_delivery>na</op_delivery>
		<op_flag>1</op_flag>
		<op_style>na</op_style>
		<op_subclass>na</op_subclass>
		<openinterest>na</openinterest>
		<opn>213.9000</opn>
		<opt_val>na</opt_val>
		<pchg>2.16</pchg>
		<pchg_sign>na</pchg_sign>
		<pcls>213.0400</pcls>
		<pe>18.1465</pe>
		<phi>221.3700</phi>
		<plo>211.3000</plo>
		<popn>216.4200</popn>
		<pr_adp_100>197.1878</pr_adp_100>
		<pr_adp_200>187.7591</pr_adp_200>
		<pr_adp_50>201.6477</pr_adp_50>
		<pr_date>2019-07-31</pr_date>
		<pr_openinterest>na</pr_openinterest>
		<prbook>9.9441</prbook>
		<prchg>-3.3800</prchg>
		<prem_mult>na</prem_mult>
		<put_call>na</put_call>
		<pvol>0</pvol>
		<qcond>0</qcond>
		<rootsymbol>na</rootsymbol>
		<secclass>0</secclass>
		<sesn>na</sesn>
		<sho>4601075000</sho>
		<strikeprice>na</strikeprice>
		<symbol>Test1</symbol>
		<tcond>29</tcond>
		<timestamp>1564691755</timestamp>
		<tr_num>450273</tr_num>
		<tradetick>e</tradetick>
		<trend>na</trend>
		<under_cusip>na</under_cusip>
		<undersymbol>na</undersymbol>
		<vl>53755343</vl>
		<volatility12>0.314</volatility12>
		<vwap>212.2052</vwap>
		<wk52hi>233.47</wk52hi>
		<wk52hidate>20181003</wk52hidate>
		<wk52lo>142.00</wk52lo>
		<wk52lodate>20190103</wk52lodate>
		<xdate>na</xdate>
		<xday>na</xday>
		<xmonth>na</xmonth>
		<xyear>na</xyear>
		<yield>1.44574</yield>
	  </quote>
	  <quote>
		<adp_100>30.0418</adp_100>
		<adp_200>27.1697</adp_200>
		<adp_50>31.8860</adp_50>
		<adv_21>55136491</adv_21>
		<adv_30>52062046</adv_30>
		<adv_90>67201879</adv_90>
		<ask>29.90</ask>
		<ask_time>16:35</ask_time>
		<asksz>11</asksz>
		<basis>na</basis>
		<beta>3.2202</beta>
		<bid>29.87</bid>
		<bid_time>16:35</bid_time>
		<bidsz>5</bidsz>
		<bidtick>u</bidtick>
		<chg>0.5900</chg>
		<chg_sign>d</chg_sign>
		<chg_t>na</chg_t>
		<cl>30.45</cl>
		<contract_size>na</contract_size>
		<cusip>na</cusip>
		<date>2019-08-01</date>
		<datetime>2019-08-01T15:59:00-04:00</datetime>
		<days_to_expiration>na</days_to_expiration>
		<div>0</div>
		<divexdate>na</divexdate>
		<divfreq>N</divfreq>
		<divpaydt>na</divpaydt>
		<dollar_value>2433485305.67</dollar_value>
		<eps>0.17</eps>
		<exch>NASD</exch>
		<exch_desc>NASDAQ</exch_desc>
		<hi>31.48</hi>
		<iad>na</iad>
		<idelta>na</idelta>
		<igamma>na</igamma>
		<imp_volatility>na</imp_volatility>
		<incr_vl>1162664</incr_vl>
		<irho>na</irho>
		<issue_desc>na</issue_desc>
		<itheta>na</itheta>
		<ivega>na</ivega>
		<last>29.86</last>
		<lo>29.10</lo>
		<name>ADVANCED MICRO DEVICES INC</name>
		<op_delivery>na</op_delivery>
		<op_flag>1</op_flag>
		<op_style>na</op_style>
		<op_subclass>na</op_subclass>
		<openinterest>na</openinterest>
		<opn>30.50</opn>
		<opt_val>na</opt_val>
		<pchg>1.94</pchg>
		<pchg_sign>na</pchg_sign>
		<pcls>30.45</pcls>
		<pe>179.1177</pe>
		<phi>32.30</phi>
		<plo>30.30</plo>
		<popn>32.08</popn>
		<pr_adp_100>30.0094</pr_adp_100>
		<pr_adp_200>27.1457</pr_adp_200>
		<pr_adp_50>31.9420</pr_adp_50>
		<pr_date>2019-07-31</pr_date>
		<pr_openinterest>na</pr_openinterest>
		<prbook>25.52</prbook>
		<prchg>-1.63</prchg>
		<prem_mult>na</prem_mult>
		<put_call>na</put_call>
		<pvol>0</pvol>
		<qcond>0</qcond>
		<rootsymbol>na</rootsymbol>
		<secclass>0</secclass>
		<sesn>na</sesn>
		<sho>1081601000</sho>
		<strikeprice>na</strikeprice>
		<symbol>Test2</symbol>
		<tcond>29</tcond>
		<timestamp>1564691755</timestamp>
		<tr_num>237906</tr_num>
		<tradetick>e</tradetick>
		<trend>na</trend>
		<under_cusip>na</under_cusip>
		<undersymbol>na</undersymbol>
		<vl>79915913</vl>
		<volatility12>0.676</volatility12>
		<vwap>30.45</vwap>
		<wk52hi>34.86</wk52hi>
		<wk52hidate>20190716</wk52hidate>
		<wk52lo>16.03</wk52lo>
		<wk52lodate>20181226</wk52lodate>
		<xdate>na</xdate>
		<xday>na</xday>
		<xmonth>na</xmonth>
		<xyear>na</xyear>
		<yield>na</yield>
	  </quote>
	  <quote>
		<adp_100>10.1125</adp_100>
		<adp_200>9.8569</adp_200>
		<adp_50>10.3560</adp_50>
		<adv_21>51899793</adv_21>
		<adv_30>46842706</adv_30>
		<adv_90>49743560</adv_90>
		<ask>10.090</ask>
		<ask_time>16:32</ask_time>
		<asksz>203</asksz>
		<basis>na</basis>
		<beta>1.0311</beta>
		<bid>10.080</bid>
		<bid_time>16:32</bid_time>
		<bidsz>13</bidsz>
		<bidtick>u</bidtick>
		<chg>0.3700</chg>
		<chg_sign>d</chg_sign>
		<chg_t>na</chg_t>
		<cl>10.450</cl>
		<contract_size>na</contract_size>
		<cusip>na</cusip>
		<date>2019-08-01</date>
		<datetime>2019-08-01T00:00:00-04:00</datetime>
		<days_to_expiration>na</days_to_expiration>
		<div>0.01</div>
		<divexdate>20190628</divexdate>
		<divfreq>Q</divfreq>
		<divpaydt>20190725</divpaydt>
		<dollar_value>875480114.014</dollar_value>
		<eps>0</eps>
		<exch>NYSE</exch>
		<exch_desc>New York Stock Exchange</exch_desc>
		<hi>10.485</hi>
		<iad>0.04</iad>
		<idelta>na</idelta>
		<igamma>na</igamma>
		<imp_volatility>na</imp_volatility>
		<incr_vl>0</incr_vl>
		<irho>na</irho>
		<issue_desc>na</issue_desc>
		<itheta>na</itheta>
		<ivega>na</ivega>
		<last>10.080</last>
		<lo>9.980</lo>
		<name>GENERAL ELECTRIC CO</name>
		<op_delivery>na</op_delivery>
		<op_flag>1</op_flag>
		<op_style>na</op_style>
		<op_subclass>na</op_subclass>
		<openinterest>na</openinterest>
		<opn>10.370</opn>
		<opt_val>na</opt_val>
		<pchg>3.54</pchg>
		<pchg_sign>na</pchg_sign>
		<pcls>10.450</pcls>
		<pe>0</pe>
		<phi>10.790</phi>
		<plo>10.050</plo>
		<popn>10.760</popn>
		<pr_adp_100>10.0964</pr_adp_100>
		<pr_adp_200>9.8526</pr_adp_200>
		<pr_adp_50>10.3465</pr_adp_50>
		<pr_date>2019-07-31</pr_date>
		<pr_openinterest>na</pr_openinterest>
		<prbook>2.847</prbook>
		<prchg>-0.310</prchg>
		<prem_mult>na</prem_mult>
		<put_call>na</put_call>
		<pvol>0</pvol>
		<qcond>0</qcond>
		<rootsymbol>na</rootsymbol>
		<secclass>0</secclass>
		<sesn>na</sesn>
		<sho>8727072000</sho>
		<strikeprice>na</strikeprice>
		<symbol>Test3</symbol>
		<tcond>29</tcond>
		<timestamp>1564691755</timestamp>
		<tr_num>146647</tr_num>
		<tradetick>e</tradetick>
		<trend>na</trend>
		<under_cusip>na</under_cusip>
		<undersymbol>na</undersymbol>
		<vl>86147063</vl>
		<volatility12>0.4318</volatility12>
		<vwap>10.162</vwap>
		<wk52hi>13.2460</wk52hi>
		<wk52hidate>20181009</wk52hidate>
		<wk52lo>6.4019</wk52lo>
		<wk52lodate>20181211</wk52lodate>
		<xdate>na</xdate>
		<xday>na</xday>
		<xmonth>na</xmonth>
		<xyear>na</xyear>
		<yield>0.38278</yield>
	  </quote>
	</quotes><error>Success</error></response>`

	stockList := parseMonitoredStockQuery(response)

	for i, v := range stockList {
		// fmt.Println(v.Prchg)
		insertStock(v)
		i++
	}
}

func queryMonitoredStocks(params ...interface{}) {
	monitoredSymbolList := selectAllMonitoringStock()
	fmt.Println(monitoredSymbolList)

	// response := `<?xml version="1.0" encoding="UTF-8"?><response id="2d8f6fce-8ac2-459c-b3cd-6d32775d7792"><elapsedtime>0</elapsedtime>
	// <quotes>
	//   <quotetype>Real Time -  market data real time, National Best Bid and Offer</quotetype>
	//   <quote>
	// 	<adp_100>197.3094</adp_100>
	// 	<adp_200>187.9423</adp_200>
	// 	<adp_50>202.1685</adp_50>
	// 	<adv_21>22456709</adv_21>
	// 	<adv_30>21245400</adv_30>
	// 	<adv_90>26286361</adv_90>
	// 	<ask>208.4000</ask>
	// 	<ask_time>16:35</ask_time>
	// 	<asksz>33</asksz>
	// 	<basis>na</basis>
	// 	<beta>1.0714</beta>
	// 	<bid>208.3000</bid>
	// 	<bid_time>16:33</bid_time>
	// 	<bidsz>1</bidsz>
	// 	<bidtick>u</bidtick>
	// 	<chg>4.6100</chg>
	// 	<chg_sign>d</chg_sign>
	// 	<chg_t>na</chg_t>
	// 	<cl>213.0400</cl>
	// 	<contract_size>na</contract_size>
	// 	<cusip>na</cusip>
	// 	<date>2019-08-01</date>
	// 	<datetime>2019-08-01T16:00:00-04:00</datetime>
	// 	<days_to_expiration>na</days_to_expiration>
	// 	<div>0.77</div>
	// 	<divexdate>20190809</divexdate>
	// 	<divfreq>Q</divfreq>
	// 	<divpaydt>20190815</divpaydt>
	// 	<dollar_value>11407167379.2735</dollar_value>
	// 	<eps>11.74</eps>
	// 	<exch>NASD</exch>
	// 	<exch_desc>NASDAQ</exch_desc>
	// 	<hi>218.0300</hi>
	// 	<iad>3.08</iad>
	// 	<idelta>na</idelta>
	// 	<igamma>na</igamma>
	// 	<imp_volatility>na</imp_volatility>
	// 	<incr_vl>1482512</incr_vl>
	// 	<irho>na</irho>
	// 	<issue_desc>na</issue_desc>
	// 	<itheta>na</itheta>
	// 	<ivega>na</ivega>
	// 	<last>208.4300</last>
	// 	<lo>206.7435</lo>
	// 	<name>APPLE INC</name>
	// 	<op_delivery>na</op_delivery>
	// 	<op_flag>1</op_flag>
	// 	<op_style>na</op_style>
	// 	<op_subclass>na</op_subclass>
	// 	<openinterest>na</openinterest>
	// 	<opn>213.9000</opn>
	// 	<opt_val>na</opt_val>
	// 	<pchg>2.16</pchg>
	// 	<pchg_sign>na</pchg_sign>
	// 	<pcls>213.0400</pcls>
	// 	<pe>18.1465</pe>
	// 	<phi>221.3700</phi>
	// 	<plo>211.3000</plo>
	// 	<popn>216.4200</popn>
	// 	<pr_adp_100>197.1878</pr_adp_100>
	// 	<pr_adp_200>187.7591</pr_adp_200>
	// 	<pr_adp_50>201.6477</pr_adp_50>
	// 	<pr_date>2019-07-31</pr_date>
	// 	<pr_openinterest>na</pr_openinterest>
	// 	<prbook>9.9441</prbook>
	// 	<prchg>-3.3800</prchg>
	// 	<prem_mult>na</prem_mult>
	// 	<put_call>na</put_call>
	// 	<pvol>0</pvol>
	// 	<qcond>0</qcond>
	// 	<rootsymbol>na</rootsymbol>
	// 	<secclass>0</secclass>
	// 	<sesn>na</sesn>
	// 	<sho>4601075000</sho>
	// 	<strikeprice>na</strikeprice>
	// 	<symbol>Test1</symbol>
	// 	<tcond>29</tcond>
	// 	<timestamp>1564691755</timestamp>
	// 	<tr_num>450273</tr_num>
	// 	<tradetick>e</tradetick>
	// 	<trend>na</trend>
	// 	<under_cusip>na</under_cusip>
	// 	<undersymbol>na</undersymbol>
	// 	<vl>53755343</vl>
	// 	<volatility12>0.314</volatility12>
	// 	<vwap>212.2052</vwap>
	// 	<wk52hi>233.47</wk52hi>
	// 	<wk52hidate>20181003</wk52hidate>
	// 	<wk52lo>142.00</wk52lo>
	// 	<wk52lodate>20190103</wk52lodate>
	// 	<xdate>na</xdate>
	// 	<xday>na</xday>
	// 	<xmonth>na</xmonth>
	// 	<xyear>na</xyear>
	// 	<yield>1.44574</yield>
	//   </quote>
	//   <quote>
	// 	<adp_100>30.0418</adp_100>
	// 	<adp_200>27.1697</adp_200>
	// 	<adp_50>31.8860</adp_50>
	// 	<adv_21>55136491</adv_21>
	// 	<adv_30>52062046</adv_30>
	// 	<adv_90>67201879</adv_90>
	// 	<ask>29.90</ask>
	// 	<ask_time>16:35</ask_time>
	// 	<asksz>11</asksz>
	// 	<basis>na</basis>
	// 	<beta>3.2202</beta>
	// 	<bid>29.87</bid>
	// 	<bid_time>16:35</bid_time>
	// 	<bidsz>5</bidsz>
	// 	<bidtick>u</bidtick>
	// 	<chg>0.5900</chg>
	// 	<chg_sign>d</chg_sign>
	// 	<chg_t>na</chg_t>
	// 	<cl>30.45</cl>
	// 	<contract_size>na</contract_size>
	// 	<cusip>na</cusip>
	// 	<date>2019-08-01</date>
	// 	<datetime>2019-08-01T15:59:00-04:00</datetime>
	// 	<days_to_expiration>na</days_to_expiration>
	// 	<div>0</div>
	// 	<divexdate>na</divexdate>
	// 	<divfreq>N</divfreq>
	// 	<divpaydt>na</divpaydt>
	// 	<dollar_value>2433485305.67</dollar_value>
	// 	<eps>0.17</eps>
	// 	<exch>NASD</exch>
	// 	<exch_desc>NASDAQ</exch_desc>
	// 	<hi>31.48</hi>
	// 	<iad>na</iad>
	// 	<idelta>na</idelta>
	// 	<igamma>na</igamma>
	// 	<imp_volatility>na</imp_volatility>
	// 	<incr_vl>1162664</incr_vl>
	// 	<irho>na</irho>
	// 	<issue_desc>na</issue_desc>
	// 	<itheta>na</itheta>
	// 	<ivega>na</ivega>
	// 	<last>29.86</last>
	// 	<lo>29.10</lo>
	// 	<name>ADVANCED MICRO DEVICES INC</name>
	// 	<op_delivery>na</op_delivery>
	// 	<op_flag>1</op_flag>
	// 	<op_style>na</op_style>
	// 	<op_subclass>na</op_subclass>
	// 	<openinterest>na</openinterest>
	// 	<opn>30.50</opn>
	// 	<opt_val>na</opt_val>
	// 	<pchg>1.94</pchg>
	// 	<pchg_sign>na</pchg_sign>
	// 	<pcls>30.45</pcls>
	// 	<pe>179.1177</pe>
	// 	<phi>32.30</phi>
	// 	<plo>30.30</plo>
	// 	<popn>32.08</popn>
	// 	<pr_adp_100>30.0094</pr_adp_100>
	// 	<pr_adp_200>27.1457</pr_adp_200>
	// 	<pr_adp_50>31.9420</pr_adp_50>
	// 	<pr_date>2019-07-31</pr_date>
	// 	<pr_openinterest>na</pr_openinterest>
	// 	<prbook>25.52</prbook>
	// 	<prchg>-1.63</prchg>
	// 	<prem_mult>na</prem_mult>
	// 	<put_call>na</put_call>
	// 	<pvol>0</pvol>
	// 	<qcond>0</qcond>
	// 	<rootsymbol>na</rootsymbol>
	// 	<secclass>0</secclass>
	// 	<sesn>na</sesn>
	// 	<sho>1081601000</sho>
	// 	<strikeprice>na</strikeprice>
	// 	<symbol>Test2</symbol>
	// 	<tcond>29</tcond>
	// 	<timestamp>1564691755</timestamp>
	// 	<tr_num>237906</tr_num>
	// 	<tradetick>e</tradetick>
	// 	<trend>na</trend>
	// 	<under_cusip>na</under_cusip>
	// 	<undersymbol>na</undersymbol>
	// 	<vl>79915913</vl>
	// 	<volatility12>0.676</volatility12>
	// 	<vwap>30.45</vwap>
	// 	<wk52hi>34.86</wk52hi>
	// 	<wk52hidate>20190716</wk52hidate>
	// 	<wk52lo>16.03</wk52lo>
	// 	<wk52lodate>20181226</wk52lodate>
	// 	<xdate>na</xdate>
	// 	<xday>na</xday>
	// 	<xmonth>na</xmonth>
	// 	<xyear>na</xyear>
	// 	<yield>na</yield>
	//   </quote>
	//   <quote>
	// 	<adp_100>10.1125</adp_100>
	// 	<adp_200>9.8569</adp_200>
	// 	<adp_50>10.3560</adp_50>
	// 	<adv_21>51899793</adv_21>
	// 	<adv_30>46842706</adv_30>
	// 	<adv_90>49743560</adv_90>
	// 	<ask>10.090</ask>
	// 	<ask_time>16:32</ask_time>
	// 	<asksz>203</asksz>
	// 	<basis>na</basis>
	// 	<beta>1.0311</beta>
	// 	<bid>10.080</bid>
	// 	<bid_time>16:32</bid_time>
	// 	<bidsz>13</bidsz>
	// 	<bidtick>u</bidtick>
	// 	<chg>0.3700</chg>
	// 	<chg_sign>d</chg_sign>
	// 	<chg_t>na</chg_t>
	// 	<cl>10.450</cl>
	// 	<contract_size>na</contract_size>
	// 	<cusip>na</cusip>
	// 	<date>2019-08-01</date>
	// 	<datetime>2019-08-01T00:00:00-04:00</datetime>
	// 	<days_to_expiration>na</days_to_expiration>
	// 	<div>0.01</div>
	// 	<divexdate>20190628</divexdate>
	// 	<divfreq>Q</divfreq>
	// 	<divpaydt>20190725</divpaydt>
	// 	<dollar_value>875480114.014</dollar_value>
	// 	<eps>0</eps>
	// 	<exch>NYSE</exch>
	// 	<exch_desc>New York Stock Exchange</exch_desc>
	// 	<hi>10.485</hi>
	// 	<iad>0.04</iad>
	// 	<idelta>na</idelta>
	// 	<igamma>na</igamma>
	// 	<imp_volatility>na</imp_volatility>
	// 	<incr_vl>0</incr_vl>
	// 	<irho>na</irho>
	// 	<issue_desc>na</issue_desc>
	// 	<itheta>na</itheta>
	// 	<ivega>na</ivega>
	// 	<last>10.080</last>
	// 	<lo>9.980</lo>
	// 	<name>GENERAL ELECTRIC CO</name>
	// 	<op_delivery>na</op_delivery>
	// 	<op_flag>1</op_flag>
	// 	<op_style>na</op_style>
	// 	<op_subclass>na</op_subclass>
	// 	<openinterest>na</openinterest>
	// 	<opn>10.370</opn>
	// 	<opt_val>na</opt_val>
	// 	<pchg>3.54</pchg>
	// 	<pchg_sign>na</pchg_sign>
	// 	<pcls>10.450</pcls>
	// 	<pe>0</pe>
	// 	<phi>10.790</phi>
	// 	<plo>10.050</plo>
	// 	<popn>10.760</popn>
	// 	<pr_adp_100>10.0964</pr_adp_100>
	// 	<pr_adp_200>9.8526</pr_adp_200>
	// 	<pr_adp_50>10.3465</pr_adp_50>
	// 	<pr_date>2019-07-31</pr_date>
	// 	<pr_openinterest>na</pr_openinterest>
	// 	<prbook>2.847</prbook>
	// 	<prchg>-0.310</prchg>
	// 	<prem_mult>na</prem_mult>
	// 	<put_call>na</put_call>
	// 	<pvol>0</pvol>
	// 	<qcond>0</qcond>
	// 	<rootsymbol>na</rootsymbol>
	// 	<secclass>0</secclass>
	// 	<sesn>na</sesn>
	// 	<sho>8727072000</sho>
	// 	<strikeprice>na</strikeprice>
	// 	<symbol>Test3</symbol>
	// 	<tcond>29</tcond>
	// 	<timestamp>1564691755</timestamp>
	// 	<tr_num>146647</tr_num>
	// 	<tradetick>e</tradetick>
	// 	<trend>na</trend>
	// 	<under_cusip>na</under_cusip>
	// 	<undersymbol>na</undersymbol>
	// 	<vl>86147063</vl>
	// 	<volatility12>0.4318</volatility12>
	// 	<vwap>10.162</vwap>
	// 	<wk52hi>13.2460</wk52hi>
	// 	<wk52hidate>20181009</wk52hidate>
	// 	<wk52lo>6.4019</wk52lo>
	// 	<wk52lodate>20181211</wk52lodate>
	// 	<xdate>na</xdate>
	// 	<xday>na</xday>
	// 	<xmonth>na</xmonth>
	// 	<xyear>na</xyear>
	// 	<yield>0.38278</yield>
	//   </quote>
	// </quotes><error>Success</error></response>`

	// stockList := parseMonitoredStockQuery(response)

	// for i, v := range stockList {
	// 	// fmt.Println(v.Prchg)
	// 	insertStock(v)
	// 	i++
	// }
}
