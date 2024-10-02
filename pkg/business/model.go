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

package business

import (
	"github.com/bit-fever/inventory-server/pkg/db"
	"github.com/bit-fever/sick-engine/session"
)

//=============================================================================

type ConnectionSpec struct {
	Code         string `json:"code"       binding:"required"`
	Name         string `json:"name"       binding:"required"`
	SystemCode   string `json:"systemCode" binding:"required"`
	SystemConfig string `json:"systemConfig"`
}

//=============================================================================

type TradingSystemSpec struct {
	PortfolioId      uint   `json:"portfolioId"       binding:"required"`
	DataProductId    uint   `json:"dataProductId"     binding:"required"`
	BrokerProductId  uint   `json:"brokerProductId"   binding:"required"`
	TradingSessionId uint   `json:"tradingSessionId"  binding:"required"`
	WorkspaceCode    string `json:"workspaceCode"     binding:"required"`
	Name             string `json:"name"              binding:"required"`
}

//=============================================================================

type DataProductSpec struct {
	ConnectionId uint    `json:"connectionId"   binding:"required"`
	ExchangeId   uint    `json:"exchangeId"     binding:"required"`
	Symbol       string  `json:"symbol"         binding:"required"`
	Name         string  `json:"name"           binding:"required"`
	MarketType   string  `json:"marketType"     binding:"required"`
	ProductType  string  `json:"productType"    binding:"required"`
}

//=============================================================================

type BrokerProductSpec struct {
	ConnectionId uint    `json:"connectionId"   binding:"required"`
	ExchangeId   uint    `json:"exchangeId"     binding:"required"`
	Symbol       string  `json:"symbol"         binding:"required"`
	Name         string  `json:"name"           binding:"required"`
	PointValue   float32 `json:"pointValue"     binding:"required,min=0"`
	CostPerTrade float32 `json:"costPerTrade"   binding:"required,min=0"`
	MarginValue  float32 `json:"marginValue"    binding:"required,min=0"`
	Increment    float64 `json:"increment"      binding:"required,min=0,max=1"`
	MarketType   string  `json:"marketType"     binding:"required"`
	ProductType  string  `json:"productType"    binding:"required"`
}

//=============================================================================

type TradingSession struct {
	db.Common
	Username  string                  `json:"username"`
	Name      string                  `json:"name"`
	Session   *session.TradingSession `json:"session"`
}

//=============================================================================
//===
//=== Portfolio tree
//===
//=============================================================================

type PortfolioTree struct {
	db.Portfolio
	Children       []*PortfolioTree        `json:"children"`
	TradingSystems []*db.TradingSystemFull `json:"tradingSystems"`
}

//-----------------------------------------------------------------------------

func (pt *PortfolioTree) AddChild(p *PortfolioTree) {
	pt.Children = append(pt.Children, p)
}

//-----------------------------------------------------------------------------

func (pt *PortfolioTree) AddTradingSystem(ts *db.TradingSystemFull) {
	pt.TradingSystems = append(pt.TradingSystems, ts)
}

//=============================================================================
//===
//=== ProductBroker & ProductData composite structs
//===
//=============================================================================

type BrokerProductExt struct {
	db.BrokerProduct
	Connection  db.Connection         `json:"connection"`
	Exchange    db.Exchange           `json:"exchange"`
	Instruments []db.BrokerInstrument `json:"instruments,omitempty"`
}

//=============================================================================

type DataProductExt struct {
	db.DataProduct
	Connection  db.Connection  `json:"connection,omitempty"`
	Exchange    db.Exchange    `json:"exchange,omitempty"`
}

//=============================================================================
//===
//=== Messages
//===
//=============================================================================

type TradingSystemMessage struct {
	TradingSystem db.TradingSystem `json:"tradingSystem"`
	BrokerProduct db.BrokerProduct `json:"brokerProduct"`
	Currency      db.Currency      `json:"currency"`
}

//=============================================================================

type DataProductMessage struct {
	DataProduct db.DataProduct `json:"dataProduct"`
	Connection  db.Connection  `json:"connection"`
	Exchange    db.Exchange    `json:"exchange"`
}

//=============================================================================

type BrokerProductMessage struct {
	BrokerProduct db.BrokerProduct `json:"brokerProduct"`
	Connection    db.Connection    `json:"connection"`
	Exchange      db.Exchange      `json:"exchange"`
	Currency      db.Currency      `json:"currency"`
}

//=============================================================================
