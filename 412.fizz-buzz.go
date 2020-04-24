/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	re := []string{}

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			re = append(re, "FizzBuzz")
		} else if i%5 == 0 {
			re = append(re, "Buzz")
		} else if i%3 == 0 {
			re = append(re, "Fizz")
		} else {
			re = append(re, strconv.Itoa(i))
		}
	}

	return re
}
// @lc code=end

