package bar

import (
	repo "boilerplate/internal/repositiories/bar"
	"context"
)

//go:generate mockery --name=Repository
type Repository interface {
	FindUserByBars(ctx context.Context, bar string) ([]*repo.User, error)
}

type service struct {
	repo Repository
}

func New(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

// Bar makes database request for some data
func (s *service) Bar(ctx context.Context, bar string) ([]string, error) {
	bars, err := s.repo.FindUserByBars(ctx, bar)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(bars))
	for _, bar := range bars {
		names = append(names, bar.Name)
	}
	return names, nil
}
