package main


import (
//	"flag"
	"fmt"
	"os"
	"jam/pkg/pg"
	"jam/pkg/server"
)

const (
  user     = "postgres"
  dbname   = "anything"
)

var (
	db	*pg.PgClient
)

func init(){
	var err error
	var pgAddr string
	//posthgressAddr := flag.String("postgressAddr", "localhost", "IP address of Postgress")
	//flag.Parse()
	if pgAddr = os.Getenv("POSTGRESSADDR"); pgAddr == "" {
		fmt.Println("set post gress address env variable POSTGRESSADDR")
		os.Exit(1)
	}
	db, err = pg.NewPostgresClient(pgAddr, user, dbname)
	if err != nil {
		fmt.Println("error connecting db",err)
		os.Exit(1)
	}
}
func main() {
	server := server.NewServer(":9090", db)
        server.Run()
	//db.GetTimeSpentPerPerson()
	//db.GetAllActivity()
}
