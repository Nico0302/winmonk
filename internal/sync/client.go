package sync

import (
	"github.com/nico0302/go-winestro"
	"github.com/nico0302/winmonk/pkg/listmonk"
)

type Client struct {
	winestro *winestro.Client
	listmonk *listmonk.Queries
}

func NewClient(winestro *winestro.Client, listmonk *listmonk.Queries) *Client {
	return &Client{
		winestro: winestro,
		listmonk: listmonk,
	}
}
