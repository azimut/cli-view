package twitter

type Embedded struct {
	Url          string
	AuthorName   string `json:"author_name"`
	AuthorUrl    string `json:"author_url"`
	Html         string
	width        int
	height       int
	kind         string `json:"type"`
	cacheAge     int    `json:cache_age`
	providerName string `json:provider_name`
	providerUrl  string `json:provider_url`
	version      float32
}
