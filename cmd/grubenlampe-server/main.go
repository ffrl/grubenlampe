package main

import (
	"flag"
	"log"
	"net"
	"os"

	// pb "github.com/ffrl/grubenlampe/api"
	// "github.com/ffrl/grubenlampe/database"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	log.SetFlags(0)
}

func main() {
	var (
		verbose       = flag.Bool("verbose", false, "Enable verbose logging")
		listenAddress = flag.String("listen", "[::1]:20170", "GRPC listener host:port")
		driver        = flag.String("db", "sqlite3", "Database driver (sqlite3, mysql or postgres)")
		dsn           = flag.String("dsn", "grubenlampe.db", "Database DSN (GRUBENLAMPE_DSN)")
	)
	flag.Parse()

	if envDSN := os.Getenv("GRUBENLAMPE_DSN"); envDSN != "" {
		*dsn = envDSN
	}

	db, err := gorm.Open(*driver, *dsn)
	if err != nil {
		log.Fatal(err)
	}
	if *verbose {
		db.LogMode(true)
	}
	db.AutoMigrate()

	lis, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
