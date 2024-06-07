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

func GetProductBrokers(tx *gorm.DB, c *auth.Context, filter map[string]any, offset int, limit int, details bool) (*[]db.ProductBrokerFull, error) {
	if ! c.Session.IsAdmin() {
		filter["username"] = c.Session.Username
	}

	if details {
		return db.GetProductBrokersFull(tx, filter, offset, limit)
	}

	return db.GetProductBrokers(tx, filter, offset, limit)
}

//=============================================================================

func GetProductBrokerById(tx *gorm.DB, c *auth.Context, id uint, details bool) (*ProductBrokerExt, error) {
	c.Log.Info("GetProductBrokerById: Getting a product for broker", "id", id)

	pb, err := getProductBrokerAndCheckAccess(tx, c, id, "GetProductBrokerById")
	if err != nil {
		return nil, err
	}

	//--- Get connection

	conn, err := db.GetConnectionById(tx, pb.ConnectionId)
	if err != nil {
		c.Log.Error("GetProductBrokerById: Could not retrieve connection", "error", err.Error())
		return nil, err
	}

	//--- Get exchange

	ex, err := db.GetExchangeById(tx, pb.ExchangeId)
	if err != nil {
		c.Log.Error("GetProductBrokerById: Could not retrieve exchange", "error", err.Error())
		return nil, err
	}

	//--- Add instruments, if it is the case

	var instruments *[]db.InstrumentBroker

	if details {
		instruments, err = db.GetInstrumentsByBrokerId(tx, pb.Id)
	}

	//--- Put all together

	pbe := ProductBrokerExt{
		ProductBroker: *pb,
		Connection:    *conn,
		Exchange:      *ex,
		Instruments:   *instruments,
	}

	return &pbe, nil
}

//=============================================================================

func AddProductBroker(tx *gorm.DB, c *auth.Context, pbs *ProductBrokerSpec) (*db.ProductBroker, error) {
	c.Log.Info("AddProductBroker: Adding a new product for broker", "symbol", pbs.Symbol, "name", pbs.Name)

	var pb db.ProductBroker
	pb.ConnectionId = pbs.ConnectionId
	pb.ExchangeId   = pbs.ExchangeId
	pb.Username     = c.Session.Username
	pb.Symbol       = pbs.Symbol
	pb.Name         = pbs.Name
	pb.PointValue   = pbs.PointValue
	pb.CostPerTrade = pbs.CostPerTrade
	pb.MarginValue  = pbs.MarginValue
	pb.MarketType   = pbs.MarketType
	pb.ProductType  = pbs.ProductType

	err := db.AddProductBroker(tx, &pb)

	if err != nil {
		c.Log.Error("AddProductBroker: Could not add a new product for broker", "error", err.Error())
		return nil, err
	}

	err = sendProductBrokerChangeMessage(tx, c, &pb, msg.TypeUpdate)
	if err != nil {
		return nil, err
	}

	c.Log.Info("AddProductBroker: Product for broker added", "symbol", pb.Symbol, "id", pb.Id)
	return &pb, err
}

//=============================================================================

func UpdateProductBroker(tx *gorm.DB, c *auth.Context, id uint, pbs *ProductBrokerSpec) (*db.ProductBroker, error) {
	c.Log.Info("UpdateProductBroker: Updating a product for broker", "id", id, "name", pbs.Name)

	pb, err := getProductBrokerAndCheckAccess(tx, c, id, "UpdateProductBroker")
	if err != nil {
		return nil, err
	}

	pb.ExchangeId  = pbs.ExchangeId
	pb.Symbol      = pbs.Symbol
	pb.Name        = pbs.Name
	pb.PointValue  = pbs.PointValue
	pb.CostPerTrade= pbs.CostPerTrade
	pb.MarginValue = pbs.MarginValue
	pb.MarketType  = pbs.MarketType
	pb.ProductType = pbs.ProductType

	db.UpdateProductBroker(tx, pb)

	err = sendProductBrokerChangeMessage(tx, c, pb, msg.TypeUpdate)
	if err != nil {
		return nil, err
	}

	c.Log.Info("UpdateProductBroker: Product for broker updated", "id", pb.Id, "name", pb.Name)
	return pb, err
}

//=============================================================================
//===
//=== Private functions
//===
//=============================================================================

func getProductBrokerAndCheckAccess(tx *gorm.DB, c *auth.Context, id uint, function string) (*db.ProductBroker, error) {
	pb, err := db.GetProductBrokerById(tx, id)

	if err != nil {
		c.Log.Error(function +": Could not retrieve product for broker", "error", err.Error())
		return nil, err
	}

	if pb == nil {
		c.Log.Error(function +": Product for broker was not found", "id", id)
		return nil, req.NewNotFoundError("Product for broker was not found: %v", id)
	}

	if ! c.Session.IsAdmin() {
		if pb.Username != c.Session.Username {
			c.Log.Error(function+": Product for broker not owned by user", "id", id)
			return nil, req.NewForbiddenError("Product for broker is not owned by user: %v", id)
		}
	}

	return pb, nil
}

//=============================================================================

func sendProductBrokerChangeMessage(tx *gorm.DB, c *auth.Context, pb *db.ProductBroker, msgType int) error {
	conn, err := db.GetConnectionById(tx, pb.ConnectionId)
	if err != nil {
		c.Log.Error("[Add|Update]ProductBroker: Could not retrieve connection", "error", err.Error())
		return err
	}

	exc, err := db.GetExchangeById(tx, pb.ExchangeId)
	if err != nil {
		c.Log.Error("[Add|Update]ProductBroker: Could not retrieve exchange", "error", err.Error())
		return err
	}

	pbm := ProductBrokerMessage{*pb, *conn, *exc}
	err = msg.SendMessage(msg.ExInventoryUpdates, msg.OriginDb, msgType, msg.SourceProductBroker, &pbm)

	if err != nil {
		c.Log.Error("[Add|Update]ProductBroker: Could not publish the update message", "error", err.Error())
		return err
	}

	return nil
}

//=============================================================================
