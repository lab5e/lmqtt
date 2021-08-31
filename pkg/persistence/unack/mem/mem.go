package mem

import (
	"github.com/lab5e/lmqtt/pkg/packets"
	"github.com/lab5e/lmqtt/pkg/persistence/unack"
)

var _ unack.Store = (*Store)(nil)

// Store is the memory store implementation for the unack store
type Store struct {
	clientID     string
	unackpublish map[packets.PacketID]struct{}
}

type Options struct {
	ClientID string
}

func New(opts Options) *Store {
	return &Store{
		clientID:     opts.ClientID,
		unackpublish: make(map[packets.PacketID]struct{}),
	}
}

func (s *Store) Init(cleanStart bool) error {
	if cleanStart {
		s.unackpublish = make(map[packets.PacketID]struct{})
	}
	return nil
}

func (s *Store) Set(id packets.PacketID) (bool, error) {
	if _, ok := s.unackpublish[id]; ok {
		return true, nil
	}
	s.unackpublish[id] = struct{}{}
	return false, nil
}

func (s *Store) Remove(id packets.PacketID) error {
	delete(s.unackpublish, id)
	return nil
}
