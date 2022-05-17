package orders

type APIKey struct {
	Key string
}

func GetKey() string {
	api := APIKey{}
	api.Key = "appid=YOUR_API_KEY_HERE"
	return api.Key
}
