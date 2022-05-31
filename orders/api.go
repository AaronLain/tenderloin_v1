package orders

type APIKey struct {
	Key string
}

func GetKey() string {
	api := APIKey{}
	api.Key = "appid=721b00374b19b0362abd9ab1c1680fba"
	return api.Key
}
