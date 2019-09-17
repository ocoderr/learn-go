package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str = "中文中文 This is an example of a string"
	fmt.Printf("%t\n", strings.HasPrefix(str, "中文"))
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	fmt.Printf("%t\n", strings.Contains(str, "中文"))
	fmt.Printf("%t\n", strings.Contains(str, "Th"))

	fmt.Printf("%d\n", strings.Index(str, "文"))           // 不支持包含中文
	fmt.Printf("%d\n", strings.Index(str, "T"))           // 不支持包含中文
	fmt.Printf("%d\n", strings.IndexRune(str, rune('文'))) // 不支持包含中文
	fmt.Printf("%d\n", IndexRune(str, '中'))               //
	fmt.Printf("%d\n", IndexRune(str, '文'))               //
	fmt.Printf("%d\n", IndexRune(str, 'e'))               //

	fmt.Printf("%s\n", strings.Replace(str, "中文", "", -1))
	fmt.Printf("%s\n", strings.Replace(str, "example", "🌰", -1))

	fmt.Printf("%d\n", strings.Count(str, "is"))
	fmt.Printf("%d\n", strings.Count(str, "中"))
	fmt.Printf("%d\n", strings.Count(str, "中文"))

	fmt.Printf("%s\n", strings.Repeat("s", 10))
	fmt.Printf("%s\n", strings.Repeat("中", 10))

	fmt.Printf("%s\n", strings.ToLower("A"))

	fmt.Printf("%s\n", strings.TrimSpace("  A   "))
	fmt.Printf("%s\n", strings.Trim("[A]", "[|]"))

	split := strings.Split("a,b,c,d", ",")
	for _, s := range split {
		fmt.Printf("%s\n", s)
	}

	fmt.Printf("%s\n", strings.Join(split, "|"))

	fmt.Printf("strconv.Itoa %s\n", strconv.Itoa(100))
	atoi, _ := strconv.Atoi("100")
	fmt.Printf("strconv.Atoi %d\n", atoi)

	// 用 reader 方式 读取 string
	reader := strings.NewReader(str)
	bytes := make([]byte, len(str))
	reader.Read(bytes)
	fmt.Printf("%s\n", string(bytes))

}

func IndexRune(str string, s rune) int {
	for i, c := range []rune(str) {
		if c == s {
			return i
		}
	}
	return -1
}
