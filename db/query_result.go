package db

import (
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

type QueryResult struct {
	rows pgx.Rows
}

func (qr *QueryResult) ScanAll(dst interface{}) error {
	return pgxscan.ScanAll(dst, qr.rows)
}

func (qr *QueryResult) ScanFirst(dst interface{}) error {
	return pgxscan.ScanOne(dst, qr.rows)
}
