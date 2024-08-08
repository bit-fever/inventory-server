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
	"github.com/bit-fever/core/req"
	"gorm.io/gorm"
)

//=============================================================================

func GetTradingSystems(tx *gorm.DB, filter map[string]any, offset int, limit int) (*[]TradingSystemFull, error) {
	var list []TradingSystemFull
	res := tx.Where(filter).Offset(offset).Limit(limit).Find(&list)

	if res.Error != nil {
		return nil, req.NewServerErrorByError(res.Error)
	}

	return &list, nil
}

//=============================================================================

func GetTradingSystemById(tx *gorm.DB, id uint) (*TradingSystem, error) {
	var list []TradingSystem
	res := tx.Find(&list, id)

	if res.Error != nil {
		return nil, req.NewServerErrorByError(res.Error)
	}

	if len(list) == 1 {
		return &list[0], nil
	}

	return nil, nil
}

//=============================================================================

//func GetTradingSystemsByIds(tx *gorm.DB, ids []uint) (*[]TradingSystem, error) {
//	var list []TradingSystem
//	res := tx.Find(&list, ids)
//
//	if res.Error != nil {
//		return nil, req.NewServerErrorByError(res.Error)
//	}
//
//	return &list, nil
//}

//=============================================================================

func GetTradingSystemsFull(tx *gorm.DB, filter map[string]any, offset int, limit int) (*[]TradingSystemFull, error) {
	var list []TradingSystemFull
	query := "SELECT ts.*, pd.symbol as data_symbol, pb.symbol as broker_symbol, p.name as portfolio_name, s.name as trading_session " +
		"FROM trading_system ts " +
		"LEFT JOIN product_data    pd on ts.product_data_id   = pd.id " +
		"LEFT JOIN product_broker  pb on ts.product_broker_id = pb.id " +
		"LEFT JOIN portfolio       p  on ts.portfolio_id      = p.id " +
		"LEFT JOIN trading_session s  on ts.trading_session_id= s.id"

	res := tx.Raw(query).Where(filter).Offset(offset).Limit(limit).Find(&list)

	if res.Error != nil {
		return nil, req.NewServerErrorByError(res.Error)
	}

	return &list, nil
}

//=============================================================================

func AddTradingSystem(tx *gorm.DB, ts *TradingSystem) error {
	return tx.Create(ts).Error
}

//=============================================================================

func UpdateTradingSystem(tx *gorm.DB, ts *TradingSystem) error {
	return tx.Save(ts).Error
}

//=============================================================================
