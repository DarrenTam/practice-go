package hash_table

type HashTable struct {
	hashValue [10]HashValue
	size      int
}

func (h HashTable) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = hash*31 + i
	}

	return hash % h.size
}

func (h *HashTable) Add(key string, value interface{}) {
	index := h.hash(key)
	originIndex := h.size

	if h.hashValue[index].key != key {
		for h.hashValue[index].key != "" {
			if index != originIndex {
				if index != 0 {
					index = (index + 1) % h.size
				} else {
					index = index + 1
				}
			} else {
				break
			}
		}
	}

	if index == originIndex {
		panic("The table already full")
	}

	h.hashValue[index].value = value
	h.hashValue[index].key = key
}

func (h *HashTable) Get(key string) interface{} {
	index := h.hash(key)
	originIndex := h.size

	for h.hashValue[index].key != key {

		if h.hashValue[index].key == key {
			break
		}

		if index != originIndex-1 {
			if index != 0 {
				index = (index + 1) % h.size
			} else {
				index = index + 1
			}

		} else {
			break
		}
	}

	if index == originIndex {
		panic("Can't found key")
	}

	return h.hashValue[index].value
}

func (h *HashTable) Exist(key string) bool {
	exist := false
	index := h.hash(key)
	originIndex := h.size

	if h.hashValue[index].key == key {
		exist = true
		return exist
	}

	for h.hashValue[index].key != "" {

		if h.hashValue[index].key == key {
			exist = true
			break
		}

		if index != 0 {
			index = (index + 1) % h.size
		} else {
			index = index + 1
		}

		if index == originIndex {
			break
		}

	}

	return exist
}

func (h *HashTable) Remove(key string) {
	index := h.hash(key)
	originIndex := h.size

	for h.hashValue[index].key != "" {

		if h.hashValue[index].key == key {
			h.hashValue[index].key = ""
			break
		}

		if index != 0 {
			index = (index + 1) % h.size
		} else {
			index = index + 1
		}

		if index == originIndex {
			break
		}

	}

}
