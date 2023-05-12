package link

type (
	Link struct {
		ID          string `json:"id"`
		OriginalURL string `json:"original_url"`
	}
	AddLink struct {
		URL string `json:"url" validate:"required,url_encoded"`
	}
)
