package main

/*
 * @lc app=leetcode id=676 lang=golang
 *
 * [676] Implement Magic Dictionary
 */

// @lc code=start

// MagicDictionary struct
type MagicDictionary struct {
	table map[string][]string
}

// Constructor func
/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{}
}

// BuildDict func
func (d *MagicDictionary) BuildDict(dictionary []string) {
	table := map[string][]string{}

	for _, s := range dictionary {
		for i := 0; i < len(s); i++ {
			ar := []byte(s)
			ar[i] = '*'
			if _, ok := table[string(ar)]; ok {
				table[string(ar)] = append(table[string(ar)], s)
			} else {
				table[string(ar)] = []string{s}
			}

		}
	}

	d.table = table
}

// Search func
func (d *MagicDictionary) Search(searchWord string) bool {
	for i := 0; i < len(searchWord); i++ {
		ar := []byte(searchWord)
		ar[i] = '*'
		s := string(ar)
		if d.table[s] != nil {
			for _, text := range d.table[s] {
				if text != searchWord {
					return true
				}
			}
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
// @lc code=end
