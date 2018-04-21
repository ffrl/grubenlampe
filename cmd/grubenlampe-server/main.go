package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	// pb "github.com/ffrl/grubenlampe/api"

	"github.com/ffrl/grubenlampe/database"
	"github.com/ffrl/grubenlampe/server"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

	db, err := initDatabase(*driver, *dsn, *verbose)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		log.Fatal(err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGTERM, syscall.SIGINT)

	s := server.New()

	go func() {
		log.Println("Starting GRPC server on", *listenAddress)
		err := s.Serve(lis)
		if err != nil {
			log.Fatal(err)
		}
	}()

	select {
	case sig := <-sigchan:
		log.Printf("Received %v, terminating", sig)
	}
	s.GracefulStop()
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func initDatabase(driver, dsn string, debug bool) (*database.Connection, error) {
	opts := []database.Option{}

	if debug {
		opts = append(opts, database.WithDebug())
	}

	return database.Connect(driver, dsn, opts...)
}
