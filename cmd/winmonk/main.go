package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/nico0302/winmonk/internal/sync"
	"github.com/nico0302/winmonk/pkg/listmonk"
	"github.com/nico0302/winmonk/pkg/winestro"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DB_CONN"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	uid, err := strconv.Atoi(os.Getenv("WBO_UID"))
	if err != nil {
		log.Fatal(err)
	}
	shopID, err := strconv.Atoi(os.Getenv("WBO_SHOP_ID"))
	if err != nil {
		log.Fatal(err)
	}

	winestro := winestro.NewClient(uid, os.Getenv("WBO_API_USER"), os.Getenv("WBO_API_CODE"), shopID)
	queries := listmonk.New(conn)

	groupID, err := strconv.Atoi(os.Getenv("WBO_GROUP_ID"))
	if err != nil {
		log.Fatal(err)
	}
	listID, err := strconv.Atoi(os.Getenv("LISTMONK_LIST_ID"))
	if err != nil {
		log.Fatal(err)
	}

	client := sync.NewClient(winestro, queries)
	err = client.ImportContacts(ctx, groupID, listID)
	if err != nil {
		log.Fatal(err)
	}

}
