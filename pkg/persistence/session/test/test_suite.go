package test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/lab5e/lmqtt/pkg/entities"
	"github.com/lab5e/lmqtt/pkg/persistence/session"
)

// Suite runs the tests on a session.Store type
func Suite(t *testing.T, store session.Store) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var tt = []*entities.Session{
		{
			ClientID: "client",
			Will: &entities.Message{
				Topic:   "topicA",
				Payload: []byte("abc"),
			},
			WillDelayInterval: 1,
			ConnectedAt:       time.Unix(1, 0),
			ExpiryInterval:    2,
		}, {
			ClientID:          "client2",
			Will:              nil,
			WillDelayInterval: 0,
			ConnectedAt:       time.Unix(2, 0),
			ExpiryInterval:    0,
		},
	}
	for _, v := range tt {
		a.Nil(store.Set(v))
	}
	for _, v := range tt {
		sess, err := store.Get(v.ClientID)
		a.Nil(err)
		a.EqualValues(v, sess)
	}
	var sess []*entities.Session
	err := store.Iterate(func(session *entities.Session) bool {
		sess = append(sess, session)
		return true
	})
	a.Nil(err)
	a.ElementsMatch(sess, tt)
}
