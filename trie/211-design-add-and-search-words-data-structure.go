package trie

type WordDictionary struct {
	Nodes map[rune]*WordDictionary
	IsEnd bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		Nodes: make(map[rune]*WordDictionary),
	}
}

func (this *WordDictionary) AddWord(word string) {
	p := this
	for _, ch := range word {
		if _, ok := p.Nodes[ch]; !ok {
			node := Constructor()
			p.Nodes[ch] = &node
		}

		p = p.Nodes[ch]
	}

	p.IsEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	p := this
	for i, ch := range word {
		if _, ok := p.Nodes[ch]; !ok {
			if ch != '.' {
				return false
			}
			for _, node := range p.Nodes {
				if node.Search(word[i+1:]) {
					return true
				}
			}
			return false
		}

		p = p.Nodes[ch]
	}

	return p.IsEnd
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
