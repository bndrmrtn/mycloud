package main

import (
	"encoding/json"

	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/middlewares"
	"github.com/bndrmrtn/my-cloud/utils"
	"gorm.io/gorm"
)

// NewWSServer creates a new WebSocket server
func NewWSServer(app *gale.Gale, db *gorm.DB) gale.WSServer {
	dispatcher := func(s gale.WSServer, msg gale.WSMessage) error {
		type EchoMsg struct {
			Type string `json:"type"`
		}

		var echoMsg EchoMsg
		if err := json.Unmarshal(msg.Content(), &echoMsg); err != nil {
			return err
		}

		if echoMsg.Type == "echo" {
			return msg.Conn().SendJSON(EchoMsg{
				Type: "echo",
			})
		}

		return nil
	}

	server := gale.NewWebSocketServer(dispatcher)

	// Register the WebSocket server endpoint
	app.WS("/ws", func(conn gale.WSConn) {
		user := conn.Ctx().Get(utils.RequestAuthUserKey)
		if user == nil {
			return
		}

		userID := user.(*models.User).ID
		conn.Ctx().Set(utils.WSUserID, userID)
		server.AddConn(conn)
	}, middlewares.AuthMiddleware(db)).Name("ws")

	return server
}
