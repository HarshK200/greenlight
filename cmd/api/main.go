package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/harshk200/greenlight/internal/data"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string // (Can be production or development or staging) This will be used for testing later
	db   struct {
		dns          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
}

func openDB(cfg *config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dns)
	if err != nil {
		return nil, err
	}

	// setting connection limits
	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	// the context cancels/timesout after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// testing the db connection by pinging it
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// loading Environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env variables", err)
	}

	var cfg config

	// CLI flags
	flag.IntVar(&cfg.port, "addr", 4000, "the port/addr on which server will run")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	// db connection flags
	flag.StringVar(&cfg.db.dns, "db-dns", os.Getenv("POSTGRES_URL"), "DNS for the database (postgres-db)")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-timeout", "15m", "PostgreSQL max connection idle time")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// DB connection pool setup
	db, err := openDB(&cfg)
	if err != nil {
		log.Fatal("Error connecting to the db", err)
	}
	defer db.Close()

	logger.Println("postgres DB connection established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}
