package twitter

type Embedded struct {
	Url          string
	AuthorName   string `json:"author_name"`
	AuthorUrl    string `json:"author_url"`
	Html         string
	width        int
	height       int
	Kind         string `json:"type"`
	PacheAge     string `json:"cache_age"`
	ProviderName string `json:"provider_name"`
	ProviderUrl  string `json:"provider_url"`
	version      float32
	Links        []string // NOT in the response
}
