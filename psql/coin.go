package psql

func Favorite(id int, names string) (err error) {
	_, err = DB().Exec(`UPDATE public.users SET favorites=$1 WHERE uid=$2`, names, id)
	return
}

func GetFavorites(id int) (name string) {
	DB().QueryRow(`SELECT favorites FROM public.users WHERE uid=$1`, id).Scan(&name)
	return
}

func SetSearch(search string) (err error) {
	count := getSearch(search)
	if count == 0 {
		_, err = DB().Exec(`INSERT INTO public.search (search, count) VALUES ($1, $2)`, search, 1)
	} else {
		_, err = DB().Exec(`UPDATE public.search SET count=$1 WHERE search=$2`, count+1, search)
	}
	return
}

func getSearch(search string) (count int) {
	DB().QueryRow(`SELECT count FROM public.search WHERE search=$1`, search).Scan(&count)
	return
}
