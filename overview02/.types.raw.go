package overview02

import (
	"context"
	"database/sql"
	"net/http"
	"time"
)

//go:generate gozz run -p "wire" ./

type (
	// root config for unmarshal config file
	Config struct {
		Server ServerConfig `yaml:"server"`
		Sql    SqlConfig    `yaml:"sql"`
		Redis  RedisConfig  `yaml:"redis"`
	}

	// http server config
	ServerConfig struct {
		Addr string `yaml:"addr"`
	}

	// sql config
	SqlConfig struct {
		Dsn string `yaml:"dsn"`
	}

	// redis config
	RedisConfig struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
)

// provide http server from server config
func ProvideHttpServer(config ServerConfig) *http.Server {
	return &http.Server{
		Addr: config.Addr,
	}
}

// interface of sql connection
type SqlConn interface {
	QueryContext(ctx context.Context, statement string, args ...interface{}) (rows *sql.Rows, err error)
}

// interface of key value store
type Store interface {
	Get(ctx context.Context, key string) (value []byte, err error)
	Set(ctx context.Context, key string, value []byte, exp time.Duration) (err error)
}

// provide sql connection from sql config
func ProvideSql(config SqlConfig) (*sql.DB, error) {
	panic("not implemented")
}

// provide kv store from redis config
func ProvideRedisStore(config RedisConfig) (Store, error) {
	panic("not implemented")
}

// biz service handler
type ServiceHandler interface {
	GetInt(ctx context.Context) (int, error)
	GetString(ctx context.Context) (string, error)
}

// implement of server handler
type ServerHandlerImpl struct {
	Sql   SqlConn
	Store Store
}

func (impl *ServerHandlerImpl) GetInt(ctx context.Context) (int, error) {
	panic("not implemented")
}

func (impl *ServerHandlerImpl) GetString(ctx context.Context) (string, error) {
	panic("not implemented")
}

// the entry of application
type Application interface {
	Run()
}

// web application implement
type application struct {
	Server  *http.Server
	Handler ServiceHandler
}

func (application application) Run() {
	panic("not implemented")
}
