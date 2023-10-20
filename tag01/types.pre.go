package tag01

import (
	"time"
)

//go:generate gozz run -p "tag" ./

// +zz:tag:json,bson:{{ snake .FieldName }}
type User struct {
	Id        string
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
