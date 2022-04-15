package main

import "fmt"

func main() {
	// 1. Map是一种Key-Value的数据结构，内部存储的二元数据对
	//    无序的，通过Key来索引，引用类型

	// 2. Map的创建方法：
	// 方法一：直接量创建
	m := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
	}

	fmt.Println(m)

	// 方法二：make函数 empty map
	m2 := make(map[string]string)

	fmt.Println(m2, m2 == nil)

	// 方法三，定义不初始化，nil
	var m3 map[string]string
	fmt.Println(m3, m3 == nil)

	// 3. Map的遍历
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 4. 访问某个value
	fmt.Println(m["key1"])

	fmt.Println(m["key5"])
	// 如何知道key不存在？
	// value, ok := m["key1"]
	// fmt.Println(value, ok)

	// value, ok = m["key5"]
	// fmt.Println(value, ok)
	if value, ok := m["key5"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("wrong key")
	}

	// 5. 删除元素
	delete(m, "key1")
	if value, ok := m["key1"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("wrong key")
	}

	// 6. 获得Map中二元对的个数
	fmt.Println(len(m))

	// 7. 哪些类型能做Map的key
	// key 必须能够用 == 比较是否相等
	// 除了slice，map，function以外的所有内建类型

}
