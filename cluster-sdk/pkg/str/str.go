package str

import (
	"strings"
)

// CompareIgnoreCase 比较两个字符串是否一样
// @param str1: 字符串1
// @param str2: 字符串2
// @return bool: 两个字符串是否一致
func CompareIgnoreCase(str1, str2 string) bool {

	// 去除空格
	str1 = strings.TrimSpace(str1)
	str2 = strings.TrimSpace(str2)

	// 去除str中间的空格
	str1 = strings.Join(strings.Fields(str1), "")
	str2 = strings.Join(strings.Fields(str2), "")

	// 全部转小写
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	// 比较字符串
	return str1 == str2
}
