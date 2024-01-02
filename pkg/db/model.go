//=============================================================================
/*
Copyright © 2023 Andrea Carboni andrea.carboni71@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
//=============================================================================

package db

import (
	"time"
)

//=============================================================================

type Common struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

//=============================================================================

type Currency struct {
	Id   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

//=============================================================================

type Connection struct {
	Common
	Username              string `json:"username"`
	Code                  string `json:"code"`
	Name                  string `json:"name"`
	SystemCode            string `json:"systemCode"`
	SystemName            string `json:"systemName"`
	SystemConfig          string `json:"systemConfig"`
	ConnectionCode        string `json:"connectionCode"`
	SupportsFeed          bool   `json:"supportsFeed"`
	SupportsBroker        bool   `json:"supportsBroker"`
	SupportsMultipleFeeds bool   `json:"supportsMultipleFeeds"`
	SupportsInventory     bool   `json:"supportsInventory"`
}

//=============================================================================

type Portfolio struct {
	Common
	ParentId  uint    `json:"parentId"`
	Username  string  `json:"username"`
	Name      string  `json:"name"`
}

//=============================================================================

type ProductFeed struct {
	Common
	ConnectionId uint     `json:"connectionId"`
	Username     string   `json:"username"`
	Symbol       string   `json:"symbol"`
	Name         string   `json:"name"`
	PriceScale   int      `json:"priceScale"`
	MinMovement  float32  `json:"minMovement"`
	MarketType   string   `json:"marketType"`
	ProductType  string   `json:"productType"`
	Exchange     string   `json:"exchange"`
}

//=============================================================================


type ProductFeedFull struct {
	ProductFeed
	ConnectionCode  string  `json:"connectionCode,omitempty"`
}

//=============================================================================

type ProductBroker struct {
	Common
	ConnectionId     uint       `json:"connectionId"`
	Username         string     `json:"username"`
	Symbol           string     `json:"symbol"`
	Name             string     `json:"name"`
	PointValue       float32    `json:"pointValue"`
	CostPerTrade     float32    `json:"costPerTrade"`
	MarginValue      float32    `json:"marginValue"`
	CurrencyId       uint       `json:"currencyId"`
	MarketType       string     `json:"marketType"`
	ProductType      string     `json:"productType"`
	Exchange         string     `json:"exchange"`
}

//=============================================================================

type ProductBrokerFull struct {
	ProductBroker
	CurrencyCode    string  `json:"currencyCode,omitempty"`
	ConnectionCode  string  `json:"connectionCode,omitempty"`
}

//=============================================================================

type InstrumentFeed struct {
	Id               uint      `json:"id" gorm:"primaryKey"`
	ProductFeedId    uint       `json:"productFeedId"`
	Symbol           string     `json:"symbol"`
	Name             string     `json:"name"`
	ExpirationDate   time.Time  `json:"expirationDate"`
	IsContinuous     bool       `json:"isContinuous"`
}

//=============================================================================

type InstrumentBroker struct {
	Id               uint       `json:"id" gorm:"primaryKey"`
	ProductBrokerId  uint       `json:"productBrokerId"`
	Symbol           string     `json:"symbol"`
	Name             string     `json:"name"`
	ExpirationDate   time.Time  `json:"expirationDate"`
}

//=============================================================================

type TradingSession struct {
	Common
	Username  string `json:"username"`
	Name      string `json:"name"`
	Config    string `json:"config"`
}

//=============================================================================

type TradingSystem struct {
	Common
	PortfolioId       uint    `json:"portfolioId"`
	ProductFeedId     uint    `json:"productFeedId"`
	ProductBrokerId   uint    `json:"productBrokerId"`
	TradingSessionId  uint    `json:"tradingSessionId"`
	Username          string  `json:"username"`
	WorkspaceCode     string  `json:"workspaceCode"`
	Name              string  `json:"name"`
}

//=============================================================================

type TradingSystemFull struct {
	TradingSystem
	FeedSymbol     string `json:"feedSymbol,omitempty"`
	BrokerSymbol   string `json:"brokerSymbol,omitempty"`
	PortfolioName  string `json:"portfolioName,omitempty"`
	TradingSession string `json:"tradingSession,omitempty"`
}

//=============================================================================
//===
//=== Table names
//===
//=============================================================================

func (Currency)         TableName() string { return "currency" }
func (Connection)       TableName() string { return "connection" }
func (Portfolio)        TableName() string { return "portfolio" }
func (ProductFeed)      TableName() string { return "product_feed" }
func (ProductBroker)    TableName() string { return "product_broker" }
func (InstrumentFeed)   TableName() string { return "instrument_feed" }
func (InstrumentBroker) TableName() string { return "instrument_broker" }
func (TradingSession)   TableName() string { return "trading_session" }
func (TradingSystem)    TableName() string { return "trading_system" }

//=============================================================================
