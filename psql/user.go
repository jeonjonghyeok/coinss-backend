package psql

import (
	"github.com/jeonjonghyeok/coinss-backend/model"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(u model.User) (id int, err error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password),
		bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	err = DB().QueryRow(`INSERT INTO public.users (email, password, name, access_key, secret_key) 
	VALUES ($1, $2, $3, $4, $5 ) RETURNING uid
	`, u.Email, passwordHash, u.Name, u.Accesskey, u.Secretkey).Scan(&id)
	return
}

func IsExistUser(email string) (exists bool, err error) {
	err = DB().QueryRow(`SELECT EXISTS(SELECT 1 FROM public.users where email = $1)`, email).Scan(&exists)
	return
}
func IsExistUserById(id int) (exists bool, err error) {
	err = DB().QueryRow(`SELECT EXISTS(SELECT 1 FROM public.users where uid = $1)`, id).Scan(&exists)
	return
}

func FindLoginUser(email, password string) (id int, err error) {
	var passwordHash string
	DB().QueryRow(`SELECT uid, password FROM public.users
	WHERE email = $1`, email).Scan(&id, &passwordHash)
	if err := bcrypt.CompareHashAndPassword(
		[]byte(passwordHash), []byte(password)); err != nil {
		return 0, err
	}
	return
}

func FindUserById(id int) (user model.User, err error) {
	DB().QueryRow(`SELECT email, password, name, access_key, secret_key FROM public.users WHERE uid = $1`, id).Scan(&user.Email, &user.Password, &user.Name, &user.Accesskey, &user.Secretkey)
	return
}
