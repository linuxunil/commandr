package commandr

type BaseHeader struct {
	data map[string][]byte
}

func (h *BaseHeader) Del(k string) bool {
	if _, exists := h.data[k]; exists {
		delete(h.data, k)
		return true
	}
	return false
}

func NewHeader() *BaseHeader { return &BaseHeader{} }

func (h *BaseHeader) Set(k string, v []byte) {
	if h.data == nil {
		h.data = make(map[string][]byte)
	}
	h.data[k] = v
}

func (h *BaseHeader) Get(k string) ([]byte, error) {
	if val, exists := h.data[k]; exists {
		return val, nil
	} else {
		return nil, ErrNotFound
	}
}

func (h *BaseHeader) Has(k string) bool {
	_, exists := h.data[k]
	return exists
}
