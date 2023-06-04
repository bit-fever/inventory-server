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

import "time"

//=============================================================================

type Common struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
}

//=============================================================================

type Exchange struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

//=============================================================================

type DataSource struct {
	Id    int    `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Local bool   `json:"local"`
}

//=============================================================================

type Currency struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

//=============================================================================

type Instrument struct {
	Common
	ExchangeId     int       `json:"exchange-id"`
	DatasourceId   int       `json:"datasource-id"`
	Symbol         string    `json:"symbol"`
	Name           string    `json:"name"`
	ExpirationDate time.Time `json:"expiration-date,omitempty"`
	PriceScale     int       `json:"price-scale"`
	MinMovement    float32   `json:"min-movement"`
	BigPointValue  int       `json:"big-point-value"`
	CurrencyId     int       `json:"currency-id"`
	MarketType     string    `json:"market-type"`
	SecurityType   string    `json:"security-type"`
}

//=============================================================================
