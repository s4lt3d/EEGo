package main

const AdvisorTemplate = 
`*******************************************************************************
* The Status of cnum : %4s                                    Government: %s  *
*                                                                             *
* Turns: %3d   Money: $%11d   Food: %10d   Networth: $%10d * 
*******************************************************************************

The Basics                                  Current Status
*************************************       ***********************************
Turns Left            %15d       Tax Revenue              $%9d
Turns Taken           %15d         Tax Rate               %9d%%
Turns Stored          %15d         Per Capital Income     $%9.2f
Rank                  %15d       Expenses                 $%9d
Networth             $%15d       Net Income               $%9d
Land                  %15d         Cashing                $%9d
Money                $%15d       Food                      %9d
Population            %15d         Production              %9d
At War                %15d         Consumption             %9d
GDI Member            %15d         Decay                   %9d
                                              Net Change              %9d
Land Distribution                           Oil                       %9d
*************************************         Production              %9d
Enterprise Zones      %15d       Building Rate             %9d
Residences            %15d       Exploration Rate          %9d
Industrial Complexes  %15d       Tech Rate                 %9d
Military Bases        %15d       
Research Labs         %15d       Military Forces
Farms                 %15d       ***********************************
Oil Rigs              %15d       Spies                     %9d
Construction Sites    %15d       Troops                    %9d
Used Land             %15d       Jets                      %9d
Unused Land           %15d       Turrets                   %9d
                                            Tanks                     %9d
Technology                                  Nuclear Missiles          %9d
*************************************       Chemical Missiles         %9d
Military           %8d   %7.3f%%      Cruise Missiles           %9d
Medical            %8d   %7.3f%%       
Business           %8d   %7.3f%%       Expense Breakdown
Residential        %8d   %7.3f%%       **********************************
Agricultural       %8d   %7.3f%%       Expenses                $%9d
Warfare            %8d   %7.3f%%         Military              $%9d
Military Strategy  %8d   %7.3f%%           Spies               $%9d
Weapons            %8d   %7.3f%%           Troops              $%9d
Industrial         %8d   %7.3f%%           Jets                $%9d
Spy                %8d   %7.3f%%           Turrets             $%9d
SDI                %8d   %7.3f%%           Tanks               $%9d
                                            Alliance/GDI             $%9d
Total       %15d                 Land                     $%9d
                                            Corruption               $%9d`



//##########################################################################
const ServerInfoTemplate = 
`
Server Information                       Round Information 
*************************************    ***********************************
Current Time:              %10d    Round Num:               %10d
Reset Start:               %10d    Turn Rate:               %10d
Reset End:                 %10d    Time Left (min):         %10d
Countries Allowed:         %10s    Turns Left:              %10d

Country Infomation
*************************************
Alive Count:               %10d
`

//##########################################################################
const PMInfoTemplate = 
`
                                Private Market 
*******************************************************************************
Units                   Buy Price              Sell Price             Available
*******************************************************************************
Spies                          0              %10d                      0
Troops                %10d              %10d             %10d
Jets                  %10d              %10d             %10d
Turrets               %10d              %10d             %10d
Tanks                 %10d              %10d             %10d
Food                  %10d              %10d             %10d
Oil                   %10d              %10d             %10d
`   
const PublicMarketTemplate = 
`
                                Public Market 
*******************************************************************************
Units                   Buy Price              Sell Price             Available
*******************************************************************************
Troops                %10d              %10d             %10d
Jets                  %10d              %10d             %10d
Turrets               %10d              %10d             %10d
Tanks                 %10d              %10d             %10d
Food                  %10d              %10d             %10d
Oil                   %10d              %10d             %10d
******************************************************************************
Military Tech         %10d              %10d             %10d
Medical Tech          %10d              %10d             %10d
Business Tech         %10d              %10d             %10d
Residental Tech       %10d              %10d             %10d
Agriculture Tech      %10d              %10d             %10d
Warfare Tech          %10d              %10d             %10d
MilitaryStrategy Tech %10d              %10d             %10d
Weapon Tech           %10d              %10d             %10d
Industrial Tech       %10d              %10d             %10d
Spy Tech              %10d              %10d             %10d
SDI Tech              %10d              %10d             %10d
`   