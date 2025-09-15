package messagecenter

import "sync"

type Observer interface {
	Update(msg NotifyMessage)
}

type NotifyMessage struct {
	Type string `json:"type"`
	Data []byte `json:"data"`
}

type MessageCenter struct {
	observers map[Observer]struct{}
	mu        sync.RWMutex
}

func NewMessageCenter() *MessageCenter {
	return &MessageCenter{
		observers: make(map[Observer]struct{}),
	}
}

func (mc *MessageCenter) Register(obj Observer) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.observers[obj] = struct{}{}
}

func (mc *MessageCenter) Unregister(obj Observer) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	delete(mc.observers, obj)
}

func (mc *MessageCenter) Notify(msg NotifyMessage) {
	mc.mu.RLock()
	defer mc.mu.RUnlock()
	for o := range mc.observers {
		o.Update(msg)
	}
}
