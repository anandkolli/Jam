package pg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PgClient struct {
	Pdb	*sql.DB
	dbname	string
	user	string
}

func NewPostgresClient(addr, uName, dbName string) (*PgClient, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable",
		addr, uName, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	client := new(PgClient)
	client.Pdb = db
	client.dbname = uName
	client.user = uName

	return client, nil
}

func (c *PgClient) Close() error {
	return c.Pdb.Close()
}
