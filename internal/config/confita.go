package config

import (
	"context"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
)

type ConfitaProvider struct {
	loc string
	*baseProvider
}

func NewConfitaProvider(loc string) *ConfitaProvider {
	return &ConfitaProvider{
		loc:          loc,
		baseProvider: newBaseProvider(),
	}
}

func (p *ConfitaProvider) Load() error {
	return confita.NewLoader(
		file.NewOptionalBackend(p.loc),
	).Load(context.Background(), p.instance)
}
