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

import "github.com/bit-fever/inventory-server/pkg/db"

//=============================================================================

type ConnectionSpec struct {
	Code         string `json:"code"       binding:"required"`
	Name         string `json:"name"       binding:"required"`
	SystemCode   string `json:"systemCode" binding:"required"`
	SystemConfig string `json:"systemConfig"`
}

//=============================================================================

type TradingSystemSpec struct {
	Id               uint   `json:"id"`
	PortfolioId      uint   `json:"portfolioId"       binding:"required"`
	ProductFeedId    uint   `json:"productFeedId"     binding:"required"`
	ProductBrokerId  uint   `json:"productBrokerId"   binding:"required"`
	TradingSessionId uint   `json:"tradingSessionId"  binding:"required"`
	WorkspaceCode    string `json:"workspaceCode"     binding:"required"`
	Name             string `json:"name"              binding:"required"`
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
//=== ProductBroker & ProductFeed composite structs
//===
//=============================================================================

type ProductBrokerExt struct {
	db.ProductBroker
	Connection  db.Connection         `json:"connection"`
	Currency    db.Currency           `json:"currency"`
	Instruments []db.InstrumentBroker `json:"instruments,omitempty"`
}

//=============================================================================

type ProductFeedExt struct {
	db.ProductFeed
	Connection  db.Connection       `json:"connection"`
	Instruments []db.InstrumentFeed `json:"instruments,omitempty"`
}

//=============================================================================
//===
//=== Messages
//===
//=============================================================================

type TradingSystemMessage struct {
	TradingSystem db.TradingSystem `json:"tradingSystem"`
	ProductBroker db.ProductBroker `json:"productBroker"`
	Currency      db.Currency      `json:"currency"`
}

//=============================================================================
