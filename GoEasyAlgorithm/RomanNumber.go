package main

// Declare a map of char and number
var symbolValue = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var answer int
	var stringLength int = len(s)
	// If the right-hand side is bigger than itself then you take the opposite of yourself
	for i := 0; i < stringLength; i++ {
		// Anyway the latest number will add in answer and determine if it's less than the last digit
		value := symbolValue[s[i]]
		if i < stringLength - 1 && value < symbolValue[s[i + 1]] {
			// Take the symbols
			answer -= value
		}else {
			answer += value
		}
	}
	return answer
}

func romanToIntNo2(s string) int {
	/*
    创建一个Map存储映射关系
    循环拿到输入的字符串
    用映射的值判断：
    1、如果索引0的值>索引1的值--->求和
    2、如果索引0的值<索引1的值--->做差
    */
    m := map[rune]int{
        rune('I'): 1,
		rune('V'): 5,
		rune('X'): 10,
		rune('L'): 50,
		rune('C'): 100,
		rune('D'): 500,
		rune('M'): 1000,
    }
    //定义一个字符数组存储输入的字符串
    str := []rune(s)
    if len(str) == 1 {
        return m[str[0]]
    }
    answer := m[str[0]]
    for j := 1; j < len(str); j++ {
        if m[str[j-1]] < m[str[j]] {
            //重新赋值result，让他为前两位之差
            answer += m[str[j]] - 2*m[str[j-1]]
        }else {
            answer += m[str[j]]
        }
    }
	return answer
}