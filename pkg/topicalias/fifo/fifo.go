package fifo

import (
	"container/list"

	"github.com/lab5e/lmqtt/pkg/config"
	"github.com/lab5e/lmqtt/pkg/lmqtt"
	"github.com/lab5e/lmqtt/pkg/packets"
)

var _ lmqtt.TopicAliasManager = (*Queue)(nil)

func init() {
	lmqtt.RegisterTopicAliasMgrFactory("fifo", New)
}

// New is the constructor of Queue.
func New(config config.Config, maxAlias uint16, clientID string) lmqtt.TopicAliasManager {
	return &Queue{
		clientID: clientID,
		topicAlias: &topicAlias{
			max:   int(maxAlias),
			alias: list.New(),
			index: make(map[string]uint16),
		},
	}
}

// Queue is the fifo queue which store all topic alias for one client
type Queue struct {
	clientID   string
	topicAlias *topicAlias
}
type topicAlias struct {
	max   int
	alias *list.List
	// topic name => alias
	index map[string]uint16
}
type aliasElem struct {
	topic string
	alias uint16
}

// Check checks if the message can be published
func (q *Queue) Check(publish *packets.Publish) (alias uint16, exist bool) {
	topicName := string(publish.TopicName)
	// alias exist
	if a, ok := q.topicAlias.index[topicName]; ok {
		return a, true
	}
	l := q.topicAlias.alias.Len()
	// alias has been exhausted
	if l == q.topicAlias.max {
		first := q.topicAlias.alias.Front()
		elem := first.Value.(*aliasElem)
		q.topicAlias.alias.Remove(first)
		delete(q.topicAlias.index, elem.topic)
		alias = elem.alias
	} else {
		alias = uint16(l + 1)
	}
	q.topicAlias.alias.PushBack(&aliasElem{
		topic: topicName,
		alias: alias,
	})
	q.topicAlias.index[topicName] = alias
	return
}
