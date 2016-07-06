// ****************************************************************************
// EEWeb.go
// A collection of functions to get data in and out of the ee server
// ****************************************************************************
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ****************************************************************************
// doPost - All web communication for the game goes this this function
// ****************************************************************************
func doPost(payload string) (json string, success bool) {
	url := "http://qz.earthempires.com/api"
	_success := true
	var jsonStr = []byte("api_payload=" + payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		_success = false
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//  fmt.Println("response Body:", string(body))
	return string(body), _success
}

//GetServer gets the server status from one player from ee
func GetServer() {
	s := serverParams{
		defaultParams{
			Username:    "salted",
			AIKey:       "49ee125ad5e9a3b81dfb771ac0d3d2fb",
			Server:      "ai",
			ApiFunction: "server"}}

	b, _ := json.Marshal(s)

	body, _ := doPost(string(b))
	//fmt.Println(body)

	server := ServerInfo{}

	_ = json.Unmarshal([]byte(body), &server)

	fmt.Println(server)
}

//GetAdvisor gets the advisor from ee
func GetAdvisor(cnum int) {
	s := advisorParams{
		Cnum: cnum,
		defaultParams: defaultParams{
			Username:    "salted",
			AIKey:       "49ee125ad5e9a3b81dfb771ac0d3d2fb",
			Server:      "ai",
			ApiFunction: "advisor"}}

	b, _ := json.Marshal(s)

	body, _ := doPost(string(b))
	//fmt.Println(body)

	advisor := AdvisorInfo{}

	_ = json.Unmarshal([]byte(body), &advisor)

	fmt.Println(advisor)
}

//CreateCountry creates and names a country
func CreateCountry(cnum int, cname string) {
	s := createParams{
		Cnum:  cnum,
		Cname: cname,
		defaultParams: defaultParams{
			Username:    "salted",
			AIKey:       "49ee125ad5e9a3b81dfb771ac0d3d2fb",
			Server:      "ai",
			ApiFunction: "create"}}

	b, _ := json.Marshal(s)

	body, _ := doPost(string(b))
	fmt.Println(body)

	//    publicmarket := PublicMarket{}

	//    _ = json.Unmarshal([]byte(body), &publicmarket)

	//    fmt.Println(publicmarket)
}

//GetPrivateMarket returns the private market and success
func GetPrivateMarket(cnum int) {
	s := pmInfoParams{
		Cnum: cnum,
		defaultParams: defaultParams{
			Username:    "salted",
			AIKey:       "49ee125ad5e9a3b81dfb771ac0d3d2fb",
			Server:      "ai",
			ApiFunction: "pm_info"}}

	b, _ := json.Marshal(s)

	body, _ := doPost(string(b))

	pminfo := PMInfo{}

	_ = json.Unmarshal([]byte(body), &pminfo)

	fmt.Println(pminfo)
}

// GetPublicMarket returns public market and success
func GetPublicMarket(cnum int) (market PublicMarket, success bool) {
	s := marketParams{
		Cnum: cnum,
		defaultParams: defaultParams{
			Username:    "salted",
			AIKey:       "49ee125ad5e9a3b81dfb771ac0d3d2fb",
			Server:      "ai",
			ApiFunction: "market"}}

	b, _ := json.Marshal(s)

	body, _ := doPost(string(b))

	publicmarket := PublicMarket{}

	_ = json.Unmarshal([]byte(body), &publicmarket)

	fmt.Println(publicmarket)
	return publicmarket, true
}

//WebExplore uses turns on ee to explore
func WebExplore(cnum int, turns int) {
	s := exploreParams{
		Cnum:  cnum,
		Turns: turns,
		defaultParams: defaultParams{
			Username:    "salted",
			AIKey:       "49ee125ad5e9a3b81dfb771ac0d3d2fb",
			Server:      "ai",
			ApiFunction: "explore"}}

	b, _ := json.Marshal(s)

	body, _ := doPost(string(b))
	fmt.Println(body)

	//    publicmarket := PublicMarket{}

	//    _ = json.Unmarshal([]byte(body), &publicmarket)

	//    fmt.Println(publicmarket)
}

//WebBuild builds on the ee server
func WebBuild(cnum int, residentialZones int, enterprizeZones int, industrialZones int,
	militaryZones int, researchZones int, farmZones int, oilZones int,
	constructionZones int) {
	s := buildParams{
		Cnum: cnum,
		build: build{
			ResidentialZones:  residentialZones,
			EnterprizeZones:   enterprizeZones,
			IndustrialZones:   industrialZones,
			MilitaryZones:     militaryZones,
			ResearchZones:     researchZones,
			FarmZones:         farmZones,
			OilZones:          oilZones,
			ConstructionZones: constructionZones},
		defaultParams: defaultParams{
			Username:    "salted",
			AIKey:       "49ee125ad5e9a3b81dfb771ac0d3d2fb",
			Server:      "ai",
			ApiFunction: "build"}}

	b, _ := json.Marshal(s)
	fmt.Println(b)
	body, _ := doPost(string(b))
	//  body := `{"BUILD":{"built":{"res":5,"cs":1},"cost":11160,"bpt":5,"turns":{"1":{"taxrevenue":7000,"foodproduced":46,"popgrowth":13,"foodconsumed":30,"expenses":438,"troopsproduced":0,"jetsproduced":0,"turretsproduced":0,"tanksproduced":0,"spiesproduced":0},"2":{"taxrevenue":6098,"foodproduced":46,"popgrowth":9,"foodconsumed":31,"expenses":438,"troopsproduced":0,"jetsproduced":0,"turretsproduced":0,"tanksproduced":0,"spiesproduced":0}},"tpt":3}}`
	fmt.Println(body)

	buildTurn := BuildTurn{}

	_ = json.Unmarshal([]byte(body), &buildTurn)

	fmt.Println(buildTurn)
}

//WebCash uses turns on ee to generate cash
func WebCash(cnum int, turns int) {
	s := cashParams{
		Cnum:  cnum,
		Turns: turns,
		defaultParams: defaultParams{
			Username:    "salted",
			AIKey:       "49ee125ad5e9a3b81dfb771ac0d3d2fb",
			Server:      "ai",
			ApiFunction: "cash"}}

	b, _ := json.Marshal(s)

	body, _ := doPost(string(b))
	fmt.Println(body)

	//    publicmarket := PublicMarket{}

	//    _ = json.Unmarshal([]byte(body), &publicmarket)

	//    fmt.Println(publicmarket)
}

//WebGovernment changes governments
func WebGovernment(cnum int, government string) {
	s := govtParams{
		Cnum: cnum,
		Govt: government,
		defaultParams: defaultParams{
			Username:    "salted",
			AIKey:       "49ee125ad5e9a3b81dfb771ac0d3d2fb",
			Server:      "ai",
			ApiFunction: "govt"}}

	b, _ := json.Marshal(s)

	body, _ := doPost(string(b))
	fmt.Println(body)

	//    publicmarket := PublicMarket{}

	//    _ = json.Unmarshal([]byte(body), &publicmarket)

	//    fmt.Println(publicmarket)
}

//WebTech techs on the ee server
func WebTech(cnum int, militaryTech int, medicalTech int, businessTech int,
	residentalTech int, agricultureTech int, warfareTech int, militaryStrategyTech int,
	weaponTech int, industrialTech int, spyTech int, sdiTech int) {
	s := techParams{
		Cnum: cnum,
		tech: tech{
			MilitaryTech:         militaryTech,
			MedicalTech:          medicalTech,
			BusinessTech:         businessTech,
			ResidentalTech:       residentalTech,
			AgricultureTech:      agricultureTech,
			WarfareTech:          warfareTech,
			MilitaryStrategyTech: militaryStrategyTech,
			WeaponTech:           weaponTech,
			IndustrialTech:       industrialTech,
			SpyTech:              spyTech,
			SDITech:              sdiTech},

		defaultParams: defaultParams{
			Username:    "salted",
			AIKey:       "49ee125ad5e9a3b81dfb771ac0d3d2fb",
			Server:      "ai",
			ApiFunction: "build"}}

	b, _ := json.Marshal(s)
	fmt.Println(b)
	body, _ := doPost(string(b))
	//  body := `{"BUILD":{"built":{"res":5,"cs":1},"cost":11160,"bpt":5,"turns":{"1":{"taxrevenue":7000,"foodproduced":46,"popgrowth":13,"foodconsumed":30,"expenses":438,"troopsproduced":0,"jetsproduced":0,"turretsproduced":0,"tanksproduced":0,"spiesproduced":0},"2":{"taxrevenue":6098,"foodproduced":46,"popgrowth":9,"foodconsumed":31,"expenses":438,"troopsproduced":0,"jetsproduced":0,"turretsproduced":0,"tanksproduced":0,"spiesproduced":0}},"tpt":3}}`
	fmt.Println(body)

	buildTurn := BuildTurn{}

	_ = json.Unmarshal([]byte(body), &buildTurn)

	fmt.Println(buildTurn)
}
