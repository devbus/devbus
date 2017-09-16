package common

import (
	"github.com/go-ini/ini"
	"github.com/pkg/errors"
)

type AppConfig struct {
	Addr       string
	DB         string
	DBUser     string
	DBPassword string
	DBHost     string
}

var Config AppConfig

func Parse(confPath string) error {
	var section *ini.Section
	var key *ini.Key
	cfg, err := ini.Load(confPath)
	if err != nil {
		return errors.Wrap(err, "failed to parse config")
	}

	// parse 'server' section
	if section, err = cfg.GetSection("server"); err != nil {
		return errors.New("section 'server' is required")
	}

	if key, err = section.GetKey("addr"); err != nil {
		return errors.New("field 'addr' in section 'server' is required")
	}
	Config.Addr = key.Value()

	// parse 'db' section
	if section, err = cfg.GetSection("db"); err != nil {
		return errors.New("section 'db' is required")
	}

	if key, err = section.GetKey("DB"); err != nil {
		return errors.New("field 'db' in section 'db' is required")
	}
	Config.DB = key.Value()

	if key, err = section.GetKey("user"); err != nil {
		return errors.New("field 'user' in section 'db' is required")
	}
	Config.DBUser = key.Value()

	if key, err = section.GetKey("password"); err != nil {
		return errors.New("field 'password' in section 'db' is required")
	}
	Config.DBPassword = key.Value()

	if key, err = section.GetKey("host"); err != nil {
		return errors.New("field 'host' in section 'db' is required")
	}
	Config.DBHost = key.Value()

	return nil
}