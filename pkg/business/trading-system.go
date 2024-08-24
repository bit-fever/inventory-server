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
	"github.com/bit-fever/core/req"
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
	c.Log.Info("AddTradingSystem: Adding a new trading system", "workspaceCode", tss.WorkspaceCode, "name", tss.Name)

	var ts db.TradingSystem
	ts.Username         = c.Session.Username
	ts.WorkspaceCode    = tss.WorkspaceCode
	ts.Name             = tss.Name
	ts.PortfolioId      = tss.PortfolioId
	ts.DataProductId    = tss.DataProductId
	ts.BrokerProductId  = tss.BrokerProductId
	ts.TradingSessionId = tss.TradingSessionId

	err := db.AddTradingSystem(tx, &ts)

	if err != nil {
		c.Log.Error("AddTradingSystem: Could not add a new trading system", "error", err.Error())
		return nil, err
	}

	err = sendChangeMessage(tx, c, &ts, msg.TypeCreate)
	if err != nil {
		return nil, err
	}

	c.Log.Info("AddTradingSystem: Trading system added", "workspaceCode", ts.WorkspaceCode, "id", ts.Id)
	return &ts, err
}

//=============================================================================

func UpdateTradingSystem(tx *gorm.DB, c *auth.Context, id uint, tss *TradingSystemSpec) (*db.TradingSystem, error) {
	c.Log.Info("UpdateTradingSystem: Updating a trading system", "id", id, "name", tss.Name)

	ts, err := db.GetTradingSystemById(tx, id)
	if err != nil {
		c.Log.Error("UpdateTradingSystem: Could not retrieve trading system", "error", err.Error())
		return nil, err
	}
	if ts == nil {
		c.Log.Error("UpdateTradingSystem: Trading system was not found", "id", id)
		return nil, req.NewNotFoundError("Trading system was not found: %v", id)
	}

	if ts.Username != c.Session.Username {
		c.Log.Error("UpdateTradingSystem: Trading system not owned by user", "id", id)
		return nil, req.NewForbiddenError("Trading system is not owned by user: %v", id)
	}

	ts.WorkspaceCode    = tss.WorkspaceCode
	ts.Name             = tss.Name
	ts.PortfolioId      = tss.PortfolioId
	ts.DataProductId    = tss.DataProductId
	ts.BrokerProductId  = tss.BrokerProductId
	ts.TradingSessionId = tss.TradingSessionId

	db.UpdateTradingSystem(tx, ts)

	err = sendChangeMessage(tx, c, ts, msg.TypeUpdate)
	if err != nil {
		return nil, err
	}

	c.Log.Info("UpdateTradingSystem: Trading system updated", "id", ts.Id, "name", ts.Name)
	return ts, err
}

//=============================================================================
//===
//=== Private functions
//===
//=============================================================================

func sendChangeMessage(tx *gorm.DB, c *auth.Context, ts *db.TradingSystem, msgType int) error {
	bp, err := db.GetBrokerProductById(tx, ts.BrokerProductId)
	if err != nil {
		c.Log.Error("[Add|Update]TradingSystem: Could not retrieve broker product", "error", err.Error())
		return err
	}

	ex, err := db.GetExchangeById(tx, bp.ExchangeId)
	if err != nil {
		c.Log.Error("[Add|Update]TradingSystem: Could not retrieve exchange", "error", err.Error())
		return err
	}

	cu, err := db.GetCurrencyById(tx, ex.CurrencyId)
	if err != nil {
		c.Log.Error("[Add|Update]TradingSystem: Could not retrieve currency", "error", err.Error())
		return err
	}

	tsm := TradingSystemMessage{*ts, *bp, *cu}
	err = msg.SendMessage(msg.ExInventoryUpdates, msg.OriginDb, msgType, msg.SourceTradingSystem, &tsm)

	if err != nil {
		c.Log.Error("[Add|Update]TradingSystem: Could not publish the update message", "error", err.Error())
		return err
	}

	return nil
}

//=============================================================================
