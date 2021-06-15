package psql

import (
	"github.com/jeonjonghyeok/coinss-backend/vo"
)

func CreateUser(u vo.User) (id int, err error) {
	err = db.QueryRow(`INSERT INTO public.user (email, password, name, phone_number, access_key, secret_key) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING uid
	`, u.Email, u.Password, u.Name, u.PhoneNumber, u.Accesskey, u.Secretkey).Scan(&id)
	return
}

func IsExistUser(email string) (exists bool, err error) {
	err = db.QueryRow(`SELECT EXISTS(SELECT 1 FROM public.user where email = $1)`, email).Scan(&exists)
	return

}
