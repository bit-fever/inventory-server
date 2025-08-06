//=============================================================================
/*
Copyright © 2025 Andrea Carboni andrea.carboni71@gmail.com

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

package system

import (
	"encoding/json"
	"github.com/bit-fever/core/msg"
	"github.com/bit-fever/inventory-server/pkg/db"
	"gorm.io/gorm"
	"log/slog"
)

//=============================================================================

func InitMessageListener() {
	slog.Info("Starting system adapter message listener...")

	go msg.ReceiveMessages(msg.QuSystemToInventory, handleMessage)
}

//=============================================================================

func handleMessage(m *msg.Message) bool {

	slog.Info("handleMessage: New message received", "source", m.Source, "type", m.Type)

	if m.Source == msg.SourceSystem {
		if m.Type == msg.TypeRestart {
			return handleSystemAdapterRestart()
		}
	} else if m.Source == msg.SourceConnection {
		ccm := ConnectionChangeSystemMessage{}
		err := json.Unmarshal(m.Entity, &ccm)
		if err != nil {
			slog.Error("Dropping badly formatted message!", "entity", string(m.Entity))
			return true
		}

		if m.Type == msg.TypeChange {
			return handleConnectionChange(&ccm)
		}
	}

	slog.Error("handleMessage: Dropping message with unknown source/type!", "source", m.Source, "type", m.Type)
	return true
}

//=============================================================================

func handleSystemAdapterRestart() bool {
	slog.Info("handleSystemAdapterRestart: Unsetting connection status flag to all connections")

	err := db.RunInTransaction(func(tx *gorm.DB) error {
		return db.DisconnectAll(tx)
	})

	if err != nil {
		slog.Error("handleSystemAdapterRestart: Raised error while disconnecting", "error", err.Error())
	} else {
		slog.Info("handleSystemAdapterRestart: Disconnection complete")
	}

	return err == nil
}

//=============================================================================

func handleConnectionChange(ccm *ConnectionChangeSystemMessage) bool {
	if ccm.Status == ConnectionStatusConnecting {
		return true
	}

	slog.Info("handleConnectionChange: Updating connection status", "user", ccm.Username, "connectionCode", ccm.ConnectionCode, "status", ccm.Status)

	err := db.RunInTransaction(func(tx *gorm.DB) error {
		connected := ccm.Status == ConnectionStatusConnected
		return db.SetConnectionStatus(tx, ccm.Username, ccm.ConnectionCode, connected)
	})

	if err != nil {
		slog.Info("handleConnectionChange: Raise an error while changing connection status", "user", ccm.Username, "connectionCode", ccm.ConnectionCode, "status", ccm.Status, "error", err.Error())
	} else {
		slog.Info("handleConnectionChange: Connection status update complete", "user", ccm.Username, "connectionCode", ccm.ConnectionCode, "status", ccm.Status)
	}

	return err == nil
}

//=============================================================================
