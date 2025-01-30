package types

// User struct with validation tags
type User struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=8"`
	Id       string `json:"id,omitempty"`
}

// UserResponse struct for responses
type UserResponse struct {
	Email string `json:"email,omitempty"`
	Id    string `json:"id,omitempty"`
}

type Config struct {
	Port        string
	Db_User     string
	Db_Pwd      string
	Db_Port     string
	Db_URL      string
	Redis_URL   string
	Schema_Path string
	Kafka_URL   string
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserLoginResponse struct {
	Email string `json:"email,omitempty"`
	Id    string `json:"id,omitempty"`
}
