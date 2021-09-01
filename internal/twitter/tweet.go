package twitter

type Tweet struct {
}

func (t Tweet) GetHeader(url string) {
	println(url)
}
