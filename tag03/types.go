package tag03

import (
	"time"
)

//go:generate gozz run -p "tag" ./

// +zz:tag:json,bson:{{ snake .FieldName }}
type (
	User struct {
		Id        string    `bson:"id" json:"id"`
		Name      string    `bson:"name" json:"name"`
		Address   string    `bson:"address" json:"address"`
		CreatedAt time.Time `bson:"created_at" json:"created_at"`
		UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	}

	Book struct {
		Id        string    `bson:"id" json:"id"`
		Title     string    `bson:"title" json:"title"`
		CreatedAt time.Time `bson:"created_at" json:"created_at"`
		UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	}

	Order struct {
		Id        string    `bson:"id" json:"id"`
		UserId    string    `bson:"user_id" json:"user_id"`
		BookId    string    `bson:"book_id" json:"book_id"`
		CreatedAt time.Time `bson:"created_at" json:"created_at"`
		UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	}
)
