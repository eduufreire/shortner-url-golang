package shortner

import "net/http"

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
}

type ShortnerHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type ShortnerService interface {
	Create(originalUrl string, userID int) responseDTO
	GetByHash(hashUrl string) responseDTO
}

type ShortnerRepository interface {
	Save(data Shortner)  error
	GetByHash(hashUrl string) (*Shortner, error)
}
