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
	err = db.QueryRow(`INSERT INTO public.user (email, password, name, phone_number, access_key, secret_key) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING uid
	`, u.Email, passwordHash, u.Name, u.PhoneNumber, u.Accesskey, u.Secretkey).Scan(&id)
	return
}

func IsExistUser(email string) (exists bool, err error) {
	err = db.QueryRow(`SELECT EXISTS(SELECT 1 FROM public.user where email = $1)`, email).Scan(&exists)
	return

}

func FindUser(email, password string) (id int, err error) {
	var passwordHash string
	err = db.QueryRow(`SELECT uid, password FROM public.user
	WHERE email = $1`, email).Scan(&id, &passwordHash)
	if err != nil {
		return 0, err
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(passwordHash), []byte(password)); err != nil {
		return 0, err
	}
	return
}

func FindUserKey(id int) (access_key, secret_key string, err error) {
	err = db.QueryRow(`SELECT access_key, secret_key FROM public.user WHERE uid = $1`, id).Scan(&access_key, &secret_key)
	if err != nil {
		return "", "", err
	}

	return
}
