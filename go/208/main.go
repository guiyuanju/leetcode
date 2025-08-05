package main

func main() {
}

type Trie struct {
	val      byte
	children map[byte]*Trie
	isWord   bool
}

func Constructor() Trie {
	return Trie{0, map[byte]*Trie{}, false}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i, c := range word {
		if _, ok := cur.children[byte(c)]; !ok {
			node := Constructor()
			node.val = byte(c)
			cur.children[byte(c)] = &node
		}
		cur = cur.children[byte(c)]
		if i == len(word)-1 {
			cur.isWord = true
		}
	}
}

func findNode(trie *Trie, word string) *Trie {
	cur := trie
	for _, c := range []byte(word) {
		if node, ok := cur.children[c]; ok {
			cur = node
		} else {
			return nil
		}
	}
	return cur
}

func (this *Trie) Search(word string) bool {
	node := findNode(this, word)
	if node == nil {
		return false
	}
	return node.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	return findNode(this, prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
