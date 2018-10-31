package container

import "strings"

type HashSet struct {
	acceptEmpty bool
	set map[string]bool
}

func NewHashSet(acceptEmpty bool) *HashSet {
	return &HashSet{acceptEmpty, make(map[string]bool) }
}

func (hashset *HashSet) Add(input string) bool {
	if !hashset.acceptEmpty && len(strings.TrimSpace(input)) == 0 {
		return false;
	}

	_, found := hashset.set[input]
	hashset.set[input] = true
	return !found //False if it existed already
}

func (hashset *HashSet) IsExist(input string) bool {
	_, found := hashset.set[input]
	return found //true if it existed already
}

func (hashset *HashSet) Remove(i string) {
	delete(hashset.set, i)
}

func (hashset *HashSet) AsArray() []string {
	v := make([]string, 0, len(hashset.set))

	for key, _ := range hashset.set {
		v = append(v, key)
	}

	return v;
}