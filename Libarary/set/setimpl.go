package set

type HashSet struct {
	items map[interface{}]status
}

type status bool

const statusExists status = true

func New() *HashSet {
	return &HashSet{items: make(map[interface{}]status)}
}

func (set *HashSet) Add(item interface{}) bool {
	if _, exists := set.items[item]; exists {
		return false
	}
	set.items[item] = statusExists
	return true
}

func (set *HashSet) Remove(item interface{}) bool {
	if !set.IsElementOf(item) {
		return false
	}

	if _, exists := set.items[item]; !exists {
		return false
	}
	delete(set.items, item)
	return true
}

func (set *HashSet) Size() int {
	return len(set.items)
}

func (set *HashSet) IsElementOf(item interface{}) bool {
	if _, exists := set.items[item]; exists {
		return true
	} else {
		return false
	}
}

func (set *HashSet) Values() []interface{} {
	values := make([]interface{}, len(set.items))
	count := 0
	for item := range set.items {
		values = append(values, item)
		count++
	}
	return values

}
