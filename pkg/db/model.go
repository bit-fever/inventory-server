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

type Exchange struct {
	Id         uint   `json:"id"`
	CurrencyId uint   `json:"currencyId"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Timezone   string `json:"timezone"`
	Url        string `json:"url"`
}

//=============================================================================

type Connection struct {
	Common
	Username             string `json:"username"`
	Code                 string `json:"code"`
	Name                 string `json:"name"`
	SystemCode           string `json:"systemCode"`
	SystemName           string `json:"systemName"`
	SystemConfig         string `json:"systemConfig"`
	ConnectionCode       string `json:"connectionCode"`
	SupportsData         bool   `json:"supportsData"`
	SupportsBroker       bool   `json:"supportsBroker"`
	SupportsMultipleData bool   `json:"supportsMultipleData"`
	SupportsInventory    bool   `json:"supportsInventory"`
}

//=============================================================================

type Portfolio struct {
	Common
	ParentId  uint    `json:"parentId"`
	Username  string  `json:"username"`
	Name      string  `json:"name"`
}

//=============================================================================

type ProductData struct {
	Common
	ConnectionId uint     `json:"connectionId"`
	ExchangeId   uint     `json:"exchangeId"`
	Username     string   `json:"username"`
	Symbol       string   `json:"symbol"`
	Name         string   `json:"name"`
	Increment    float64  `json:"increment"`
	MarketType   string   `json:"marketType"`
	ProductType  string   `json:"productType"`
	LocalClass   string   `json:"localClass"`
}

//=============================================================================

type ProductDataFull struct {
	ProductData
	ConnectionCode  string  `json:"connectionCode,omitempty"`
	SystemCode      string  `json:"systemCode,omitempty"`
	ExchangeCode    string  `json:"exchangeCode,omitempty"`
}

//=============================================================================

type ProductBroker struct {
	Common
	ConnectionId     uint     `json:"connectionId"`
	ExchangeId       uint     `json:"exchangeId"`
	Username         string   `json:"username"`
	Symbol           string   `json:"symbol"`
	Name             string   `json:"name"`
	PointValue       float32  `json:"pointValue"`
	CostPerTrade     float32  `json:"costPerTrade"`
	MarginValue      float32  `json:"marginValue"`
	MarketType       string   `json:"marketType"`
	ProductType      string   `json:"productType"`
	LocalClass       string   `json:"localClass"`
}

//=============================================================================

type ProductBrokerFull struct {
	ProductBroker
	CurrencyCode    string  `json:"currencyCode,omitempty"`
	ConnectionCode  string  `json:"connectionCode,omitempty"`
	ExchangeCode    string  `json:"exchangeCode,omitempty"`
}

//=============================================================================

type InstrumentData struct {
	Id               uint    `json:"id" gorm:"primaryKey"`
	ProductDataId    uint    `json:"productDataId"`
	ContractId       uint    `json:"contractId"`
	Symbol           string  `json:"symbol"`
	Name             string  `json:"name"`
	ExpirationDate   int     `json:"expirationDate"`
	IsContinuous     bool    `json:"isContinuous"`
}

//=============================================================================

type InstrumentBroker struct {
	Id               uint    `json:"id" gorm:"primaryKey"`
	ProductBrokerId  uint    `json:"productBrokerId"`
	ContractId       uint    `json:"contractId"`
	Symbol           string  `json:"symbol"`
	Name             string  `json:"name"`
	ExpirationDate   int     `json:"expirationDate"`
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
	ProductDataId     uint    `json:"productDataId"`
	ProductBrokerId   uint    `json:"productBrokerId"`
	TradingSessionId  uint    `json:"tradingSessionId"`
	Username          string  `json:"username"`
	WorkspaceCode     string  `json:"workspaceCode"`
	Name              string  `json:"name"`
}

//=============================================================================

type TradingSystemFull struct {
	TradingSystem
	DataSymbol     string `json:"dataSymbol,omitempty"`
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
func (Exchange)         TableName() string { return "exchange" }
func (Connection)       TableName() string { return "connection" }
func (Portfolio)        TableName() string { return "portfolio" }
func (ProductData)      TableName() string { return "product_data" }
func (ProductBroker)    TableName() string { return "product_broker" }
func (InstrumentData)   TableName() string { return "instrument_data" }
func (InstrumentBroker) TableName() string { return "instrument_broker" }
func (TradingSession)   TableName() string { return "trading_session" }
func (TradingSystem)    TableName() string { return "trading_system" }

//=============================================================================
