package psql

import "github.com/jeonjonghyeok/coinss-backend/vo"

func CreateUser(u vo.User) (err error) {
	_, err = db.Exec(`INSERT INTO public.user (email, password, name, phone_number, access_key, secret_key)
	VALUES ($1, $2, $3, $4, $5, $6)
	`, u.Email, u.Password, u.Name, u.PhoneNumber, u.Accesskey, u.Secretkey)
	return
}
