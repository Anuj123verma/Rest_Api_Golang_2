package controller

func Getlink(latitude string, longitude string, key string) string {
	return "https://revgeocode.search.hereapi.com/v1/revgeocode?at=" + latitude + "%2C" + longitude + "&lang=en-US&apikey=" + key
}
