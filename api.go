package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var isTimeMonitoringLoop bool

func coolPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the coolPage!")
	fmt.Println("Endpoint Hit: coolPage")
}

func stockQuery(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	// var brokerageQuery BrokerageQuery
	var requestCase string
	err := decoder.Decode(&requestCase)
	if err != nil {
		panic(err)
		fmt.Println("Stock error1")
	}
	fmt.Println(requestCase)
}

func databaseQuery(rw http.ResponseWriter, req *http.Request) {
	// decoder := json.NewDecoder(req.Body)
	// // var brokerageQuery BrokerageQuery
	// err := decoder.Decode(&brokerageQuery)
	// if err != nil {
	// 	panic(err)
	// 	fmt.Println("Stock error1")
	// }
	// fmt.Println(brokerageQuery.RequestType)
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

	http.HandleFunc("/databaseQuery", databaseQuery)
	log.Fatal(http.ListenAndServe(":10000", nil))

	// timein := time.Now().Local().AddDate(Hours, Mins, Sec)
	// fmt.Println(timein)

}

func main() {
	// post()
	// get()
	// handleRequests()

	// ticker := time.NewTicker(500 * time.Millisecond)
	// go func() {
	// 	for t := range ticker.C {
	// 		fmt.Println("Tick at", t)
	// 	}
	// }()
	// time.Sleep(1600 * time.Millisecond)
	// h, _ := time.ParseDuration("4h30m")
	// fmt.Printf("I've got %.1f hours of work left.", h.Hours())

	// ticker.Stop()
	// fmt.Println("Ticker stopped")

	// numLoops := 144
	// timePerIteration := time.Duration(3) * time.Second / time.Duration(numLoops)
	// ticker := time.NewTicker(timePerIteration)
	// for i := 0; i < numLoops; i++ {
	// 	// your code
	// 	fmt.Println("cool")
	// 	<-ticker.C
	// }
	// ticker.Stop()

	// index := 0
	// for index < 3 {
	// 	time.Sleep(2 * time.Second)
	// 	index++
	// 	testing()
	// }

	// insertDay("7")
	// var newsEntry = News{
	// 	DayID:    3,
	// 	NewsInfo: "Awesome news about Dow",
	// }
	// // fmt.Println(newsEntry.DayID, newsEntry.NewsInfo)
	// insertNews(newsEntry)

	// var dowEntry = Dow{
	// 	DayID:   3,
	// 	DowInfo: "Awesome Dow about Dow",
	// }
	// insertDow(dowEntry)

	// selectAllStock("DEL")

	// query := `<?xml version="1.0" encoding="UTF-8"?>
	// <response id="58d6662f-2410-4e43-ba2b-5298e22c5aae">
	// 	<elapsedtime>0</elapsedtime>
	// 	<quotes>
	// 		<quote>
	// 			<chg>28.835</chg>
	// 			<chg_sign>u</chg_sign>
	// 			<last>808.695</last>
	// 			<name>CHIPOTLE MEXICAN GRILL INC</name>
	// 			<pchg>3.70</pchg>
	// 			<pcls>779.860</pcls>
	// 			<rank>1</rank>
	// 			<symbol>CMG</symbol>
	// 			<vl>922968</vl>
	// 		</quote>
	// 		<quote>
	// 			<chg>10.20</chg>
	// 			<chg_sign>u</chg_sign>
	// 			<last>396.21</last>
	// 			<name>CHEMED CORP</name>
	// 			<pchg>2.64</pchg>
	// 			<pcls>386.01</pcls>
	// 			<rank>2</rank>
	// 			<symbol>CHE</symbol>
	// 			<vl>140181</vl>
	// 		</quote>
	// 		<quote>
	// 			<chg>5.13</chg>
	// 			<chg_sign>u</chg_sign>
	// 			<last>79.77</last>
	// 			<name>GRUBHUB INC</name>
	// 			<pchg>6.87</pchg>
	// 			<pcls>74.64</pcls>
	// 			<rank>3</rank>
	// 			<symbol>GRUB</symbol>
	// 			<vl>4742848</vl>
	// 		</quote>
	// 		<quote>
	// 			<chg>5.10</chg>
	// 			<chg_sign>u</chg_sign>
	// 			<last>94.60</last>
	// 			<name>ASBURY AUTOMOTIVE GROUP INC</name>
	// 			<pchg>5.70</pchg>
	// 			<pcls>89.50</pcls>
	// 			<rank>4</rank>
	// 			<symbol>ABG</symbol>
	// 			<vl>217864</vl>
	// 		</quote>
	// 		<quote>
	// 			<chg>4.89</chg>
	// 			<chg_sign>u</chg_sign>
	// 			<last>172.19</last>
	// 			<name>CIGNA CORP (NEW)</name>
	// 			<pchg>2.92</pchg>
	// 			<pcls>167.30</pcls>
	// 			<rank>5</rank>
	// 			<symbol>CI</symbol>
	// 			<vl>1281650</vl>
	// 		</quote>
	// 		<quote>
	// 			<chg>4.0200</chg>
	// 			<chg_sign>u</chg_sign>
	// 			<last>280.8900</last>
	// 			<name>HUMANA INC.</name>
	// 			<pchg>1.45</pchg>
	// 			<pcls>276.8700</pcls>
	// 			<rank>6</rank>
	// 			<symbol>HUM</symbol>
	// 			<vl>635934</vl>
	// 		</quote>
	// 	</quotes>
	// 	<error>Success</error>
	// </response>`

	// 	query := `<?xml version="1.0" encoding="UTF-8"?>
	// <response id="58d6662f-2410-4e43-ba2b-5298e22c5aae">
	//     <elapsedtime>0</elapsedtime>
	//     <quotes>
	//         <quote>
	//             <chg>28.835</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>808.695</last>
	//             <name>CHIPOTLE MEXICAN GRILL INC</name>
	//             <pchg>3.70</pchg>
	//             <pcls>779.860</pcls>
	//             <rank>1</rank>
	//             <symbol>CMG</symbol>
	//             <vl>922968</vl>
	//         </quote>
	//         <quote>
	//             <chg>10.20</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>396.21</last>
	//             <name>CHEMED CORP</name>
	//             <pchg>2.64</pchg>
	//             <pcls>386.01</pcls>
	//             <rank>2</rank>
	//             <symbol>CHE</symbol>
	//             <vl>140181</vl>
	//         </quote>
	//         <quote>
	//             <chg>5.13</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>79.77</last>
	//             <name>GRUBHUB INC</name>
	//             <pchg>6.87</pchg>
	//             <pcls>74.64</pcls>
	//             <rank>3</rank>
	//             <symbol>GRUB</symbol>
	//             <vl>4742848</vl>
	//         </quote>
	//         <quote>
	//             <chg>5.10</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>94.60</last>
	//             <name>ASBURY AUTOMOTIVE GROUP INC</name>
	//             <pchg>5.70</pchg>
	//             <pcls>89.50</pcls>
	//             <rank>4</rank>
	//             <symbol>ABG</symbol>
	//             <vl>217864</vl>
	//         </quote>
	//         <quote>
	//             <chg>4.89</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>172.19</last>
	//             <name>CIGNA CORP (NEW)</name>
	//             <pchg>2.92</pchg>
	//             <pcls>167.30</pcls>
	//             <rank>5</rank>
	//             <symbol>CI</symbol>
	//             <vl>1281650</vl>
	//         </quote>
	//         <quote>
	//             <chg>4.0200</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>280.8900</last>
	//             <name>HUMANA INC.</name>
	//             <pchg>1.45</pchg>
	//             <pcls>276.8700</pcls>
	//             <rank>6</rank>
	//             <symbol>HUM</symbol>
	//             <vl>635934</vl>
	//         </quote>
	//         <quote>
	//             <chg>3.9900</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>18.3000</last>
	//             <name>NA</name>
	//             <pchg>27.88</pchg>
	//             <pcls>14.3100</pcls>
	//             <rank>7</rank>
	//             <symbol>MDLQ</symbol>
	//             <vl>78743</vl>
	//         </quote>
	//         <quote>
	//             <chg>3.975</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>157.935</last>
	//             <name>GRUPO AEROPORTUARIO DEL SURESTE SA DE CV</name>
	//             <pchg>2.58</pchg>
	//             <pcls>153.960</pcls>
	//             <rank>8</rank>
	//             <symbol>ASR</symbol>
	//             <vl>66202</vl>
	//         </quote>
	//         <quote>
	//             <chg>3.830</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>307.900</last>
	//             <name>ESSEX PROPERTY TRUST INC</name>
	//             <pchg>1.26</pchg>
	//             <pcls>304.070</pcls>
	//             <rank>9</rank>
	//             <symbol>ESS</symbol>
	//             <vl>270050</vl>
	//         </quote>
	//         <quote>
	//             <chg>3.8050</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>18.2500</last>
	//             <name>NA</name>
	//             <pchg>26.34</pchg>
	//             <pcls>14.4450</pcls>
	//             <rank>10</rank>
	//             <symbol>MDLX</symbol>
	//             <vl>23802</vl>
	//         </quote>
	//         <quote>
	//             <chg>3.65</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>99.97</last>
	//             <name>GRUPO AEROPORTUARIO DEL PACIFICO, S.A.B DE C.V.</name>
	//             <pchg>3.79</pchg>
	//             <pcls>96.32</pcls>
	//             <rank>11</rank>
	//             <symbol>PAC</symbol>
	//             <vl>42256</vl>
	//         </quote>
	//         <quote>
	//             <chg>3.54</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>341.52</last>
	//             <name>COOPER COMPANIES, INC. (THE)</name>
	//             <pchg>1.05</pchg>
	//             <pcls>337.98</pcls>
	//             <rank>12</rank>
	//             <symbol>COO</symbol>
	//             <vl>160936</vl>
	//         </quote>
	//         <quote>
	//             <chg>3.40</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>216.25</last>
	//             <name>EDWARDS LIFESCIENCES CORP</name>
	//             <pchg>1.60</pchg>
	//             <pcls>212.85</pcls>
	//             <rank>13</rank>
	//             <symbol>EW</symbol>
	//             <vl>969765</vl>
	//         </quote>
	//         <quote>
	//             <chg>3.380</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>350.540</last>
	//             <name>TELEFLEX INCORPORATED</name>
	//             <pchg>0.97</pchg>
	//             <pcls>347.160</pcls>
	//             <rank>14</rank>
	//             <symbol>TFX</symbol>
	//             <vl>288078</vl>
	//         </quote>
	//         <quote>
	//             <chg>3.25</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>177.48</last>
	//             <name>UNION PACIFIC CORP</name>
	//             <pchg>1.87</pchg>
	//             <pcls>174.23</pcls>
	//             <rank>15</rank>
	//             <symbol>UNP</symbol>
	//             <vl>2672074</vl>
	//         </quote>
	//         <quote>
	//             <chg>3.2000</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>355.7200</last>
	//             <name>NORTHROP GRUMMAN CORP</name>
	//             <pchg>0.91</pchg>
	//             <pcls>352.5200</pcls>
	//             <rank>16</rank>
	//             <symbol>NOC</symbol>
	//             <vl>580233</vl>
	//         </quote>
	//         <quote>
	//             <chg>3.13</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>242.56</last>
	//             <name>PUBLIC STORAGE</name>
	//             <pchg>1.31</pchg>
	//             <pcls>239.43</pcls>
	//             <rank>17</rank>
	//             <symbol>PSA</symbol>
	//             <vl>574935</vl>
	//         </quote>
	//         <quote>
	//             <chg>2.95</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>168.03</last>
	//             <name>DIAGEO PLC</name>
	//             <pchg>1.79</pchg>
	//             <pcls>165.08</pcls>
	//             <rank>18</rank>
	//             <symbol>DEO</symbol>
	//             <vl>544454</vl>
	//         </quote>
	//         <quote>
	//             <chg>2.7200</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>176.7000</last>
	//             <name>3M CO</name>
	//             <pchg>1.56</pchg>
	//             <pcls>173.9800</pcls>
	//             <rank>19</rank>
	//             <symbol>MMM</symbol>
	//             <vl>2313329</vl>
	//         </quote>
	//         <quote>
	//             <chg>2.470</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>129.340</last>
	//             <name>RESMED INC.</name>
	//             <pchg>1.95</pchg>
	//             <pcls>126.870</pcls>
	//             <rank>20</rank>
	//             <symbol>RMD</symbol>
	//             <vl>698084</vl>
	//         </quote>
	//         <quote>
	//             <chg>2.44</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>293.04</last>
	//             <name>WELLCARE HEALTH PLANS INC</name>
	//             <pchg>0.84</pchg>
	//             <pcls>290.60</pcls>
	//             <rank>21</rank>
	//             <symbol>WCG</symbol>
	//             <vl>430786</vl>
	//         </quote>
	//         <quote>
	//             <chg>2.4000</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>744.4500</last>
	//             <name>TEXAS PACIFIC LAND TRUST</name>
	//             <pchg>0.32</pchg>
	//             <pcls>742.0500</pcls>
	//             <rank>22</rank>
	//             <symbol>TPL</symbol>
	//             <vl>8513</vl>
	//         </quote>
	//         <quote>
	//             <chg>2.39</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>71.59</last>
	//             <name>BOOZ ALLEN HAMILTON HOLDING CORP.</name>
	//             <pchg>3.45</pchg>
	//             <pcls>69.20</pcls>
	//             <rank>23</rank>
	//             <symbol>BAH</symbol>
	//             <vl>2440438</vl>
	//         </quote>
	//         <quote>
	//             <chg>2.36</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>104.98</last>
	//             <name>FTI CONSULTING INC.</name>
	//             <pchg>2.30</pchg>
	//             <pcls>102.62</pcls>
	//             <rank>24</rank>
	//             <symbol>FCN</symbol>
	//             <vl>238463</vl>
	//         </quote>
	//         <quote>
	//             <chg>2.34</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>133.07</last>
	//             <name>JOHNSON &amp; JOHNSON</name>
	//             <pchg>1.79</pchg>
	//             <pcls>130.73</pcls>
	//             <rank>25</rank>
	//             <symbol>JNJ</symbol>
	//             <vl>6126923</vl>
	//         </quote>
	//         <quote>
	//             <chg>2.2900</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>141.2400</last>
	//             <name>MCKESSON CORP</name>
	//             <pchg>1.65</pchg>
	//             <pcls>138.9500</pcls>
	//             <rank>26</rank>
	//             <symbol>MCK</symbol>
	//             <vl>881164</vl>
	//         </quote>
	//         <quote>
	//             <chg>2.2291</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>22.9791</last>
	//             <name>MEDLEY CAPITAL CORP</name>
	//             <pchg>10.74</pchg>
	//             <pcls>20.7500</pcls>
	//             <rank>27</rank>
	//             <symbol>MCV</symbol>
	//             <vl>27980</vl>
	//         </quote>
	//         <quote>
	//             <chg>2.1066</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>43.7066</last>
	//             <name>CHESAPEAKE ENERGY CORP.</name>
	//             <pchg>5.06</pchg>
	//             <pcls>41.6000</pcls>
	//             <rank>28</rank>
	//             <symbol>CHK.D</symbol>
	//             <vl>9288</vl>
	//         </quote>
	//         <quote>
	//             <chg>2.09</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>211.26</last>
	//             <name>NEXTERA ENERGY INC</name>
	//             <pchg>1.00</pchg>
	//             <pcls>209.17</pcls>
	//             <rank>29</rank>
	//             <symbol>NEE</symbol>
	//             <vl>1007708</vl>
	//         </quote>
	//         <quote>
	//             <chg>2.070</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>210.450</last>
	//             <name>AVALONBAY COMMUNITIES, INC.</name>
	//             <pchg>0.99</pchg>
	//             <pcls>208.380</pcls>
	//             <rank>30</rank>
	//             <symbol>AVB</symbol>
	//             <vl>381813</vl>
	//         </quote>
	//         <quote>
	//             <chg>2.07</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>192.06</last>
	//             <name>NORFOLK SOUTHERN CORP</name>
	//             <pchg>1.09</pchg>
	//             <pcls>189.99</pcls>
	//             <rank>31</rank>
	//             <symbol>NSC</symbol>
	//             <vl>1382116</vl>
	//         </quote>
	//         <quote>
	//             <chg>2.06</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>88.78</last>
	//             <name>GROUP 1 AUTOMOTIVE, INC.</name>
	//             <pchg>2.38</pchg>
	//             <pcls>86.72</pcls>
	//             <rank>32</rank>
	//             <symbol>GPI</symbol>
	//             <vl>282155</vl>
	//         </quote>
	//         <quote>
	//             <chg>2.015</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>78.285</last>
	//             <name>KIRBY CORP.</name>
	//             <pchg>2.64</pchg>
	//             <pcls>76.270</pcls>
	//             <rank>33</rank>
	//             <symbol>KEX</symbol>
	//             <vl>631439</vl>
	//         </quote>
	//         <quote>
	//             <chg>1.86</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>75.55</last>
	//             <name>COLGATE-PALMOLIVE CO.</name>
	//             <pchg>2.52</pchg>
	//             <pcls>73.69</pcls>
	//             <rank>34</rank>
	//             <symbol>CL</symbol>
	//             <vl>4139703</vl>
	//         </quote>
	//         <quote>
	//             <chg>1.85</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>137.95</last>
	//             <name>WORLDPAY INC</name>
	//             <pchg>1.36</pchg>
	//             <pcls>136.10</pcls>
	//             <rank>35</rank>
	//             <symbol>WP</symbol>
	//             <vl>3321890</vl>
	//         </quote>
	//         <quote>
	//             <chg>1.82</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>171.46</last>
	//             <name>GLOBAL PAYMENTS INC</name>
	//             <pchg>1.07</pchg>
	//             <pcls>169.64</pcls>
	//             <rank>36</rank>
	//             <symbol>GPN</symbol>
	//             <vl>2346662</vl>
	//         </quote>
	//         <quote>
	//             <chg>1.730</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>288.940</last>
	//             <name>THERMO FISHER SCIENTIFIC INC</name>
	//             <pchg>0.60</pchg>
	//             <pcls>287.210</pcls>
	//             <rank>37</rank>
	//             <symbol>TMO</symbol>
	//             <vl>783580</vl>
	//         </quote>
	//         <quote>
	//             <chg>1.68</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>218.96</last>
	//             <name>WEX INC</name>
	//             <pchg>0.77</pchg>
	//             <pcls>217.28</pcls>
	//             <rank>38</rank>
	//             <symbol>WEX</symbol>
	//             <vl>373960</vl>
	//         </quote>
	//         <quote>
	//             <chg>1.67</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>146.32</last>
	//             <name>DISNEY (WALT) CO. (THE)</name>
	//             <pchg>1.15</pchg>
	//             <pcls>144.65</pcls>
	//             <rank>39</rank>
	//             <symbol>DIS</symbol>
	//             <vl>9874091</vl>
	//         </quote>
	//         <quote>
	//             <chg>1.67</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>201.50</last>
	//             <name>L3HARRIS TECHNOLOGIES INC</name>
	//             <pchg>0.84</pchg>
	//             <pcls>199.83</pcls>
	//             <rank>40</rank>
	//             <symbol>LHX</symbol>
	//             <vl>1123499</vl>
	//         </quote>
	//         <quote>
	//             <chg>1.670</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>115.630</last>
	//             <name>ROYAL CARIBBEAN CRUISES LTD</name>
	//             <pchg>1.47</pchg>
	//             <pcls>113.960</pcls>
	//             <rank>41</rank>
	//             <symbol>RCL</symbol>
	//             <vl>1807224</vl>
	//         </quote>
	//         <quote>
	//             <chg>1.660</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>34.110</last>
	//             <name>CHEWY INC</name>
	//             <pchg>5.12</pchg>
	//             <pcls>32.450</pcls>
	//             <rank>42</rank>
	//             <symbol>CHWY</symbol>
	//             <vl>1819448</vl>
	//         </quote>
	//         <quote>
	//             <chg>1.63</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>93.23</last>
	//             <name>FOMENTO ECONOMICO MEXICANO, S.A.B. DE C.V.</name>
	//             <pchg>1.78</pchg>
	//             <pcls>91.60</pcls>
	//             <rank>43</rank>
	//             <symbol>FMX</symbol>
	//             <vl>362280</vl>
	//         </quote>
	//         <quote>
	//             <chg>1.6100</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>1147.7900</last>
	//             <name>AUTOZONE, INC.</name>
	//             <pchg>0.14</pchg>
	//             <pcls>1146.1800</pcls>
	//             <rank>44</rank>
	//             <symbol>AZO</symbol>
	//             <vl>170542</vl>
	//         </quote>
	//         <quote>
	//             <chg>1.590</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>23.600</last>
	//             <name>TENET HEALTHCARE CORP.</name>
	//             <pchg>7.22</pchg>
	//             <pcls>22.010</pcls>
	//             <rank>45</rank>
	//             <symbol>THC</symbol>
	//             <vl>3724356</vl>
	//         </quote>
	//         <quote>
	//             <chg>1.5800</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>134.5000</last>
	//             <name>CATERPILLAR INC.</name>
	//             <pchg>1.19</pchg>
	//             <pcls>132.9200</pcls>
	//             <rank>46</rank>
	//             <symbol>CAT</symbol>
	//             <vl>2412553</vl>
	//         </quote>
	//         <quote>
	//             <chg>1.5700</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>219.2000</last>
	//             <name>WATERS CORP.</name>
	//             <pchg>0.72</pchg>
	//             <pcls>217.6300</pcls>
	//             <rank>47</rank>
	//             <symbol>WAT</symbol>
	//             <vl>729133</vl>
	//         </quote>
	//         <quote>
	//             <chg>1.51</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>143.52</last>
	//             <name>CARLISLE COMPANIES INC.</name>
	//             <pchg>1.06</pchg>
	//             <pcls>142.01</pcls>
	//             <rank>48</rank>
	//             <symbol>CSL</symbol>
	//             <vl>384400</vl>
	//         </quote>
	//         <quote>
	//             <chg>1.490</chg>
	//             <chg_sign>u</chg_sign>
	//             <last>91.590</last>
	//             <name>ENERGIZER HOLDINGS INC (NEW)</name>
	//             <pchg>1.65</pchg>
	//             <pcls>90.100</pcls>
	//             <rank>49</rank>
	//             <symbol>ENR.A</symbol>
	//             <vl>9327</vl>
	//         </quote>
	//     </quotes>
	//     <error>Success</error>
	// </response>
	// `

	// stockList := parseTopStockQuery(query)

	// fmt.Println(stockList[0])

	// var stockEntry = Stock{
	// 	DayID:        3,
	// 	Monitoring:   true,
	// 	Symbol:       "TAM",
	// 	Bid:          "2.31",
	// 	Ask:          "2.30",
	// 	Last:         "2.28",
	// 	Pchg:         "13.2",
	// 	Pcls:         "2.20",
	// 	Opn:          "2.25",
	// 	Vl:           "2001",
	// 	Pvol:         "3001",
	// 	Volatility12: "1",
	// 	Wk52hi:       "1",
	// 	Wk52hidate:   "1",
	// 	Wk52lo:       "1",
	// 	Wk52lodate:   "1",
	// 	Hi:           "1",
	// 	Low:          "1",
	// 	PrAdp50:      "1",
	// 	PrAdp100:     "1",
	// 	Prchg:        "1",
	// 	Adp50:        "1",
	// 	Adp100:       "1",
	// 	Adv30:        "1",
	// 	Adv90:        "1",
	// }

	// fmt.Println(stockEntry.DayID)
	// insertStock(stockEntry)

	//STORE INSERT STOCKS. Monitoring condition.

	// resultSet := selectMonitoringStock()

	// symbolChecking := "DOH"
	// for i, v := range resultSet {
	// 	if symbolChecking == v {
	// 		fmt.Println("hit", v)
	// 	}
	// 	i++
	// }

	//if not monitoring create query cycle
	//sequential or all the same
	//all the same for set query
	//else individual

	//split process.

	//handle monitor
	//store monitoring
	// var stockEntry = Stock{
	// 	DayID:        3,
	// 	Monitoring:   false,
	// 	Symbol:       symbolChecking,
	// 	Bid:          "2.31",
	// 	Ask:          "2.30",
	// 	Last:         "2.28",
	// 	Pchg:         "13.2",
	// 	Pcls:         "2.20",
	// 	Opn:          "2.25",
	// 	Vl:           "2001",
	// 	Pvol:         "3001",
	// 	Volatility12: "1",
	// 	Wk52hi:       "1",
	// 	Wk52hidate:   "1",
	// 	Wk52lo:       "1",
	// 	Wk52lodate:   "1",
	// 	Hi:           "1",
	// 	Low:          "1",
	// 	PrAdp50:      "1",
	// 	PrAdp100:     "1",
	// 	Prchg:        "1",
	// 	Adp50:        "1",
	// 	Adp100:       "1",
	// 	Adv30:        "1",
	// 	Adv90:        "1",
	// }

	// insertStock(stockEntry)

	// resultSet = selectMonitoringStock()
	// fmt.Println(resultSet)

	//Handle formulate query
	//Query multistock
	// post() Multi dynamic query response, need to set on cycle
	// var stringListToQuery = []string{"AAPL", "AMD", "GE"}
	// queryMultiStockPull(stringListToQuery)

	//Handle cycleManager

	//Now handle
	// deleteStock("TAM")

	//handle query set

	//for a set there will always be three.

	//Query monitoring
	//Select all matching monitored condition

	//Create query cycle.

	//then if not being monitored,
	//Add query cycle.

	// isTimeMonitoringLoop = true
	// go initTimeMonitoring()

	// go say("world")
	// say("hello")

	// processQueryStockSet()

	//processMonitorStockSetQuery

	// processTSPRefresh()
	// processTimelineStart()
	// checkIfHoliday()
	// performWebscrape()
	response := queryWebscrape()

	currentDowValue, pointsChanged, percentageChange := parseDowWebscrape(response)
	insertDow(currentDowValue, pointsChanged, percentageChange)

	fmt.Scanln()
	fmt.Println("done")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// func main() {
// 	go say("world")
// 	say("hello")
// }

func initTimeMonitoring() {
	// time excution process looped on regular basis
	// triggered every few seconds
	i := 0
	for i < 3 { //isTimeMonitoringLoop {
		// timeConditionExecutionProcess()
		// if i == 3 {
		// 	fmt.Println("is false")
		// 	isTimeMonitoringLoop = false
		// }
		fmt.Println("Awesome sauce")
		// time.Sleep(3 * time.Second)
		i++
	}

}
func timeConditionExecutionProcess() {
	currentTime := time.Now()
	fmt.Println(currentTime.Hour())
	fmt.Println(currentTime.Minute())
	fmt.Println(currentTime.Second())

	//Initial monitoring pool
	if currentTime.Hour() == 15 && currentTime.Minute() == 37 {
		fmt.Println("hit")
	}
	//TSP, trigger individual stock query process,
	//Update stock monitor and query decision process
	//trigger update monitoring pool analytics
	if currentTime.Hour() == 15 && currentTime.Minute() == 37 {
		fmt.Println("hit")
	}
	//TSP, trigger individual stock query process,
	//Update stock monitor and query decision process
	//trigger update monitoring pool analytics
	if currentTime.Hour() == 15 && currentTime.Minute() == 37 {
		fmt.Println("hit")
	}
	//TSP, trigger individual stock query process,
	//Update stock monitor and query decision process
	//trigger update monitoring pool analytics
	if currentTime.Hour() == 15 && currentTime.Minute() == 37 {
		fmt.Println("hit")
	}
}

func testing() {
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
