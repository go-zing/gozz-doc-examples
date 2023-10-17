package tag05

import (
	"time"
)

//go:generate gozz run -p "tag" ./

// +zz:tag:json,bson:{{ snake .FieldName }}
type (
	UserStruct struct {
		Id        string    `bson:"id" json:"id"`
		Name      string    `bson:"name" json:"name"`
		Address   string    `bson:"address" json:"address"`
		CreatedAt time.Time `bson:"created_at" json:"created_at"`
		UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
		Friends   []struct {
			Id        string    `bson:"id" json:"id"`
			Name      string    `bson:"name" json:"name"`
			Address   string    `bson:"address" json:"address"`
			CreatedAt time.Time `bson:"created_at" json:"created_at"`
			UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
		} `bson:"friends" json:"friends"`
	}

	UserMap map[string]struct {
		Id        string    `bson:"id" json:"id"`
		Name      string    `bson:"name" json:"name"`
		Address   string    `bson:"address" json:"address"`
		CreatedAt time.Time `bson:"created_at" json:"created_at"`
		UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	}

	UserSlice []struct {
		Id        string    `bson:"id" json:"id"`
		Name      string    `bson:"name" json:"name"`
		Address   string    `bson:"address" json:"address"`
		CreatedAt time.Time `bson:"created_at" json:"created_at"`
		UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	}

	UserInterface interface {
		User() struct {
			Id        string    `bson:"id" json:"id"`
			Name      string    `bson:"name" json:"name"`
			Address   string    `bson:"address" json:"address"`
			CreatedAt time.Time `bson:"created_at" json:"created_at"`
			UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
		}
	}

	UserFunc func(struct {
		Id        string    `bson:"id" json:"id"`
		Name      string    `bson:"name" json:"name"`
		Address   string    `bson:"address" json:"address"`
		CreatedAt time.Time `bson:"created_at" json:"created_at"`
		UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	})
)
