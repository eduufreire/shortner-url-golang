package shortner

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash/maphash"
	"net/http"
	"strings"
)

type handler struct {
	repository *repository
}

func Handler(repository *repository) *handler {
	return &handler{
		repository: repository,
	}
}

func hashUrl(url string) string {
	hash := maphash.Hash{}
	hash.SetSeed(maphash.MakeSeed())
	hash.WriteString(url)
	result := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return strings.ReplaceAll(result, "=", "")
}

func (h *handler) CreateUrl(w http.ResponseWriter, r *http.Request) {

	body := RequestDTO{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	hashUrl := hashUrl(body.OriginalUrl)

	data := Shortner{
		HashUrl:     hashUrl,
		OriginalUrl: body.OriginalUrl,
		Clicks:      0,
	}

	if err := h.repository.Save(data); err != nil {
		fmt.Println("deu erro ao salvar")
		w.Write([]byte(err.Error()))
		return
	}

	dto := ResponseDTO(data)
	response, err := json.Marshal(dto)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(response)

}

func (h *handler) GetUrl(w http.ResponseWriter, r *http.Request) {
	hash := r.PathValue("hash")

	data, err := h.repository.GetByHash(hash)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if url := data.OriginalUrl; url == "" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, data.OriginalUrl, 302)
}
