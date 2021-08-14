package main

import (
	"Algorithms/hashTable/hashMap"
	"fmt"
)

/*思路
1、通过hash 求key的列表index
2、通过index找到指定的链表
3、遍历链表找到和k相等的节点，返回 v

1、创建长度为16的列表
2、
*/

func main() {
	hashTable := hashMap.HashTable{}
	hashTable.Add("Name", "xlc")
	hashTable.Add("Gender", "man")
	hashTable.Add("Age", "26")
	if v, ok := hashTable.Search("Age");ok{
		fmt.Println(v)
	}

}
