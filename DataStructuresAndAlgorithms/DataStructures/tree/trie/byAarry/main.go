/*
 * File: main.go
 * Created Date: 2022-01-07 10:10:15
 * Author: ysj
 * Description:  前缀树、字典树 - 数组实现，用于英文单词匹配
 */
package main

import "fmt"

type Trier interface {
	Insert(word string)                    // 插入一个值
	Search(word string) (int, bool)        // 查找值，看是否存在，如存在，返回添加了几次
	StartsWith(prefix string) (int, bool)   // 前缀匹配是否存在，如存在，返回存在多少个
	AllWords() []string                     // 获取所有的词
	StartsWithWords(prefix string) []string // 获取所有以prefix开头的词
	Delete(word string)                    // 删除一个值
}

type TrieNode struct {
	pass  int
	end   int
	Value string
	Nexts [26]*TrieNode
}

// 初始化
func NewTrieNode() *TrieNode {
	return &TrieNode{
		pass:  0,
		end:   0,
		Value: "",
		Nexts: [26]*TrieNode{},
	}
}

// 插入一个值
func (t *TrieNode) Insert(word string) {
	for i, v := range word {
		j := int(v - 'a')
		if t.Nexts[j] == nil { // 不存在，添加node
			trieNode := NewTrieNode()
			trieNode.pass++
			trieNode.Value = string(v)
			t.Nexts[j] = trieNode

			if len(word) == 1 { // 结束字符，end+1
				trieNode.end++
			} else {
				trieNode.Insert(word[i+1:])
			}
			break
		}
		// 存在node，pass+1，并判断是否结束
		next := t.Nexts[j]
		next.pass++
		if len(word) == 1 { // 存在且为结束字符
			next.end++
		} else {
			next.Insert(word[i+1:])
			break
		}
	}
}

// 查找一个值
func (t *TrieNode) Search(word string) (int, bool) {
	for i, v := range word {
		j := int(v - 'a')
		next := t.Nexts[j]
		// 越界了，不存在
		if next == nil {
			return 0, false
		}

		// 最后一个字符
		if len(word) == 1 {
			if next.end > 0 { // 如果为结束点
				return next.end, true
			}
			return 0, false
		} else {
			return next.Search(word[i+1:])
		}
	}
	return 0, false
}

// 前缀匹配一个值
func (t *TrieNode) StartsWith(prefix string) (int, bool) {
	for i, v := range prefix {
		j := int(v - 'a')
		next := t.Nexts[j]
		// 不存在
		if next == nil {
			return 0, false
		}
		// 存在
		if len(prefix) == 1 { // 最后一个字符
			return next.pass, true
		} else {
			return next.StartsWith(prefix[i+1:])
		}
	}
	return 0, false
}

// 获取所有词
func (t *TrieNode) AllWords() []string {
	words := []string{}
	for _, node := range t.Nexts {
		if node == nil {
			continue
		}
		nodeWords := node.AllWords()
		if len(nodeWords) > 0 { // 非末端节点
			for _, v := range nodeWords {
				//当前值+子节点值
				words = append(words, node.Value+v)
				// 如果是结束字符！！！
				if node.end > 0 {
					words = append(words, node.Value)
				}
			}
		} else {
			//末端节点pass有多少就要加几次！！！
			for i := 0; i < node.pass; i++ {
				words = append(words, node.Value)
			}
		}
	}
	return words
}

// 获取所有以prefix开头的词
func (t *TrieNode) StartsWithWords(prefix string) []string {
	words := []string{}
	for i, v := range prefix {
		j := int(v - 'a')
		next := t.Nexts[j]
		if next == nil { // 不存在
			return words
		}
		// 存在
		if len(prefix) == 1 { // 结束字符
			nodeWords := next.AllWords()
			for _, nw := range nodeWords {
				words = append(words, prefix+nw)
			}
			return words
		} else { // 非结束字符, 继续判断前缀
			nextWords := next.StartsWithWords(prefix[i+1:])
			for _, nw := range nextWords {
				words = append(words, prefix+nw)
			}
			return words
		}
	}

	return words
}

// 删除一个值
func (t *TrieNode) Delete(word string) {
	// 必须先判断是否存在
	_, exists := t.Search(word)
	if !exists {
		return
	}
	for i, v := range word {
		j := int(v - 'a')
		next := t.Nexts[j]
		if next == nil { // 不存在
			return
		}
		// 存在则pass--
		next.pass--
		// 如果pass=0，说明已经没有经过该字符的词存在了
		if next.pass == 0 {
			// next = nil // 这行是错误的，这是一个新的指针
			t.Nexts[j] = nil
			return
		}
		// 如果pass>0
		if len(word) == 1 { // 删除结尾字符
			next.end--
			return
		} else {
			next.Delete(word[i+1:])
			break
		}
	}
}

func main() {
	words := []string{"apple", "appl",
		"banana", "bana",
		"chestnut", "chestnut",
		"strawberry", "strawberry",
		"peach", "pomegranate", "pomelo",
		"mango", "mangosteen", "mandarin",
		"walnut",
	}
	trie := NewTrieNode()
	for _, w := range words {
		trie.Insert(w)
	}

	fmt.Println("所有的词:", trie.AllWords())

	// 查找
	var (
		n      int
		exists bool
	)

	n, exists = trie.Search("apple")
	fmt.Println("apple:", n, exists)
	n, exists = trie.Search("walnut")
	fmt.Println("walnut:", n, exists)
	n, exists = trie.Search("fruits")
	fmt.Println("fruits:", n, exists)

	n, exists = trie.StartsWith("p")
	fmt.Println("p开头:", n, exists)
	n, exists = trie.StartsWith("po")
	fmt.Println("po开头:", n, exists)
	n, exists = trie.StartsWith("pre")
	fmt.Println("pre开头:", n, exists)
	n, exists = trie.StartsWith("a")
	fmt.Println("a开头:", n, exists)
	n, exists = trie.StartsWith("m")
	fmt.Println("m开头:", n, exists)

	trie.Delete("pomelo")
	n, exists = trie.Search("pomelo")
	fmt.Println("删除pomelo后:", n, exists)
	n, exists = trie.StartsWith("po")
	fmt.Println("po开头:", n, exists)

	fmt.Println("a开头的词有:", trie.StartsWithWords("a"))
}
