package backend

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	EnvPrefix       string
	ApiEntrypoint   string
	ServerBindAddr  string
	ServerBindPort  uint64
	MetricsBindAddr string
	MetricsBindPort uint64
}

func NewConfig() *Config {
	cfg := &Config{
		EnvPrefix:       "APP_",
		ApiEntrypoint:   "http://localhost:8000",
		ServerBindAddr:  "127.0.0.1",
		ServerBindPort:  8080,
		MetricsBindAddr: "127.0.0.1",
		MetricsBindPort: 2112,
	}

	return cfg
}

func (c *Config) PopulateFromEnv() error {

	// c.EnvPrefix can not be asseigned from ENV

	var arg string
	var err error
	var val_ui uint64

	arg = c.EnvPrefix + "API_ENTRYPOINT"
	if val, found := os.LookupEnv(arg); found {
		c.ApiEntrypoint = val
	}

	arg = c.EnvPrefix + "SERVER_BIND_ADDR"
	if val, found := os.LookupEnv(arg); found {
		c.ServerBindAddr = val
	}

	arg = c.EnvPrefix + "SERVER_BIND_PORT"
	if val, found := os.LookupEnv(arg); found {
		if val_ui, err = strconv.ParseUint(val, 10, 64); err != nil {
			return errors.New(fmt.Sprintf("ERROR argument %s wrong type %T expects unsigned integer", arg, arg))
		}
		c.ServerBindPort = val_ui
	}

	arg = c.EnvPrefix + "METRICS_BIND_ADDR"
	if val, found := os.LookupEnv(arg); found {
		c.MetricsBindAddr = val
	}

	arg = c.EnvPrefix + "METRICS_BIND_PORT"
	if val, found := os.LookupEnv(arg); found {
		if val_ui, err = strconv.ParseUint(val, 10, 64); err != nil {
			return errors.New(fmt.Sprintf("ERROR argument %s wrong type %T expects unsigned integer", arg, arg))
		}
		c.MetricsBindPort = val_ui
	}

	return nil
}
