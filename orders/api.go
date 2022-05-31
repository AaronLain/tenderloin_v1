package orders

type APIKey struct {
	Key string
}

func GetKey() string {
	api := APIKey{}
	api.Key = "appid=e9b2b6a4c6346f1ea46a948b13f129a5"

	return api.Key
}
