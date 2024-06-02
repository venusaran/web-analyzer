package interfaces

type TargetURL struct {
	URL string `json:"url"`
}

type PageData struct {
	HTMLVersion       string
	Title             string
	Headings          map[string]int
	InternalLinks     int
	ExternalLinks     int
	InaccessibleLinks int
	LoginForm         bool
	AccessibleURLs    map[string]bool
}
