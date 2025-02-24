package url_types

type UrlCode string

func (ct UrlCode) String() string {
	return string(ct)
}
