package tag01

import (
	"time"
)

//go:generate gozz run -p "tag" ./

// +zz:tag:json,bson:{{ snake .FieldName }}
type User struct {
	Id        string    `bson:"id" json:"id"`
	Name      string    `bson:"name" json:"name"`
	Address   string    `bson:"address" json:"address"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
