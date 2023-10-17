package option01

//go:generate gozz run -p "option" ./

// +zz:option
type Config struct {
	// connect host
	Host string
	// connect port
	Port string
	// database username
	Username string
	// database password
	Password string
}
