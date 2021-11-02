package config

import (
	"github.com/frankffenn/go-utils/config"
	"github.com/frankffenn/go-utils/db"
	"github.com/frankffenn/go-utils/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"golang.org/x/xerrors"
	"os"
	"path/filepath"
)

const defaultFile = "config.toml"

var (
	Cfg            *AppConfig
	_defaultEngine *xorm.Engine
)

type AppConfig struct {
	API *API
}

type API struct {
	RunMode       string
	ListenAddress string
	JwtUserSecret string
	DBURL         string
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func InitConfig(filePath string) error {
	dir, _ := os.Getwd()
	if filePath == "" {
		filePath = filepath.Join(dir, defaultFile)
	}

	_, err := os.Stat(filePath)
	if err != nil {
		return xerrors.Errorf("stat config file (%s): %w", filePath, err)
	}

	ff, err := config.FromFile(filePath, &AppConfig{})
	if err != nil {
		return xerrors.Errorf("loading config: %w", err)
	}

	Cfg = ff.(*AppConfig)
	if err := initDB(); err != nil {
		return err
	}
	return nil
}

func initDB() error {
	var err error

	_defaultEngine, err = db.OpenDB("mysql", Cfg.API.DBURL)
	if err != nil {
		return err
	}

	return nil
}

func InitLog(logLevel string) error {
	dir, _ := os.Getwd()
	ld := filepath.Join(dir, "log")
	_, err := os.Stat(ld)
	if os.IsNotExist(err) {
		err = os.MkdirAll(ld, os.ModePerm)
	}
	if err != nil {
		return xerrors.Errorf("stat log dir err: %w", err)
	}

	logger, _ := log.NewLogger(filepath.Join(filepath.Base(ld), filepath.Base(os.Args[0])+".log"), logLevel)
	log.SetDefault(logger)

	return nil
}

func Session() *xorm.Session {
	sess := _defaultEngine.NewSession()
	return sess
}
