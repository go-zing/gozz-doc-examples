package option02

//go:generate gozz run -p "option" ./

// +zz:option:type={{ .Name }}Option
type (
	Config struct {
		// connect host
		Host string
		// connect port
		Port string
		// database username
		Username string
		// database password
		Password string
	}

	Config2 struct {
		// config url
		Url string
	}
)
