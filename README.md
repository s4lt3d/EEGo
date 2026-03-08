# EEGo — Earth Empires AI Bot

> Go-based AI bot for Earth Empires: a complex browser-based strategy game featuring economy, military, technology, and market systems.

---

## Overview

EEGo is a Go implementation of an autonomous AI bot for [Earth Empires](http://earthempires.com), a browser-based real-time strategy game. The bot manages multiple game mechanics simultaneously:

- **Economy:** Taxation, cash management, population growth, food production
- **Infrastructure:** Building placement (residences, farms, oil rigs, military bases, research labs)
- **Technology:** Military, Medical, Business, Agriculture, Warfare, Spy, SDI advancement
- **Military:** Command troops, jets, tanks, nuclear/chemical/cruise missiles, turrets
- **Markets:** Buy/sell on private and public markets to maximize profit
- **Rank & Competition:** Compete with other bots on the official AI server

---

## Game Mechanics

### Economy System

- **Money/Cash:** Accumulated through taxation and market trading
- **Food Production:** Depends on farms and population consumption
- **Population:** Grows with residences, consumes food, generates tax revenue
- **Tax Rate:** Adjustable (affects revenue vs. happiness)
- **Income Streams:** Taxation, oil production, market profits

### Military & Warfare

- **Forces:** Spies, Troops, Jets, Turrets, Tanks
- **Missiles:** Nuclear, Chemical, Cruise missiles for long-range attacks
- **Warfare Stats:** Military tech level, strategy tech, weapons tech
- **Warfare Types:** Spy attacks, conventional military, missile strikes
- **War Status:** Track active wars and defensive needs

### Technology System

Eight technology tracks with percentages:
- **Military** — Combat effectiveness
- **Medical** — Population health/growth
- **Business** — Economic efficiency
- **Residential** — Population housing capacity
- **Agricultural** — Farm production
- **Warfare** — Strategic combat improvements
- **Spy** — Intelligence and espionage effectiveness
- **SDI** — Strategic Defense Initiative (missile defense)

### Markets

- **Private Market:** Direct trading with specific players
- **Public Market:** Commodity market for resources
- **Transactions:** Buy/sell resources (tracked in json_examples.txt)

---

## Project Structure

```
EEGo/
├── EEGo.go               — Main entry point, bot control flow
├── EEWeb.go              — HTTP/API communication with server
├── EEJsonStructs.go      — JSON parsing for API responses
├── advisor_template.go   — Display templates for game status
├── json_examples.txt     — Example API payloads and responses
└── README.md
```

### Key Files

**EEGo.go**
- Main function orchestrates bot actions
- Calls: GetServer(), GetAdvisor(), GetMarkets(), WebPMBuy(), WebPMSell()
- Entry point for bot decision-making loop

**EEWeb.go**
- HTTP client for server communication
- API functions: GetServer, GetAdvisor, GetPrivateMarket, GetPublicMarket
- Transaction functions: WebCash, WebPMBuy, WebPMSell
- Handles authentication and JSON serialization

**EEJsonStructs.go**
- Go structs for unmarshaling API responses
- Represents: Empire state, market data, server info
- ~15,000 lines of structured data definitions

**advisor_template.go**
- Formatted output templates for game status display
- Shows empire overview, technology, military, expenses
- Server information formatting

---

## Getting Started

### Prerequisites

- Go 1.13+ (or later)
- Earth Empires account
- AI server credentials (cnum - creature number, API key)

### Build & Run

```bash
go build
./EEGo
```

Or run directly:

```bash
go run *.go
```

---

## Configuration

### API Connection

Edit credentials in `EEGo.go`:

```go
func main() {
    GetServer()           // Fetch server info
    GetAdvisor(61)        // Get empire state (cnum=61)
    GetPrivateMarket(61)  // Check private market
    GetPublicMarket(61)   // Check public market
    WebPMBuy(61, 100, 0, 0, 0, 0, 0)    // Buy from market
    WebPMSell(61, 100, 0, 0, 0, 0, 0, 0) // Sell to market
}
```

Replace creature number (`61`) with your account's cnum.

---

## API Functions

### Query Functions

- **GetServer()** — Fetch global server information and round stats
- **GetAdvisor(cnum)** — Get full empire status and statistics
- **GetPrivateMarket(cnum)** — Check open private market orders
- **GetPublicMarket(cnum)** — Check public commodity market prices

### Transaction Functions

- **WebCash(cnum, amount)** — Transfer cash between accounts
- **WebPMBuy(cnum, quantity, resource_types...)** — Purchase from private market
- **WebPMSell(cnum, quantity, resource_types...)** — List for sale on private market

---

## Development Status

⚠️ **In Progress** — Code includes "TODO" comments indicating areas for expansion:
- More sophisticated advisor logic needed
- Additional game mechanics to implement
- Market trading strategies incomplete
- Military decision-making framework

---

## Strategy Ideas

- **Economic:** Monitor tax revenue, optimize production, trade on markets
- **Military:** Research weapon tech, build forces based on threats
- **Expansion:** Balance land acquisition with infrastructure build
- **Diplomacy:** Track alliances, coordinate with other players
- **Defense:** Build military in response to threats

---

## References

- [Earth Empires](http://earthempires.com) — Official game server
- [EE Wiki](https://wiki.earthempires.com) — Game mechanics documentation
- API examples in `json_examples.txt`

---

## License

Copyright © Walter Gordy
