package infrastructure

import (
	"time"

	"github.com/jpillora/backoff"
	"github.com/krobus00/auth-service/internal/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB           *gorm.DB
	StopTickerCh chan bool
)

func InitializeDBConn() {
	conn, err := openDBConn(config.DatabaseDSN())
	if err != nil {
		logrus.WithField("databaseDSN", config.DatabaseDSN()).Fatal("failed to connect  database: ", err)
	}

	DB = conn
	StopTickerCh = make(chan bool)

	go checkConnection(time.NewTicker(config.DatabasePingInterval()))

	switch config.LogLevel() {
	case "error":
		DB.Logger = DB.Logger.LogMode(gormLogger.Error)
	case "warn":
		DB.Logger = DB.Logger.LogMode(gormLogger.Warn)
	default:
		DB.Logger = DB.Logger.LogMode(gormLogger.Info)
	}

	logrus.Info("Connection to database Server success...")
}

func checkConnection(ticker *time.Ticker) {
	for {
		select {
		case <-StopTickerCh:
			ticker.Stop()
			return
		case <-ticker.C:
			if _, err := DB.DB(); err != nil {
				reconnectDBConn()
			}
		}
	}
}

func reconnectDBConn() {
	b := backoff.Backoff{
		Factor: float64(config.DatabaseConnReconnectFactor()),
		Jitter: true,
		Min:    config.DatabaseConnReconnectMinJitter(),
		Max:    config.DatabaseConnReconnectMaxJitter(),
	}

	dbRetryAttempts := config.DatabaseRetryAttempts()

	for b.Attempt() < dbRetryAttempts {
		conn, err := openDBConn(config.DatabaseDSN())
		if err != nil {
			logrus.WithField("databaseDSN", config.DatabaseDSN()).Error("failed to connect database: ", err)
		}

		if conn != nil {
			DB = conn
			break
		}
		time.Sleep(b.Duration())
	}

	if b.Attempt() >= dbRetryAttempts {
		logrus.Fatal("maximum retry to connect database")
	}
	b.Reset()
}

func openDBConn(dsn string) (*gorm.DB, error) {
	logrus.Info("trying connect to database")
	psqlDialector := postgres.Open(dsn)
	db, err := gorm.Open(psqlDialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Use(tracing.NewPlugin(tracing.WithoutMetrics()))
	if err != nil {
		return nil, err
	}

	conn, err := db.DB()
	if err != nil {
		logrus.Fatal(err)
	}
	conn.SetMaxIdleConns(config.DatabaseMaxIdleConns())
	conn.SetMaxOpenConns(config.DatabaseMaxOpenConns())
	conn.SetConnMaxLifetime(config.DatabaseConnMaxLifetime())

	return db, nil
}
