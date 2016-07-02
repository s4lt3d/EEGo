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

type defaultParams struct {
    Username string `json:"username"`
    AI_Key string   `json:"ai_key"`
    Server string   `json:"server"`
    ApiFunction string `json:"api_function"`
}

type serverParams struct {
    defaultParams
}

type ServerInfo struct {
    Server struct {
        RoundNum int `json:"round_num"`
        ResetStart int `json:"reset_start"`
        ResetEnd int `json:"reset_end"`
        TurnRate int `json:"turn_rate"`
        CountriesAllowed string `json:"countries_allowed"`
        AliveCount int `json:"alive_count"`
        CnumList struct {
            Alive []int `json:"alive"`
            Dead []interface{} `json:"dead"`
        } `json:"cnum_list"`
        Time int `json:"time"`
    } `json:"SERVER_INFO"`
}

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
    fmt.Println("response Body:", string(body))
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
    fmt.Println(body)
    
    server := ServerInfo{}

    _ = json.Unmarshal([]byte(body), &server)

    fmt.Println(server)
}
