package store

// Data - simple k v map
type Data map[string]string

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

// NewData - get new Data
func NewData() Data {
	return Data{}
}

// Add - Add a k:v pair
func (d *Data) Add(k, v string) {
	(*d)[k] = v
}

// Get - get all kv pairs given an array of keys
func (d *Data) Get(keys []string) []KeyValue {
	var s []KeyValue
	for _, k := range keys {
		v, ok := (*d)[k]
		if ok {
			s = append(s, KeyValue{k, v})
		}
	}

	return s
}
