package psql

func Favorite(id int, names string) (err error) {
	_, err = db.Exec(`UPDATE public.users SET favorites=$1 WHERE uid=$2`, names, id)
	return
}

func GetFavorites(id int) (name string) {
	db.QueryRow(`SELECT favorites FROM public.users WHERE uid=$1`, id).Scan(&name)
	return
}

func SetSearch(search string) (err error) {
	count := getSearch(search)
	if count == 0 {
		_, err = db.Exec(`INSERT INTO public.search (search, count) VALUES ($1, $2)`, search, 1)
	} else {
		_, err = db.Exec(`UPDATE public.search SET count=$1 WHERE search=$2`, count+1, search)
	}
	return
}

func getSearch(search string) (count int) {
	db.QueryRow(`SELECT count FROM public.search WHERE search=$1`, search).Scan(&count)
	return
}
