package tag02

import (
	"time"
)

//go:generate gozz run -p "tag" ./

// +zz:tag:json,bson:{{ camel .FieldName }}
type User struct {
	Id        string    `json:"id" bson:"id"`
	Name      string    `json:"name" bson:"name"`
	Address   string    `json:"address" bson:"address"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
