package url_types

type UrlOriginal string

func (ct UrlOriginal) String() string {
	return string(ct)
}
