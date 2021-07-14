package psql

func Favorite(id int, names string) (err error) {
	_, err = db.Exec(`UPDATE public.users SET favorites=$1 WHERE uid=$2`, names, id)
	return
}

func GetFavorites(id int) (name string) {
	db.QueryRow(`SELECT favorites FROM public.users WHERE uid=$1`, id).Scan(&name)
	return
}
