/*******
* @Author:qingmeng
* @Description:
* @File:check
* @Date:2022/7/30
 */

package tool

// CheckIfSensitive 检查敏感词和sql关键符号，防止sql注入，存在返回true
func CheckIfSensitive(s string) bool {
	return defaultTrie.CheckWords(s)
}
