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

package service

import (
	"github.com/bit-fever/core/auth"
	"github.com/bit-fever/inventory-server/pkg/business"
	"github.com/bit-fever/inventory-server/pkg/db"
	"gorm.io/gorm"
)

//=============================================================================

func getTradingSystems(c *auth.Context) {
	filter := map[string]any{}
	offset, limit, err := c.GetPagingParams()

	if err == nil {
		var details bool
		details, err = c.GetParamAsBool("details", false)

		if err == nil {
			err = db.RunInTransaction(func(tx *gorm.DB) error {
				list, err := business.GetTradingSystems(tx, c, filter, offset, limit, details)

				if err != nil {
					return err
				}

				return c.ReturnList(list, offset, limit, len(*list))
			})
		}
	}

	c.ReturnError(err)
}

//=============================================================================

func addTradingSystem(c *auth.Context) {
	var tss business.TradingSystemSpec
	err := c.BindParamsFromBody(&tss)

	if err == nil {
		err = db.RunInTransaction(func(tx *gorm.DB) error {
			ts, err := business.AddTradingSystem(tx, c, &tss)

			if err != nil {
				return err
			}

			return c.ReturnObject(ts)
		})
	}

	c.ReturnError(err)
}

//=============================================================================

func updateTradingSystem(c *auth.Context) {
	var tss business.TradingSystemSpec
	err := c.BindParamsFromBody(&tss)

	if err == nil {
		var id uint
		id,err = c.GetIdFromUrl()

		if err == nil {
			err = db.RunInTransaction(func(tx *gorm.DB) error {
				var ts *db.TradingSystem
				ts, err = business.UpdateTradingSystem(tx, c, id, &tss)

				if err != nil {
					return err
				}

				return c.ReturnObject(ts)
			})
		}
	}

	c.ReturnError(err)
}

//=============================================================================

func deleteTradingSystem(c *auth.Context) {
	id,err := c.GetIdFromUrl()

	if err == nil {
		err = db.RunInTransaction(func(tx *gorm.DB) error {
			ts,err := business.DeleteTradingSystem(tx, c, id)

			if err != nil {
				return err
			}

			return c.ReturnObject(ts)
		})
	}

	c.ReturnError(err)
}

//=============================================================================

func finalizeTradingSystem(c *auth.Context) {
	id,err := c.GetIdFromUrl()

	if err == nil {
		err = db.RunInTransaction(func(tx *gorm.DB) error {
			ts, err := business.FinalizeTradingSystem(tx, c, id)

			if err != nil {
				return err
			}

			return c.ReturnObject(ts)
		})
	}
	c.ReturnError(err)
}

//=============================================================================
