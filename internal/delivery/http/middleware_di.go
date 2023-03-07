package http

import (
	"github.com/krobus00/auth-service/internal/model"
)

// InjectTokenRepo :nodoc:
func (m *HTTPMiddleware) InjectTokenRepo(repo model.TokenRepository) error {
	if repo == nil {
		return model.ErrTokenInvalid
	}
	m.tokenRepo = repo
	return nil
}
