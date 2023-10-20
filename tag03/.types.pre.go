package tag03

import (
	"time"
)

//go:generate gozz run -p "tag" ./

// +zz:tag:json,bson:{{ snake .FieldName }}
type (
	User struct {
		Id        string
		Name      string
		Address   string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	Book struct {
		Id        string
		Title     string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	Order struct {
		Id        string
		UserId    string
		BookId    string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
