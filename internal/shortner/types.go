package shortner

type Shortner struct {
	HashUrl     string `json:"hash_url"`
	OriginalUrl string `json:"original_url"`
	Clicks      int    `json:"clicks"`
	UserID      int    `json:"user_id"`
}

type RequestDTO struct {
	OriginalUrl string `json:"original_url"`
}

type responseDTO struct {
	HashUrl     string `json:"hash_url"`
	OriginalUrl string `json:"original_url"`
	Clicks      int    `json:"clicks"`
	UserID      int    `json:"user_id"`
	BaseUrl     string `json:"base_url"`
}

func ResponseDTO(data Shortner) *responseDTO {
	return &responseDTO{
		OriginalUrl: data.OriginalUrl,
		HashUrl:     data.HashUrl,
		Clicks:      data.Clicks,
		UserID:      data.UserID,
		BaseUrl:     "http://encurtador-caseiro.com/",
	}
}
