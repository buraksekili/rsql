package client

import (
	"database/sql"

	"github.com/buraksekili/rsql/data"

	"github.com/buraksekili/selog"
)

type DbClient struct {
	ConnInfo *data.ConnInfo
	Log      *selog.Selog
	db       *sql.DB
}

func NewDbClient(l *selog.Selog) *DbClient {
	return &DbClient{&data.ConnInfo{}, l, nil}
}
