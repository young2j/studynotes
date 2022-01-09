/*
 * File: main.go
 * Created Date: 2022-01-09 10:55:51
 * Author: ysj
 * Description:  前缀树、字典树 - map实现
 */

package main

import "fmt"

type Trier interface {
	Insert(word string)
	Search(word string) (int, bool)
	StartsWith(prefix string) (int, bool)
	AllWords() []string
	StartsWithWords(prefix string) []string
	Delete(word string)
}

type TrieNode struct {
	pass  int
	end   int
	Value rune
	Nexts map[rune]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		pass:  0,
		end:   0,
		Value: 0,
		Nexts: make(map[rune]*TrieNode),
	}
}

// 插入一个词
func (t *TrieNode) Insert(word string) {
	runeWord := []rune(word)
	for i, v := range runeWord {
		node, ok := t.Nexts[v]
		if !ok { // 不存在，需要插入节点
			trieNode := NewTrieNode()
			t.Nexts[v] = trieNode
			trieNode.Value = v
			trieNode.pass++
			if i == len(runeWord)-1 {
				trieNode.end++
				return
			}
			trieNode.Insert(string(runeWord[i+1:]))
			break
		}
		// 存在该节点
		node.pass++
		if i == len(runeWord)-1 {
			node.end++
			return
		} else {
			node.Insert(string(runeWord[i+1:]))
			break
		}
	}
}

// 查找一个词
func (t *TrieNode) Search(word string) (int, bool) {
	runeWord := []rune(word)
	for i, v := range runeWord {
		node, ok := t.Nexts[v]
		if !ok {
			return 0, false
		}
		if i == len(runeWord)-1 {
			if node.end > 0 {
				return node.pass, true
			}
			return 0, false
		} else {
			return node.Search(string(runeWord[i+1:]))
		}
	}
	return 0, false
}

// 前缀匹配prefix
func (t *TrieNode) StartsWith(prefix string) (int, bool) {
	runePrefix := []rune(prefix)
	for i, v := range runePrefix {
		node, ok := t.Nexts[v]
		if !ok {
			return 0, false
		}
		if i == len(runePrefix)-1 {
			return node.pass, true
		} else {
			return node.StartsWith(string(runePrefix[i+1:]))
		}
	}
	return 0, false
}

// 节点下的所有词
func (t *TrieNode) AllWords() []string {
	words := []string{}
	for k, node := range t.Nexts {
		value := string(k)
		nodeWords := node.AllWords()
		if len(nodeWords) == 0 { // 如果是尾节点
			for i := 0; i < node.pass; i++ {
				words = append(words, value)
			}
		} else {
			for _, nw := range nodeWords {
				words = append(words, value+nw)
				if node.end > 0 {
					words = append(words, value)
				}
			}
		}
	}
	return words
}

// 以prefix开头的所有词
func (t *TrieNode) StartsWithWords(prefix string) []string {
	runePrefix := []rune(prefix)
	words := []string{}
	for i, v := range runePrefix {
		node, ok := t.Nexts[v]
		if !ok {
			return words
		}
		if i == len(runePrefix)-1 {
			nodeWords := node.AllWords()
			for _, nw := range nodeWords {
				words = append(words, prefix+nw)
			}
			return words
		} else {
			nextsWords := node.StartsWithWords(string(runePrefix[i+1:]))
			for _, nw := range nextsWords {
				words = append(words, prefix+nw)
			}
			return words
		}
	}
	return words
}

// 删除一个词
func (t *TrieNode) Delete(word string) {
	// 先查找是否存在
	_, exists := t.Search(word)
	if !exists {
		return
	}
	runeWord := []rune(word)
	for i, v := range runeWord {
		node, ok := t.Nexts[v]
		if !ok {
			return
		}
		node.pass--
		if node.pass == 0 {
			// t.Nexts[v] = nil 这是错的
			delete(t.Nexts, v)
			return
		}
		if i == len(runeWord)-1 {
			node.end--
			return
		} else {
			node.Delete(string(runeWord[i+1:]))
			break
		}
	}
}

func main() {
	words := []string{"apple", "apple",
		"banana", "bana",
		"chestnut", "chest",
		"中国人", "不骗", "中国人",
		"美丽的人儿", "美好的火腿肠", "美好的",
		"随便", "再", "打几个词儿",
	}

	trie := NewTrieNode()
	for _, word := range words {
		trie.Insert(word)
	}
	fmt.Println("所有的词有:", trie.AllWords())

	var (
		n      int
		exists bool
	)

	n, exists = trie.Search("apple")
	fmt.Println("apple:", n, exists)

	n, exists = trie.Search("banana")
	fmt.Println("banana:", n, exists)

	n, exists = trie.Search("中国人")
	fmt.Println("中国人:", n, exists)

	n, exists = trie.StartsWith("c")
	fmt.Println("c开头的:", n, exists)
	fmt.Println("c开头的词:", trie.StartsWithWords("c"))

	n, exists = trie.StartsWith("美")
	fmt.Println("美开头的:", n, exists)
	fmt.Println("美开头的词:", trie.StartsWithWords("美"))

	trie.Delete("美好的")
	fmt.Println("---删除'美好的'---")
	n, exists = trie.StartsWith("美")
	fmt.Println("美开头的:", n, exists)
	fmt.Println("美开头的词:", trie.StartsWithWords("美"))
}
