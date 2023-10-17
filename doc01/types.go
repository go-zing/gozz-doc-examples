package doc01

//go:generate gozz run -p "doc" ./

// +zz:doc
type (
	// uuid to define unique entity
	UUID string

	// abstract type of entity
	Entity interface {
		// get entity uuid
		Id() UUID
		// get entity name
		Name() string
	}

	// entity for users
	User struct {
		// user uuid
		Id UUID
		// user name
		Name string
		// user gender
		Gender string
	}

	// entity for books
	Book struct {
		// book uuid
		Id UUID
		// book name
		Name string
		// book type
		Type string
	}
)

// +zz:doc:label=gender_type
const (
	// female
	GenderFemale = "female"
	// male
	GenderMale = "male"
	// other
	GenderOther = "other"
)

// +zz:doc:label=book_type
const (
	// books for children
	BookTypeChildren = "children"
	// books for tech
	BookTypeTech = "tech"
	// books for cook
	BookTypeCook = "cook"
)
