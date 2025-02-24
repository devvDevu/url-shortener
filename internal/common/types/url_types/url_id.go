package url_types

type UrlId int

func (ct UrlId) Int() int {
	return int(ct)
}
