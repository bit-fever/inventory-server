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

//=============================================================================

//func getInstruments(c *gin.Context, us *auth.UserSession) {
//
//	data, err := db.GetInstruments(nil, 0, 10000)
//
//	if err != nil {
//		c.IndentedJSON(http.StatusBadRequest, err.Error)
//	} else {
//		c.IndentedJSON(http.StatusOK, &data)
//	}
//}

//=============================================================================

//func getInstrumentById(c *gin.Context) {
//	id := c.Param("id")
//
//	data, err := db.GetInstrumentById(id)
//
//	if err != nil {
//		c.IndentedJSON(http.StatusBadRequest, gin.H{
//			"message": err.Error(),
//			"param": id,
//		})
//	} else {
//		c.IndentedJSON(http.StatusOK, &data)
//	}
//}

//=============================================================================

//func addInstrument(c *gin.Context) {
//	var instr db.Instrument
//	err := c.BindJSON(&instr)
//
//	if err != nil {
//		c.IndentedJSON(http.StatusBadRequest, err)
//	} else {
//		log.Printf("addInstrument: Symbol='%v', Name='%v'", instr.Symbol, instr.Name)
//		err = db.AddInstrument(&instr)
//
//		if err != nil {
//			c.IndentedJSON(http.StatusBadRequest, &instr)
//			log.Printf("addInstrument: Cannot add instrument --> %v", err)
//		} else {
//			c.IndentedJSON(http.StatusCreated, &instr)
//			log.Printf("addInstrument: Added with id=%v", instr.Id)
//		}
//	}
//}

//=============================================================================
