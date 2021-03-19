package queue

import (
	"os"
	"testing"

	"github.com/better-go/pkg/log"
	"github.com/tal-tech/go-zero/core/conf"

	"{{cookiecutter.go_module_name}}/app/basic/demo/proto/config"
)

var (
	testDomain *Domain // for unit test case use
)

func TestMain(m *testing.M) {
	// default config path:
	configFile := "../../../configs/configs.yaml"

	// parse config:
	var cfg config.Config
	conf.MustLoad(configFile, &cfg)
	log.Infof("test dao config: %+v", cfg.GraphQL)

	// new:
	testDomain = NewDomain(cfg, true)
	defer testDomain.Close()

	if code := m.Run(); code != 0 {
		os.Exit(code)
	}
}
