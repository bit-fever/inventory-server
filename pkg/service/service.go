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
	"github.com/bit-fever/core/auth/roles"
	"github.com/bit-fever/core/req"
	"github.com/bit-fever/inventory-server/pkg/app"
	"github.com/gin-gonic/gin"
	"log/slog"
)

//=============================================================================

func Init(router *gin.Engine, cfg *app.Config, logger *slog.Logger) {

	ctrl := auth.NewOidcController(cfg.Authentication.Authority, req.GetClient("bf"), logger, cfg)

	//--- Inventory

	router.GET ("/api/inventory/v1/currencies",                   ctrl.Secure(getCurrencies,                roles.Admin_User_Service))
	router.GET ("/api/inventory/v1/exchanges",                    ctrl.Secure(getExchanges,                 roles.Admin_User_Service))

	router.GET ("/api/inventory/v1/product-data",                 ctrl.Secure(getProductData,               roles.Admin_User_Service))
	router.POST("/api/inventory/v1/product-data",                 ctrl.Secure(addProductData,               roles.Admin_User_Service))
	router.GET ("/api/inventory/v1/product-data/:id",             ctrl.Secure(getProductDataById,           roles.Admin_User_Service))
	router.PUT ("/api/inventory/v1/product-data/:id",             ctrl.Secure(updateProductData,            roles.Admin_User_Service))

	router.GET ("/api/inventory/v1/product-brokers",              ctrl.Secure(getProductBrokers,            roles.Admin_User_Service))
	router.POST("/api/inventory/v1/product-brokers",              ctrl.Secure(addProductBroker,             roles.Admin_User_Service))
	router.GET ("/api/inventory/v1/product-brokers/:id",          ctrl.Secure(getProductBrokerById,         roles.Admin_User_Service))
	router.PUT ("/api/inventory/v1/product-brokers/:id",          ctrl.Secure(updateProductBroker,          roles.Admin_User_Service))

	router.GET ("/api/inventory/v1/trading-systems",              ctrl.Secure(getTradingSystems,            roles.Admin_User_Service))
	router.POST("/api/inventory/v1/trading-systems",              ctrl.Secure(addTradingSystem,             roles.Admin_User_Service))
	router.PUT ("/api/inventory/v1/trading-systems/:id",          ctrl.Secure(updateTradingSystem,          roles.Admin_User_Service))

	//--- Portfolio

	router.GET ("/api/inventory/v1/portfolios",          ctrl.Secure(getPortfolios,       roles.Admin_User_Service))

	router.GET ("/api/inventory/v1/portfolio/tree",      ctrl.Secure(getPortfolioTree,    roles.Admin_User_Service))

	router.GET ("/api/inventory/v1/trading-sessions",    ctrl.Secure(getTradingSessions,  roles.Admin_User_Service))

	//--- Administration

	router.GET ("/api/inventory/v1/connections",         ctrl.Secure(getConnections,      roles.Admin_User_Service))
	router.GET ("/api/inventory/v1/connections/:id",     ctrl.Secure(getConnectionById,   roles.Admin_User_Service))
	router.POST("/api/inventory/v1/connections",         ctrl.Secure(addConnection,       roles.Admin_User_Service))
}

//=============================================================================
