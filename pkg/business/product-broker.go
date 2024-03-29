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

	//--- Get product broker

	pb, err := db.GetProductBrokerById(tx, id)
	if err != nil {
		return nil, err
	}
	if pb == nil {
		return nil, req.NewNotFoundError("Product broker not found: %v", id)
	}

	//--- Check access

	if ! c.Session.IsAdmin() {
		if pb.Username != c.Session.Username {
			return nil, req.NewForbiddenError("Product broker not owned by user: %v", id)
		}
	}

	//--- Get connection

	co, err := db.GetConnectionById(tx, pb.ConnectionId)
	if err != nil {
		return nil, err
	}

	//--- Get currency

	cu, err := db.GetCurrencyById(tx, pb.CurrencyId)
	if err != nil {
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
		Connection:    *co,
		Currency:      *cu,
		Instruments:   *instruments,
	}

	return &pbe, nil
}

//=============================================================================
