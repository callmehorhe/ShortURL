package service

import (
	"errors"
	"net/url"
	"strings"
	"time"

	"math/rand"

	"github.com/callmehorhe/shorturl/api/pkg/models"
	"github.com/callmehorhe/shorturl/api/pkg/repository"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	length   = int64(len(alphabet))
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateURL(url string) (string, error) {
	if !isValidUrl(url) {
		return "", errors.New("incorrect URL")
	}


	newUrl := s.repo.IsCreated(url)
	if newUrl != "" {
		return newUrl, nil
	}

	rand.Seed(time.Now().UnixNano())

	for used := true; used; used = s.repo.IsUsed(newUrl) {
		newUrl = "http://shorturl.ru/" + Encode(rand.Int63())
	}

	u := models.UrlModel{
		OldURL: url,
		NewURL: newUrl,
	}
	err := s.repo.CreateURL(u)
	if err != nil {
		return "", err
	}
	return newUrl, nil
}

func (s *Service) GetURL(url string) (string, error) {
	return s.repo.GetURL(url)
}

func Encode(num int64) string {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(10)

	for ; num > 0; num = num / length {
		encodedBuilder.WriteByte(alphabet[(num % length)])
	}

	return encodedBuilder.String()[:10]
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
