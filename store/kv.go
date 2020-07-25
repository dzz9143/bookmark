package store

// KeyValue - representing the Get Result
type KeyValue [2]string

// Key - get key
func (k *KeyValue) Key() string {
	return (*k)[0]
}

// Value = get value
func (k *KeyValue) Value() string {
	return (*k)[1]
}
