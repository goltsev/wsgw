package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WSServer struct {
	server   *gin.Engine
	service  *Service
	upgrader *websocket.Upgrader
}

func NewWSServer(service *Service) *WSServer {
	server := gin.Default()
	wsserver := &WSServer{
		server:   server,
		service:  service,
		upgrader: &websocket.Upgrader{},
	}
	server.GET("/", wsserver.HandleHTTP)
	return wsserver
}

func (s *WSServer) HandleHTTP(c *gin.Context) {
	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	defer conn.Close()

	if err := s.HandleWebSocket(c.Request.Context(), conn); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

func (s *WSServer) HandleWebSocket(ctx context.Context, conn *websocket.Conn) error {
	var r Request
	sub := NewSubscriber(ctx, conn)
	defer s.service.Unsubscribe(sub)
	for {
		if err := conn.ReadJSON(&r); err != nil {
			return err
		}
		switch r.Action {
		case "subscribe":
			sub.List = r.Symbols
			s.service.Subscribe(sub)
		case "unsubscribe":
			s.service.Unsubscribe(sub)
		default:
		}
	}
}

func (s *WSServer) Serve() error {
	return s.server.Run()
}
