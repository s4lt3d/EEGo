
package main

import "fmt"

type defaultParams struct {
    Username string `json:"username"`
    AI_Key string   `json:"ai_key"`
    Server string   `json:"server"`
    ApiFunction string `json:"api_function"`
}

type serverParams struct {
    defaultParams
}

type advisorParams struct {
    defaultParams
    Cnum int   `json:"cnum"`
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

func (s ServerInfo) String() string {
    return fmt.Sprintf(ServerInfoTemplate, 
        s.Server.Time,                     s.Server.RoundNum, 
        s.Server.ResetStart,               s.Server.TurnRate,
        s.Server.ResetEnd,                 int((s.Server.ResetEnd - s.Server.Time) / 60),
        s.Server.CountriesAllowed,         int((s.Server.ResetEnd - s.Server.Time) / s.Server.TurnRate),
        s.Server.AliveCount)

}

type AdvisorInfo struct {
    Advisor struct {
        Networth int `json:"networth"`
        Land int `json:"land"`
        Population int `json:"pop"`
        Money int `json:"money"`
        PCI float64 `json:"pci"`
        Food int `json:"food"`
        Oil int `json:"oil"`
        EnterprizeZones int `json:"b_ent"`
        ResidentailZones int `json:"b_res"`
        IndustrialZones int `json:"b_indy"`
        MilitaryZones int `json:"b_mb"`
        ResearchZones int `json:"b_lab"`
        FarmZones int `json:"b_farm"`
        OilZones int `json:"b_rig"`
        ConstructionZones int `json:"b_cs"`
        MSpy int `json:"m_spy"`
        MTr int `json:"m_tr"`
        MJ int `json:"m_j"`
        MTu int `json:"m_tu"`
        MTa int `json:"m_ta"`
        TMil int `json:"t_mil"`
        TMed int `json:"t_med"`
        TBus int `json:"t_bus"`
        TRes int `json:"t_res"`
        TAgri int `json:"t_agri"`
        TWar int `json:"t_war"`
        TMs int `json:"t_ms"`
        TWeap int `json:"t_weap"`
        TIndy int `json:"t_indy"`
        TSpy int `json:"t_spy"`
        TSdi int `json:"t_sdi"`
        Empty int `json:"empty"`
        NwMil int `json:"nw_mil"`
        NwTech int `json:"nw_tech"`
        NwLand int `json:"nw_land"`
        NwOther int `json:"nw_other"`
        NwMarket int `json:"nw_market"`
        Foodpro int `json:"foodpro"`
        Foodcon int `json:"foodcon"`
        Foodnet int `json:"foodnet"`
        Oilpro int `json:"oilpro"`
        Tpt int `json:"tpt"`
        Taxes int `json:"taxes"`
        Expenses int `json:"expenses"`
        Corruption int `json:"corruption"`
        ExpensesAlliance int `json:"expenses_alliance"`
        ExpensesLand int `json:"expenses_land"`
        ExpensesMil int `json:"expenses_mil"`
        ExpenseSpy int `json:"expense_spy"`
        ExpenseTr int `json:"expense_tr"`
        ExpenseJ int `json:"expense_j"`
        ExpenseTu int `json:"expense_tu"`
        ExpenseTa int `json:"expense_ta"`
        Income int `json:"income"`
        Cashing int `json:"cashing"`
        TotJ int `json:"tot_j"`
        TTot int `json:"t_tot"`
        PtMil int `json:"pt_mil"`
        PtMed int `json:"pt_med"`
        PtBus int `json:"pt_bus"`
        PtRes int `json:"pt_res"`
        PtAgri float64 `json:"pt_agri"`
        PtWar float64 `json:"pt_war"`
        PtMs int `json:"pt_ms"`
        PtWeap int `json:"pt_weap"`
        PtIndy int `json:"pt_indy"`
        PtSpy int `json:"pt_spy"`
        PtSdi int `json:"pt_sdi"`
        Government string `json:"govt"`
        Taxrate string `json:"taxrate"`
        PsTr string `json:"ps_tr"`
        PsJ string `json:"ps_j"`
        PsTa string `json:"ps_ta"`
        Bpt int `json:"bpt"`
        BuildCost int `json:"build_cost"`
        Cnum string `json:"cnum"`
        ExploreRate int `json:"explore_rate"`
        Turns int `json:"turns"`
        TurnsStored int `json:"turns_stored"`
        TurnsPlayed int `json:"turns_played"`
        GTax int `json:"g_tax"`
        Protection string `json:"protection"`
        ProSpy string `json:"pro_spy"`
        ProTr string `json:"pro_tr"`
        ProJ string `json:"pro_j"`
        ProTu string `json:"pro_tu"`
        ProTa string `json:"pro_ta"`
    } `json:"ADVISOR"`
}

func (s AdvisorInfo) String() string {
    return fmt.Sprintf(AdvisorTemplate, 
       s.Advisor.Cnum,                  s.Advisor.Government,
       s.Advisor.Turns, s.Advisor.Money, s.Advisor.Food, s.Advisor.Networth,
       s.Advisor.Turns,                 s.Advisor.Taxes,
       s.Advisor.TurnsPlayed,           s.Advisor.Taxrate,
       s.Advisor.TurnsStored,           s.Advisor.PCI,
       0,                               0,
       s.Advisor.Networth,              s.Advisor.Income,
       s.Advisor.Land,                  s.Advisor.Cashing,
       s.Advisor.Money,                 s.Advisor.Food,
       s.Advisor.Population,            s.Advisor.Foodpro)

}