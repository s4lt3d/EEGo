// ****************************************************************************
// EEWeb.go
// A collection of functions to get data in and out of the ee server
// ****************************************************************************

package main

import (
        "fmt"
        "net/http"
        "net/url"
        "io/ioutil"
        "encoding/json")

type defaultParams struct {
    Username string `json:"username"`
    Ai_key string   `json:"ai_key"`
    Server string   `json:"server"`
    Api_function string `json:"api_function"`
}

type serverParams struct {
    defaultParams
}

// ****************************************************************************
// doPost - All web communication for the game goes this this function
// ****************************************************************************
func doPost() {
    resp, err := http.PostForm("http://qz.earthempires.com/api",
                             url.Values{"username":{"salted"}, 
                                        "ai_key":{"49ee125ad5e9a3b81dfb771ac0d3d2fb"},
                                        "server":{"ai"},
                                        "api_function":{"server"}})
    if err != nil {
        fmt.Println(err)
    }

    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}

func GetServer() {
    doPost()
}

func TestJSON() {
    s := serverParams{ 
            defaultParams {
                Username: "s4lt3d", 
                Ai_key: "49ee125ad5e9a3b81dfb771ac0d3d2fb", 
                Server: "ai", 
                Api_function: "server"}}

    b, _ := json.Marshal(s)
    fmt.Println(string(b))
}