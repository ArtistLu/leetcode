/*
 * @lc app=leetcode.cn id=820 lang=golang
 *
 * [820] 单词的压缩编码
 *
 * https://leetcode-cn.com/problems/short-encoding-of-words/description/
 *
 * algorithms
 * Medium (41.38%)
 * Likes:    82
 * Dislikes: 0
 * Total Accepted:    16.5K
 * Total Submissions: 37.5K
 * Testcase Example:  '["time", "me", "bell"]'
 *
 * 给定一个单词列表，我们将这个列表编码成一个索引字符串 S 与一个索引列表 A。
 * 
 * 例如，如果这个列表是 ["time", "me", "bell"]，我们就可以将其表示为 S = "time#bell#" 和 indexes =
 * [0, 2, 5]。
 * 
 * 对于每一个索引，我们可以通过从字符串 S 中索引的位置开始读取字符串，直到 "#" 结束，来恢复我们之前的单词列表。
 * 
 * 那么成功对给定单词列表进行编码的最小字符串长度是多少呢？
 * 
 * 
 * 
 * 示例：
 * 
 * 输入: words = ["time", "me", "bell"]
 * 输出: 10
 * 说明: S = "time#bell#" ， indexes = [0, 2, 5] 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= words.length <= 2000
 * 1 <= words[i].length <= 7
 * 每个单词都是小写字母 。
 * 
 * 
 */

// @lc code=start

func minimumLengthEncoding(words []string) int {
	//单词反转
	words = reveerseStringArray(words)

	//反转后排序
	sort.Strings(words)

	words = append(words, "*")

	//前后单词对比如果后一个是前一个前缀略过
	l := 0
	for i := len(words) - 2; i >= 0; i-- {
		if strings.HasPrefix(words[i+1], words[i]) {
			continue
		}
		l += len(words[i]) + 1
	}

	return l
}

func reveerseStringArray(words []string) []string {
	for i, w := range words {
		words[i] = reverse(w)
	}

	return words
}

func reverse(word string) string {
	bytes := []rune(word)
	s := 0
	l := len(bytes) - 1
	for s < l {
		bytes[s], bytes[l] = bytes[l], bytes[s]
		s++
		l--
	}

	return string(bytes)
}
