package trie

const (
	partial_inserted = 1
	fully_inserted   = 2
)

type Trie struct {
	Dict map[string]int
}

// func Constructor() Trie {
// 	return Trie{Dict: make(map[string]int)}
// }

func (this *Trie) Insert(word string) {
	str := ""

	for _, ch := range word {
		str += string(ch)
		if _, ok := this.Dict[str]; !ok {
			this.Dict[str] = partial_inserted

		}
	}

	this.Dict[word] = fully_inserted
}

func (this *Trie) Search(word string) bool {
	v, ok := this.Dict[word]
	if !ok {
		return false
	}

	return v == fully_inserted
}

func (this *Trie) StartsWith(prefix string) bool {
	_, ok := this.Dict[prefix]
	return ok
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
