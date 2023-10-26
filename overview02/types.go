package overview02

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

//go:generate gozz run -p "wire" ./
//go:generate sh -c "go run ./cmd/stgragh | dot -Tpng -o structure.png"
//go:generate sh -c "go run ./cmd/stgragh | dot -Tsvg -o structure.svg"

type (
	// root config for unmarshal config file
	// +zz:wire:field=*
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
// +zz:wire
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
// +zz:wire:bind=SqlConn
func ProvideSql(config SqlConfig) (*sql.DB, error) {
	return sql.Open("mysql", config.Dsn)
}

// provide kv store from redis config
// +zz:wire:bind=redis.Cmdable
func ProvideRedisStore(config RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.Host, config.Port),
	})
	return rdb, nil
}

// +zz:wire:bind=Store
type RedisStore struct {
	redis.Cmdable
}

func (s RedisStore) Set(ctx context.Context, key string, value []byte, exp time.Duration) (err error) {
	return s.Cmdable.Set(ctx, key, value, exp).Err()
}

func (s RedisStore) Get(ctx context.Context, key string) (value []byte, err error) {
	return s.Cmdable.Get(ctx, key).Bytes()
}

// biz service handler
type ServiceHandler interface {
	GetInt(ctx context.Context) (int, error)
	GetString(ctx context.Context) (string, error)
}

// implement of server handler
// +zz:wire:bind=ServiceHandler:aop
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
// +zz:wire:inject=./:param=*Config
type Application interface {
	Run()
}

// web application implement
// +zz:wire:bind=Application
type application struct {
	Server  *http.Server
	Handler ServiceHandler
}

func (application application) Run() {
	panic("not implemented")
}
