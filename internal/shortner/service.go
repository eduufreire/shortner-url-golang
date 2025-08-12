package shortner

import (
	"fmt"
	"log"
)

type shortnerService struct {
	repository ShortnerRepository
}

func NewShortnerService(repository ShortnerRepository) ShortnerService {
	return &shortnerService{
		repository: repository,
	}
}

func (s *shortnerService) Create(originalUrl string, userID int) responseDTO {
	hashUrl := hashUrl(originalUrl)

	data := Shortner{
		HashUrl:     hashUrl,
		OriginalUrl: originalUrl,
		Clicks:      0,
		UserID:      userID,
	}

	err := s.repository.Save(data)
	if err != nil {
		log.Fatal("erro")
	}

	createdShortner, err := s.repository.GetByHash(hashUrl)
	fmt.Println(createdShortner)

	return responseDTO{
		HashUrl:     createdShortner.HashUrl,
		OriginalUrl: createdShortner.OriginalUrl,
		Clicks:      createdShortner.Clicks,
		UserID:      createdShortner.UserID,
	}
}

func (s *shortnerService) GetByHash(hashUrl string) responseDTO {

	data, err := s.repository.GetByHash(hashUrl)
	if err != nil {
		log.Fatal("erro")
	}
	
	if url := data.OriginalUrl; url == "" {
		log.Fatal("nao encontrado")
	}

	return responseDTO{
		HashUrl:     data.HashUrl,
		OriginalUrl: data.OriginalUrl,
		Clicks:      data.Clicks,
		UserID:      data.UserID,
	}
}
