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
	"github.com/bit-fever/inventory-server/pkg/model/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

//=============================================================================

func GetExchanges(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, []db.Exchange{
		{Code: "A", Name: "Apple"},
		{Code: "G", Name: "Google"},
	})
}

//=============================================================================

func AddExchange(c *gin.Context) {
	var newInstr db.Instrument

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newInstr); err != nil {
		return
	}

	// Add the new album to the slice.
	//	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newInstr)
}

//=============================================================================

func GetExchangeById(c *gin.Context) {
	//	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	//for _, a := range albums {
	//	if a.ID == id {
	//		c.IndentedJSON(http.StatusOK, a)
	//		return
	//	}
	//}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

//=============================================================================
