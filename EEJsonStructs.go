
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
        SpyForces int `json:"m_spy"`
        TroopForces int `json:"m_tr"`
        JetForces int `json:"m_j"`
        TurretForces int `json:"m_tu"`
        TankForces int `json:"m_ta"`
        MilitaryTech int `json:"t_mil"`
        MedicalTech int `json:"t_med"`
        BusinessTech int `json:"t_bus"`
        ResidentalTech int `json:"t_res"`
        AgricultureTech int `json:"t_agri"`
        WarfareTech int `json:"t_war"`
        MilitaryStrategyTech int `json:"t_ms"`
        WeaponTech int `json:"t_weap"`
        IndustrialTech int `json:"t_indy"`
        SpyTech int `json:"t_spy"`
        SDITech int `json:"t_sdi"`
        EmptyLand int `json:"empty"`
        NwMil int `json:"nw_mil"`
        NwTech int `json:"nw_tech"`
        NwLand int `json:"nw_land"`
        NwOther int `json:"nw_other"`
        NwMarket int `json:"nw_market"`
        FoodProduced int `json:"foodpro"`
        FoodConsumed int `json:"foodcon"`
        FoodNet int `json:"foodnet"`
        OilProduced int `json:"oilpro"`
        TechPerTurn int `json:"tpt"`
        TaxRevenue int `json:"taxes"`
        Expenses int `json:"expenses"`
        Corruption int `json:"corruption"`
        ExpensesAlliance int `json:"expenses_alliance"`
        LandUpkeep int `json:"expenses_land"`
        MilitaryUpkeep int `json:"expenses_mil"`
        SpyUpkeep int `json:"expense_spy"`
        TroopUpkeep int `json:"expense_tr"`
        JetUpkeep int `json:"expense_j"`
        TurretUpkeep int `json:"expense_tu"`
        TankUpkeep int `json:"expense_ta"`
        Income int `json:"income"`
        Cashing int `json:"cashing"`
        TotJ int `json:"tot_j"`
        TechTotal int `json:"t_tot"`
        MilitaryTechPercentage float64 `json:"pt_mil"`
        MedicalTechPercentage float64 `json:"pt_med"`
        BusinessTechPercentage float64 `json:"pt_bus"`
        ResidentalTechPercentage float64 `json:"pt_res"`
        AgricultureTechPercentage float64 `json:"pt_agri"`
        WarfareTechPercentage float64 `json:"pt_war"`
        MilitaryStrategyTechPercentage float64 `json:"pt_ms"`
        WeaponTechPercentage float64 `json:"pt_weap"`
        IndustrialTechPercentage float64 `json:"pt_indy"`
        SpyTechPercentage float64 `json:"pt_spy"`
        SDITechPercentage float64 `json:"pt_sdi"`
        Government string `json:"govt"`
        TaxRate int `json:"taxrate,string"` // json incorrectly quotes an int
        PsTr string `json:"ps_tr"`
        PsJ string `json:"ps_j"`
        PsTa string `json:"ps_ta"`
        BuildingsPerTurn int `json:"bpt"`
        BuildCost int `json:"build_cost"`
        Cnum string `json:"cnum"`
        ExploreRate int `json:"explore_rate"`
        Turns int `json:"turns"`
        TurnsStored int `json:"turns_stored"`
        TurnsPlayed int `json:"turns_played"`
        GTax int `json:"g_tax"`
        Protection string `json:"protection"`
        SpyProduced string `json:"pro_spy"`
        TroopProduced string `json:"pro_tr"`
        JetProduced string `json:"pro_j"`
        TurretProduced string `json:"pro_tu"`
        TankProduced string `json:"pro_ta"`
    } `json:"ADVISOR"`
}

func (s AdvisorInfo) String() string {
    return fmt.Sprintf(AdvisorTemplate, 
       s.Advisor.Cnum,                          s.Advisor.Government,
       s.Advisor.Turns, s.Advisor.Money,        s.Advisor.Food, s.Advisor.Networth,
       s.Advisor.Turns,                         s.Advisor.TaxRevenue,
       s.Advisor.TurnsPlayed,                   s.Advisor.TaxRate,
       s.Advisor.TurnsStored,                   s.Advisor.PCI,
       0,                                       0,
       s.Advisor.Networth,                      s.Advisor.Income,
       s.Advisor.Land,                          s.Advisor.Cashing,
       s.Advisor.Money,                         s.Advisor.Food,
       s.Advisor.Population,                    s.Advisor.FoodProduced,
       0,                                       s.Advisor.FoodConsumed,
       0,                                       0,
                                                s.Advisor.FoodNet,
                                                s.Advisor.Oil,
                                                s.Advisor.OilProduced,
        s.Advisor.EnterprizeZones,              s.Advisor.BuildingsPerTurn,
        s.Advisor.ResidentailZones,             s.Advisor.ExploreRate,
        s.Advisor.IndustrialZones,              s.Advisor.TechPerTurn,
        s.Advisor.MilitaryZones,                
        s.Advisor.ResearchZones,
        s.Advisor.FarmZones,
        s.Advisor.OilZones,                     s.Advisor.SpyForces,
        s.Advisor.ConstructionZones,            s.Advisor.TroopForces,
        s.Advisor.Land - s.Advisor.EmptyLand,   s.Advisor.JetForces,
        s.Advisor.EmptyLand,                    s.Advisor.TurretForces,
                                                s.Advisor.TankForces,
                                                0,
                                                0,
        s.Advisor.MilitaryTech,         s.Advisor.MilitaryTechPercentage,                0,
        s.Advisor.MedicalTech,          s.Advisor.MedicalTechPercentage,        
        s.Advisor.BusinessTech,         s.Advisor.BusinessTechPercentage,        
        s.Advisor.ResidentalTech,       s.Advisor.ResidentalTechPercentage,            
        s.Advisor.AgricultureTech,      s.Advisor.AgricultureTechPercentage,        s.Advisor.Expenses,    
        s.Advisor.WarfareTech,          s.Advisor.WarfareTechPercentage,            s.Advisor.MilitaryUpkeep,
        s.Advisor.MilitaryStrategyTech, s.Advisor.MilitaryStrategyTechPercentage,   s.Advisor.SpyUpkeep,        
        s.Advisor.WeaponTech,           s.Advisor.WeaponTechPercentage,             s.Advisor.TroopUpkeep,
        s.Advisor.IndustrialTech,       s.Advisor.IndustrialTechPercentage,         s.Advisor.JetUpkeep,
        s.Advisor.SpyTech,              s.Advisor.SpyTechPercentage,                s.Advisor.TurretUpkeep,
        s.Advisor.SDITech,              s.Advisor.SDITechPercentage,                s.Advisor.TankUpkeep,
                                                0,
        s.Advisor.TechTotal,                    s.Advisor.LandUpkeep,
                                                0)

}