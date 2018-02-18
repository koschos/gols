package handlers

// Output link resource
type GetLinkResource struct {
	Slug    string `json:"slug"`
	Url     string `json:"url"`
	UrlHash string `json:"url_hash"`
}

// Input link resource
type CreateLinkResource struct {
	Url string `json:"url"`
}
