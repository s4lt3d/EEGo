// ****************************************************************************
// EEWeb.go
// A collection of functions to get data in and out of the ee server
// ****************************************************************************
package main

import ("fmt"
        "net/http"
        "io/ioutil"
        "bytes"
        "encoding/json")


// ****************************************************************************
// doPost - All web communication for the game goes this this function
// ****************************************************************************
func doPost(payload string) string  {
    url := "http://qz.earthempires.com/api"

    var jsonStr = []byte("api_payload=" + payload)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
  //  fmt.Println("response Body:", string(body))
    return string(body)
}

func GetServer() {
    s := serverParams{ 
            defaultParams {
                Username: "salted", 
                AI_Key: "49ee125ad5e9a3b81dfb771ac0d3d2fb", 
                Server: "ai", 
                ApiFunction: "server"}}

    b, _ := json.Marshal(s)

    body := doPost(string(b))
    //fmt.Println(body)
    
    server := ServerInfo{}

    _ = json.Unmarshal([]byte(body), &server)

    fmt.Println(server)
}


func GetAdvisor(cnum int) {
    s := advisorParams{ 
            Cnum: cnum,
            defaultParams: defaultParams {
                Username: "salted", 
                AI_Key: "49ee125ad5e9a3b81dfb771ac0d3d2fb", 
                Server: "ai", 
                ApiFunction: "advisor"}}

    b, _ := json.Marshal(s)

    body := doPost(string(b))
    //fmt.Println(body)
    
    advisor := AdvisorInfo{}

    _ = json.Unmarshal([]byte(body), &advisor)

    fmt.Println(advisor)
}

func GetPrivateMarket(cnum int) {
    s := pmInfoParams{ 
            Cnum: cnum,
            defaultParams: defaultParams {
                Username: "salted", 
                AI_Key: "49ee125ad5e9a3b81dfb771ac0d3d2fb", 
                Server: "ai", 
                ApiFunction: "pm_info"}}

    b, _ := json.Marshal(s)

    body := doPost(string(b))
    //fmt.Println(body)
    
    pminfo := PMInfo{}

    _ = json.Unmarshal([]byte(body), &pminfo)

    fmt.Println(pminfo)
}

func GetPublicMarket(cnum int) {
    s := marketParams{ 
            Cnum: cnum,
            defaultParams: defaultParams {
                Username: "salted", 
                AI_Key: "49ee125ad5e9a3b81dfb771ac0d3d2fb", 
                Server: "ai", 
                ApiFunction: "market"}}

    b, _ := json.Marshal(s)

    body := doPost(string(b))
    fmt.Println(body)
    
    publicmarket := PublicMarket{}

    _ = json.Unmarshal([]byte(body), &publicmarket)

    fmt.Println(publicmarket)
}