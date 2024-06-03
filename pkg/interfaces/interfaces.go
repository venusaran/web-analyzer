package interfaces

type TargetURL struct {
	URL string `json:"url"`
}

type PageData struct {
	HTMLVersion       string          `json:"html_version"`
	Title             string          `json:"title"`
	Headings          map[string]int  `json:"headings"`
	InternalLinks     int             `json:"internal_links"`
	ExternalLinks     int             `json:"external_links"`
	InaccessibleLinks int             `json:"inaccessible_links"`
	LoginForm         bool            `json:"login_form"`
	AccessibleURLs    map[string]bool `json:"accessible_urls"`
}
