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

func GetDataProducts(tx *gorm.DB, filter map[string]any, offset int, limit int) (*[]DataProductFull, error) {
	var list []DataProductFull
	res := tx.Where(filter).Offset(offset).Limit(limit).Find(&list)

	if res.Error != nil {
		return nil, req.NewServerErrorByError(res.Error)
	}

	return &list, nil
}

//=============================================================================

func GetDataProductsFull(tx *gorm.DB, filter map[string]any, offset int, limit int) (*[]DataProductFull, error) {
	var list []DataProductFull
	query :=	"SELECT dp.*, c.code as connection_code, c.system_code as system_code, e.code as exchange_code " +
				"FROM data_product dp " +
				"LEFT JOIN connection c on dp.connection_id = c.id "+
				"LEFT JOIN exchange   e on dp.exchange_id   = e.id"

	res := tx.Raw(query).Where(filter).Offset(offset).Limit(limit).Find(&list)

	if res.Error != nil {
		return nil, req.NewServerErrorByError(res.Error)
	}

	return &list, nil
}

//=============================================================================

func GetDataProductById(tx *gorm.DB, id uint) (*DataProduct, error) {
	var list []DataProduct
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

func AddDataProduct(tx *gorm.DB, ts *DataProduct) error {
	return tx.Create(ts).Error
}

//=============================================================================

func UpdateDataProduct(tx *gorm.DB, ts *DataProduct) error {
	return tx.Save(ts).Error
}

//=============================================================================
