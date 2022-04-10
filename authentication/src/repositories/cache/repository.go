package cache

type Cache interface {
	Auth
}

type cache struct {
	*auth
}

func NewCacheRepository() Cache {
	return &cache{
		auth: &auth{},
	}
}
