package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"lab5/model"
	"log"
	"net/http"
	"strconv"
)

func processRequest(conn *websocket.Conn, msg []byte) {
	var r model.Request
	err := json.Unmarshal(msg, &r)
	if err != nil {
		err = conn.WriteJSON(model.ErrResponse{Err: "bad request"})
		if err != nil {
			log.Println(err)
		}
		return
	}
	log.Printf("%+v", r)

	bin := strconv.FormatInt(int64(r.Value), 2)
	hex := strconv.FormatInt(int64(r.Value), 16)

	if err := conn.WriteJSON(model.Response{Bin: bin, Hex: hex}); err != nil {
		log.Println(err)
		return
	}
}

func listenSocket(conn *websocket.Conn) {
	for {
		messageType, b, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		} else if messageType != websocket.TextMessage {
			err := conn.WriteJSON(model.ErrResponse{Err: "incorrect message type"})
			if err != nil {
				log.Println(err)
			}
		}
		go processRequest(conn, b)
	}
}

func handleSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// upgrade this connection to a WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	defer ws.Close()

	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Welcome!"))
	if err != nil {
		log.Println(err)
	}

	listenSocket(ws)
}
