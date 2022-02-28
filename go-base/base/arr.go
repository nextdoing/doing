package base

import (
	"fmt"
	"strings"
)

// Index 获取字符串数组的下标
func Index(vs []string, t string) int {
	for index, val := range vs {
		if val == t {
			return index
		}
	}
	return -1
}

// Include 判断某字符串是否在数组中
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any 数组中至少项满足规则f的要求返回true, 否则false
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All 当所有的数组项满足规则f的要求时返回true，否则false
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Map 数组中的每一项都经过规则f的处理，然后返回处理后的新数组
func Map(vs []string, f func(string) string) []string {
	res := make([]string, len(vs))
	for index, v := range vs {
		res[index] = f(v)
	}
	return res
}

// Filter 数组中每一项经过规则f的过滤，返回符合规则要求的新数组
func Filter(vs []string, f func(string) bool) []string {
	res := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

// Str 字符串数组
type Str []string

// Len 获取字符串数组的长度
func (s Str) Len() int {
	return len(s)
}

// Less 字符串数组相邻两项的比较逻辑
func (s Str) Less(i, j int) bool {
	return len(s[i]) <= len(s[j])
}

// Swap 字符串数组比较结果后的行为
func (s Str) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 包的默认执行函数
func init() {
	p := fmt.Println

	p("包含: test包含es？ ", strings.Contains("test", "es"))
	p("计数: test里有几个t？", strings.Count("test", "t"))
	p("字符以...开头: test以te开头？", strings.HasPrefix("test", "te"))
	p("字符以...结尾: test以st结尾？", strings.HasSuffix("test", "st"))
	p("...在字符串内的位置: st在test内开始的位置下标？", strings.Index("test", "st"))
	p("将字符串数组连接成字符串: [a 2 c] => a-2-c", strings.Join([]string{"a", "2", "c"}, "-"))
	p("重复: 将ts重复5次=>tststststs", strings.Repeat("ts", 5))
	p("全部替换: 将tester内的e替换成s=>tsstsr", strings.Replace("tester", "e", "s", -1))
	p("替换部分: 将tester内的e替换成s=>tsster", strings.Replace("tester", "e", "s", 1))
	p("分割为数组: 将s-p-l-i-t分割成[s p l i t]", strings.Split("s-p-l-i-t", "-"))
	p("转小写: Test => test", strings.ToLower("Test"))
	p("转大写: Test => TEST", strings.ToUpper("Test"))
	p("字符串长度: len('test')", len("test"))
	p("用下标取某个字符: 返回ASCII十进制码", "test"[1])
}
