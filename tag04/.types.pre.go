package tag04

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

	// +zz:tag:json,bson:{{ camel .FieldName }}
	Book struct {
		Id        string
		Title     string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	Order struct {
		Id     string
		UserId string
		BookId string
		// +zz:tag:json,bson:{{ upper .FieldName | upper }}
		CreatedAt time.Time
		// +zz:tag:+json:,omitempty
		UpdatedAt time.Time
	}
)
