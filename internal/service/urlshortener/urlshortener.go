package urlshortener

import (
	"context"
	"url-shortener/internal/dto"
)

func (s *service) EncodeUrl(ctx context.Context, input dto.EncodeURLReq) (dto.EncodeURLResp, error) {
	return dto.EncodeURLResp{}, nil
}

func (s *service) DecodeUrl(ctx context.Context, input dto.DecodeURLReq) (dto.DecodeURLResp, error) {
	return dto.DecodeURLResp{}, nil
}
