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
	"time"
)

//=============================================================================

type Common struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

//=============================================================================

type Exchange struct {
	Common
	Code     string `json:"code"`
	Name     string `json:"name"`
	Timezone string `json:"timezone"`
	Offset   int    `json:"offset"`
}

//=============================================================================

type Currency struct {
	Id   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

//=============================================================================

type Connection struct {
	Common
	Username              string `json:"username"`
	Code                  string `json:"code"`
	Name                  string `json:"name"`
	SystemCode            string `json:"systemCode"`
	SystemName            string `json:"systemName"`
	SystemConfig          string `json:"systemConfig"`
	ConnectionCode        string `json:"connectionCode"`
	SupportsFeed          bool   `json:"supportsFeed"`
	SupportsBroker        bool   `json:"supportsBroker"`
	SupportsMultipleFeeds bool   `json:"supportsMultipleFeeds"`
	SupportsInventory     bool   `json:"supportsInventory"`
}

//=============================================================================

type Product struct {
	Common
	ExchangeId  uint   `json:"exchangeId"`
	CurrencyId  uint   `json:"currencyId"`
	SessionId   uint   `json:"sessionId"`
	Username    string `json:"username"`
	Symbol      string `json:"symbol"`
	MarketType  string `json:"marketType"`
	ProductType string `json:"productType"`
}

//=============================================================================

type ProductFeed struct {
	Common
	ProductId   uint     `json:"productId"`
	FeedId      uint     `json:"feedId"`
	Symbol      string   `json:"symbol"`
	Name        string   `json:"name"`
	PriceScale  int      `json:"priceScale"`
	MinMovement float32  `json:"minMovement"`
	PointValue  float32  `json:"pointValue"`
}

//=============================================================================

type ProductBroker struct {
	Common
	ProductId        uint       `json:"productId"`
	BrokerId         uint       `json:"brokerId"`
	Symbol           string     `json:"symbol"`
	Name             string     `json:"name"`
	CostPerTrade     float32    `json:"costPerTrade"`
	MarginValue      float32    `json:"marginValue"`
	MarginLastUpdate time.Time  `json:"marginLastUpdate"`
	MarginAutoSync   int        `json:"marginAutoSync"`
}

//=============================================================================

type ProductBrokerFull struct {
	ProductBroker
	CurrencyCode     string     `json:"currencyCode"`
	ConnectionCode   string     `json:"connectionCode"`
	ProductSymbol    string     `json:"productSymbol"`
}

//=============================================================================

type Instrument struct {
	Common
	ProductSourceId  uint       `json:"productSourceId"`
	Symbol           string     `json:"symbol"`
	Name             string     `json:"name"`
	ExpirationDate   time.Time  `json:"expirationDate"`
}

//=============================================================================

type ProductSession struct {
	Common
	Username  string `json:"username"`
	Name      string `json:"name"`
	Config    string `json:"config"`
}

//=============================================================================
//===
//=== Table names
//===
//=============================================================================

func (Exchange) TableName() string {
	return "exchange"
}

//=============================================================================

func (Currency) TableName() string {
	return "currency"
}

//=============================================================================

func (Connection) TableName() string {
	return "connection"
}

//=============================================================================

func (Product) TableName() string {
	return "product"
}

//=============================================================================

func (ProductFeed) TableName() string {
	return "product_feed"
}

//=============================================================================

func (ProductBroker) TableName() string {
	return "product_broker"
}

//=============================================================================

func (ProductSession) TableName() string {
	return "product_session"
}

//=============================================================================
