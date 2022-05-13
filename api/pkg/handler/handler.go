package handler

import (
	context "context"

	"github.com/callmehorhe/shorturl/api/pkg/service"
)

type Handler struct {
	serv *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{
		serv: serv,
	}
}

func (h *Handler) Create(ctx context.Context, req *UrlMessage) (*UrlMessage, error) {
	newUrl, err := h.serv.CreateURL(req.GetUrl())
	if err != nil {
		return &UrlMessage{}, err
	}
	return &UrlMessage{
		Url: newUrl,
	}, nil
}

func (h *Handler) Get(ctx context.Context, req *UrlMessage) (*UrlMessage, error) {
	oldUrl, err := h.serv.GetURL(req.GetUrl())
	if err != nil {
		return &UrlMessage{}, err
	}
	return &UrlMessage{
		Url: oldUrl,
	}, nil
}

func (h *Handler) mustEmbedUnimplementedURLServer() {}
