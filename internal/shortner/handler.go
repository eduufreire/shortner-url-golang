package shortner

import (
	"encoding/base64"
	"encoding/json"
	"hash/maphash"
	"net/http"
	"strings"

	"github.com/eduufreire/url-shortner/internal/auth"
)

type shortnerHandler struct {
	service ShortnerService
}

func NewShortnerHandler(service ShortnerService) ShortnerHandler {
	return &shortnerHandler{
		service: service,
	}
}

func hashUrl(url string) string {
	hash := maphash.Hash{}
	hash.SetSeed(maphash.MakeSeed())
	hash.WriteString(url)
	result := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return strings.ReplaceAll(result, "=", "")
}

func (h *shortnerHandler) Create(w http.ResponseWriter, r *http.Request) {

	// token := r.Header.Get("Authorization")
	// splitedToken := strings.Split(token, " ")
	// user, err := auth.VerifyToken(splitedToken[1])
	// if err != nil {
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	body := RequestDTO{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte(err.Error()))
		return
	}

	created := h.service.Create(body.OriginalUrl, 13)

	response, err := json.Marshal(created)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(response)

}

func (h *shortnerHandler) Get(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	splitedToken := strings.Split(token, " ")
	_, err := auth.VerifyToken(splitedToken[1])
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	hash := r.PathValue("hash")
	shortner := h.service.GetByHash(hash)

	http.Redirect(w, r, shortner.OriginalUrl, 302)
}
