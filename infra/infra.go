package infra

import (
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Infra interface {
	Config() *viper.Viper
	SQLDB() *sqlx.DB
}

type infraCtx struct {
	configFilePath string
}

func New(configFilePath string) Infra {
	fpath, err := filepath.Abs(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(fpath); err != nil {
		log.Fatalf("config file path %s not found", configFilePath)
	}

	return &infraCtx{
		configFilePath: configFilePath,
	}
}

var (
	cfgOnce sync.Once
	cfg     *viper.Viper
)

func (c *infraCtx) Config() *viper.Viper {
	cfgOnce.Do(func() {
		v := viper.New()
		v.SetConfigFile(c.configFilePath)

		if err := v.ReadInConfig(); err != nil {
			log.Fatalf("failed to read config file: %v", err)
		}

		cfg = v
	})

	return cfg
}

var (
	sqlDBOnce sync.Once
	sqlDB     *sqlx.DB
)

func (c *infraCtx) SQLDB() *sqlx.DB {
	sqlDBOnce.Do(func() {
		pgConfig := c.Config().Sub("postgres")
		conCfg := pgx.ConnConfig{
			Host:     pgConfig.GetString("host"),
			Port:     uint16(pgConfig.GetInt("port")),
			Database: pgConfig.GetString("database"),
			User:     pgConfig.GetString("user"),
			Password: pgConfig.GetString("password"),
		}

		if pgConfig.GetBool("log") {
			logLevel := strings.ToLower(pgConfig.GetString("log_level"))
			var lvl pgx.LogLevel
			switch logLevel {
			case "trace":
				lvl = pgx.LogLevelTrace
			case "debug":
				lvl = pgx.LogLevelDebug
			case "info":
				lvl = pgx.LogLevelInfo
			case "warn":
				lvl = pgx.LogLevelWarn
			case "error":
				lvl = pgx.LogLevelError
			case "none":
				lvl = pgx.LogLevelNone
			default:
				lvl = pgx.LogLevelInfo
			}
			conCfg.LogLevel = lvl
		}
		db := stdlib.OpenDB(conCfg)
		db.SetConnMaxLifetime(time.Duration(pgConfig.GetInt("con_max_lifetime")) * time.Second)
		db.SetMaxIdleConns(pgConfig.GetInt("con_max_idle"))
		db.SetMaxOpenConns(pgConfig.GetInt("con_max_open"))

		dbx := sqlx.NewDb(db, "postgres")
		sqlDB = dbx
	})
	return sqlDB
}
