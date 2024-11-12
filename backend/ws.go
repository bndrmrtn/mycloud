package main

import (
	"context"
	"encoding/json"

	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/middlewares"
	"github.com/bndrmrtn/my-cloud/utils"
	"gorm.io/gorm"
)

func NewWSServer(app *gale.Gale, db *gorm.DB) gale.WSServer {
	server := gale.NewWSServer(context.Background())

	app.WS("/ws", func(conn gale.WSConn) {
		user := conn.Ctx().Get(utils.RequestAuthUserKey)
		if user == nil {
			return
		}

		userID := user.(*models.User).ID
		conn.Ctx().Set(utils.WSUserID, userID)
		server.AddConn(conn)
	}, middlewares.AuthMiddleware(db)).Name("ws")

	server.OnMessage(func(s gale.WSServer, conn gale.WSConn, msg []byte) error {
		type EchoMsg struct {
			Type string `json:"type"`
		}

		var echoMsg EchoMsg
		if err := json.Unmarshal(msg, &echoMsg); err != nil {
			return err
		}

		if echoMsg.Type == "echo" {
			return conn.SendJSON(EchoMsg{
				Type: "echo",
			})
		}

		return nil
	})

	return server
}
