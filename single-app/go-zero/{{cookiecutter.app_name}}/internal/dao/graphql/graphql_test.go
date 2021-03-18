package graphql

import (
	"os"
	"testing"

	"github.com/better-go/pkg/log"
	"github.com/tal-tech/go-zero/core/conf"

	"{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/proto/config"
)

var (
	testDao *Dao // for unit test case use
)

func TestMain(m *testing.M) {
	// default config path:
	configFile := "../../../configs/configs.yaml"

	// parse config:
	var cfg config.Config
	conf.MustLoad(configFile, &cfg)
	log.Infof("test dao config: %+v", cfg.GraphQL)

	// new:
	testDao = New(cfg.GraphQL)
	defer testDao.Close()

	if code := m.Run(); code != 0 {
		os.Exit(code)
	}
}
