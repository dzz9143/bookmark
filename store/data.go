package store

// Data - simple k v map
type Data map[string]string

// NewData - get new Data
func NewData() Data {
	return Data{}
}

// Add - Add a k:v pair
func (d Data) Add(k, v string) {
	d[k] = v
}

// Get - get all kv pairs given an array of keys
func (d Data) Get(keys []string) []KeyValue {
	s := make([]KeyValue, 0, len(keys))

	for _, k := range keys {
		v, ok := d[k]
		if ok {
			s = append(s, KeyValue{k, v})
		}
	}

	return s
}
