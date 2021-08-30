package config

import (
	"github.com/pkg/errors"
)

type PersistenceType = string

const (
	PersistenceTypeMemory PersistenceType = "memory"
)

var (
	// DefaultPersistenceConfig is the default value of Persistence
	DefaultPersistenceConfig = Persistence{
		Type: PersistenceTypeMemory,
	}
)

// Persistence is the config of backend persistence.
type Persistence struct {
	// Type is the persistence type.
	// If empty, use "memory" as default.
	Type PersistenceType `yaml:"type"`
}

func (p *Persistence) Validate() error {
	if p.Type != PersistenceTypeMemory {
		return errors.New("invalid persistence type")
	}
	return nil
}
