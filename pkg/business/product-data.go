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
	"github.com/bit-fever/core/req"
	"github.com/bit-fever/inventory-server/pkg/db"
	"gorm.io/gorm"
)

//=============================================================================

func GetProductData(tx *gorm.DB, c *auth.Context, filter map[string]any, offset int, limit int, details bool) (*[]db.ProductDataFull, error) {
	if ! c.Session.IsAdmin() {
		filter["username"] = c.Session.Username
	}

	if details {
		return db.GetProductDataFull(tx, filter, offset, limit)
	}

	return db.GetProductData(tx, filter, offset, limit)
}

//=============================================================================

func AddProductData(tx *gorm.DB, c *auth.Context, pds *ProductDataSpec) (*db.ProductData, error) {
	c.Log.Info("AddProductData: Adding a new product for data", "symbol", pds.Symbol, "name", pds.Name)

	var pd db.ProductData
	pd.ConnectionId = pds.ConnectionId
	pd.ExchangeId   = pds.ExchangeId
	pd.Username     = c.Session.Username
	pd.Symbol       = pds.Symbol
	pd.Name         = pds.Name
	pd.Increment    = pds.Increment
	pd.MarketType   = pds.MarketType
	pd.ProductType  = pds.ProductType
	pd.LocalClass   = pds.LocalClass

	err := db.AddProductData(tx, &pd)

	if err != nil {
		c.Log.Error("AddProductData: Could not add a new product for data", "error", err.Error())
		return nil, err
	}

	//err = sendChangeMessage(tx, c, &pd, msg.TypeCreate)
	//if err != nil {
	//	return nil, err
	//}

	c.Log.Info("AddProductData: Product for data added", "symbol", pd.Symbol, "id", pd.Id)
	return &pd, err
}

//=============================================================================

func UpdateProductData(tx *gorm.DB, c *auth.Context, id uint, pds *ProductDataSpec) (*db.ProductData, error) {
	c.Log.Info("UpdateProductData: Updating a product for data", "id", id, "name", pds.Name)

	pd, err := db.GetProductDataById(tx, id)
	if err != nil {
		c.Log.Error("UpdateProductData: Could not retrieve product for data", "error", err.Error())
		return nil, err
	}
	if pd == nil {
		c.Log.Error("UpdateProductData: Product for data was not found", "id", id)
		return nil, req.NewNotFoundError("Product for data was not found: %v", id)
	}

	if pd.Username != c.Session.Username {
		c.Log.Error("UpdateProductData: Product for data not owned by user", "id", id)
		return nil, req.NewForbiddenError("Product for data is not owned by user: %v", id)
	}

	pd.ExchangeId  = pds.ExchangeId
	pd.Symbol      = pds.Symbol
	pd.Name        = pds.Name
	pd.Increment   = pds.Increment
	pd.MarketType  = pds.MarketType
	pd.ProductType = pds.ProductType
	pd.LocalClass  = pds.LocalClass

	db.UpdateProductData(tx, pd)

	//err = sendChangeMessage(tx, c, ts, msg.TypeUpdate)
	//if err != nil {
	//	return nil, err
	//}

	c.Log.Info("UpdateProductData: Product for data updated", "id", pd.Id, "name", pd.Name)
	return pd, err
}

//=============================================================================
//===
//=== Private functions
//===
//=============================================================================

//func sendChangeMessageX(tx *gorm.DB, c *auth.Context, ts *db.TradingSystem, msgType int) error {
//	pb, err := db.GetProductBrokerById(tx, ts.ProductBrokerId)
//	if err != nil {
//		c.Log.Error("[Add|Update]TradingSystem: Could not retrieve product broker", "error", err.Error())
//		return err
//	}
//
//	cu, err := db.GetCurrencyById(tx, pb.CurrencyId)
//	if err != nil {
//		c.Log.Error("[Add|Update]TradingSystem: Could not retrieve currency", "error", err.Error())
//		return err
//	}
//
//	tsm := TradingSystemMessage{*ts, *pb, *cu}
//	err = msg.SendMessage(msg.ExInventoryUpdates, msg.OriginDb, msgType, msg.SourceTradingSystem, &tsm)
//
//	if err != nil {
//		c.Log.Error("[Add|Update]TradingSystem: Could not publish the update message", "error", err.Error())
//		return err
//	}
//
//	return nil
//}

//=============================================================================
