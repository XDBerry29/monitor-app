package models

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	ID              string
	ConnWs          *websocket.Conn
	LogWs           *websocket.Conn
	min_severity    int
	monitoring_list map[string]bool
}

func NewClient(id string) *Client {
	return &Client{
		ID:              id,
		min_severity:    DEBUG,
		monitoring_list: make(map[string]bool),
	}
}

func (c *Client) GetID() string {
	return c.ID
}

func (c *Client) GetConnWs() *websocket.Conn {
	return c.ConnWs
}

func (c *Client) GetLogWs() *websocket.Conn {
	return c.LogWs
}

func (c *Client) GetMinSeverity() int {
	return c.min_severity
}

func (c *Client) AmIMonitoring(id string) bool {
	return c.monitoring_list[id]
}

func (c *Client) SetMonitoring(id string, state bool) {
	c.monitoring_list[id] = state
}

func (c *Client) SetMinSeverity(sev int) {
	c.min_severity = sev
}

func (c *Client) SetID(id string) {
	c.ID = id
}

func (c *Client) SetConnWs(connWs *websocket.Conn) {
	c.ConnWs = connWs
}

func (c *Client) SetLogWs(logWs *websocket.Conn) {
	c.LogWs = logWs
}
