package vo

type User struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Secretkey   string `json:"secret_key"`
	Accesskey   string `json:"access_key"`
}
