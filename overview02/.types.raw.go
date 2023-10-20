package overview02

import (
	"context"
	"database/sql"
	"net/http"
	"time"
)

type (
	// root config for unmarshal config file
	Config struct {
		Server ServerConfig
		Sql    SqlConfig
		Redis  RedisConfig
	}

	// http server config
	ServerConfig struct {
		Addr string
	}

	// sql config
	SqlConfig struct {
		Dsn string
	}

	// redis config
	RedisConfig struct {
		Host string
		Port string
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
func ProvideSql(config SqlConfig) (SqlConn, error) {
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

// the entry of application
type Application interface {
	Run()
}

// web application implement
type application struct {
	Server  *http.Server
	Handler ServiceHandler
}
