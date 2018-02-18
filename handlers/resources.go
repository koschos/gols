package handlers

// Output link resource
type getLinkResource struct {
	Slug    string `json:"slug"`
	Url     string `json:"url"`
	UrlHash string `json:"url_hash"`
}

// Input link resource
type createLinkResource struct {
	Url string `json:"url"`
}
