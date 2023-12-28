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
	"github.com/bit-fever/core/auth"
	"github.com/bit-fever/core/msg"
	"github.com/bit-fever/inventory-server/pkg/db"
	"gorm.io/gorm"
)

//=============================================================================

func GetTradingSystems(tx *gorm.DB, c *auth.Context, filter map[string]any, offset int, limit int, details bool) (*[]db.TradingSystemFull, error) {
	if ! c.Session.IsAdmin() {
		filter["username"] = c.Session.Username
	}

	if details {
		return db.GetTradingSystemsFull(tx, filter, offset, limit)
	}

	return db.GetTradingSystems(tx, filter, offset, limit)
}

//=============================================================================

func AddTradingSystem(tx *gorm.DB, c *auth.Context, tss *TradingSystemSpec) (*db.TradingSystem, error) {
	c.Log.Info("AddTradingSystem: Adding a new trading system", "strategyCode", tss.StrategyCode, "name", tss.Name)

	var ts db.TradingSystem
	ts.Username         = c.Session.Username
	ts.StrategyCode     = tss.StrategyCode
	ts.Name             = tss.Name
	ts.PortfolioId      = tss.PortfolioId
	ts.ProductFeedId    = tss.ProductFeedId
	ts.ProductBrokerId  = tss.ProductBrokerId
	ts.TradingSessionId = tss.TradingSessionId

	err := db.AddTradingSystem(tx, &ts)

	if err != nil {
		c.Log.Error("AddTradingSystem: Could not add a new connection", "error", err.Error())
		return nil, err
	}

	err = msg.SendMessage(msg.ExInventoryUpdates, msg.OriginDb, msg.TypeCreate, msg.SourceTradingSystem, &ts)

	if err != nil {
		c.Log.Error("AddTradingSystem: Could not publish the update message", "error", err.Error())
		return nil, err
	}

	c.Log.Info("AddTradingSystem: Trading system added", "trategyCode", ts.StrategyCode, "id", ts.Id)
	return &ts, err
}

//=============================================================================
