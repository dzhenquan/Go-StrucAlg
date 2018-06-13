package slist

import (
	"testing"
	"fmt"
	"strings"
)

type STest struct {
	num 	int
	age 	int
	name 	string
}

func TestLinkedSList_Insert(t *testing.T) {
	fmt.Println("\n-------------- test insert ---------------")

	slist1 := LinkedSList{}

	slist1.InsertToFirst("xiaodai")
	slist1.InsertToFirst("xiaocao")
	slist1.InsertToFirst("xiaohua")
	slist1.Insert("haotian", 1)
	slist1.Insert("wuchen", 2)

	for i := 0; i < slist1.Size(); i++ {
		fmt.Printf("i:[%d], value:[%s]\n", i, slist1.Find(i))
	}
	fmt.Printf("head:[%s], tail:[%s]\n", slist1.FindFirst(), slist1.FindLast())
}

func TestLinkedSList_InsertToFirst(t *testing.T) {

	fmt.Println("\n-------------- test insert to first ---------------")

	slist1 := LinkedSList{}

	slist1.InsertToFirst("xiaodai")
	slist1.InsertToFirst("xiaocao")
	slist1.InsertToFirst("xiaohua")
	slist1.InsertToFirst("wuchen")
	slist1.InsertToFirst("haotian")

	for i := 0; i < slist1.Size(); i++ {
		fmt.Printf("i:[%d], value:[%s]\n", i, slist1.Find(i))
	}
	fmt.Printf("head:[%s], tail:[%s]\n", slist1.FindFirst(), slist1.FindLast())
}

func TestLinkedSList_InsertToLast(t *testing.T) {
	fmt.Println("\n-------------- test insert to last ---------------")

	slist1 := LinkedSList{}

	slist1.InsertToLast("xiaodai")
	slist1.InsertToLast("xiaocao")
	slist1.InsertToLast("xiaohua")
	slist1.InsertToLast("wuchen")
	slist1.InsertToLast("haotian")

	for i := 0; i < slist1.Size(); i++ {
		fmt.Printf("i:[%d], value:[%s]\n", i, slist1.Find(i))
	}
	fmt.Printf("head:[%s], tail:[%s]\n", slist1.FindFirst(), slist1.FindLast())
}

func TestLinkedSList_IndexOf(t *testing.T) {
	fmt.Println("\n-------------- test index of ---------------")

	stest1 := &STest{3, 5, "xiaodai"}
	stest2 := &STest{2, 4, "xiaocao"}
	stest3 := &STest{1, 6, "xiaohua"}
	stest4 := &STest{7, 9, "wuchen"}
	stest5 := &STest{8, 11, "haotian"}

	slist1 := LinkedSList{}

	slist1.InsertToLast(stest1)
	slist1.InsertToLast(stest2)
	slist1.InsertToLast(stest3)
	slist1.InsertToLast(stest4)
	slist1.InsertToLast(stest5)

	index1 := slist1.IndexOf(stest2)
	fmt.Printf("index1:[%d]\n", index1)

	index2 := slist1.IndexOf(stest4)
	fmt.Printf("index2:[%d]\n", index2)
}

func TestLinkedSList_Reverse(t *testing.T) {
	fmt.Println("\n-------------- test reverse ---------------")

	slist1 := LinkedSList{}

	slist1.InsertToLast("xiaodai")
	slist1.InsertToLast("xiaocao")
	slist1.InsertToLast("xiaohua")
	slist1.InsertToLast("wuchen")
	slist1.InsertToLast("haotian")

	fmt.Println("\n-------------- reverse start ---------------")
	for i := 0; i < slist1.Size(); i++ {
		fmt.Printf("i:[%d], value:[%s]\n", i, slist1.Find(i))
	}
	fmt.Printf("head:[%s], tail:[%s]\n", slist1.FindFirst(), slist1.FindLast())


	fmt.Println("\n-------------- reverse end ---------------")

	slist1.Reverse()
	for i := 0; i < slist1.Size(); i++ {
		fmt.Printf("i:[%d], value:[%s]\n", i, slist1.Find(i))
	}
	fmt.Printf("head:[%s], tail:[%s]\n", slist1.FindFirst(), slist1.FindLast())
}

func TestLinkedSList_FindCustom(t *testing.T) {
	fmt.Println("\n-------------- test find custom ---------------")

	stest1 := &STest{3, 5, "xiaodai"}
	stest2 := &STest{2, 4, "xiaocao"}
	stest3 := &STest{1, 6, "xiaohua"}
	stest4 := &STest{7, 9, "wuchen"}
	stest5 := &STest{8, 11, "haotian"}
	stest6 := STest{9, 12, "xiaoming"}

	slist1 := LinkedSList{}

	slist1.InsertToLast(stest1)
	slist1.InsertToLast(stest2)
	slist1.InsertToLast(stest3)
	slist1.InsertToLast(stest4)
	slist1.InsertToLast(stest5)

	if value, ok := slist1.FindCustom(findCustomName, "xiaohua"); ok {
		fmt.Printf("num:[%d],age:[%d],name:[%s]\n",
			value.(*STest).num, value.(*STest).age, value.(*STest).name)
	} else {
		fmt.Println("Not exists!")
	}

	if value, ok := slist1.FindCustom(findCustomName, "xiaoming"); ok {
		fmt.Printf("num:[%d],age:[%d],name:[%s]\n",
			value.(*STest).num, value.(*STest).age, value.(*STest).name)
	} else {
		fmt.Println("Not [xiaoming]exists!")
	}

	if value, ok := slist1.FindCustom(findCustom, stest2); ok {
		fmt.Printf("num:[%d],age:[%d],name:[%s]\n",
			value.(*STest).num, value.(*STest).age, value.(*STest).name)
	} else {
		fmt.Println("Not exists!")
	}

	if value, ok := slist1.FindCustom(findCustom, stest6); ok {
		fmt.Printf("num:[%d],age:[%d],name:[%s]\n",
			value.(STest).num, value.(STest).age, value.(STest).name)
	} else {
		fmt.Println("Not [stest6] exists!")
	}
}

func findCustom(value1 interface{}, value2 interface{}) (interface{}, bool) {
	if value1 == value2 {
		return value1, true
	}
	return nil, false
}

func findCustomName(value1 interface{}, value2 interface{}) (interface{}, bool) {
	value := value1.(*STest).name

	if strings.Compare(value, value2.(string)) == 0 {
		return value1, true
	}

	return nil, false
}

func TestLinkedSList_Remove(t *testing.T) {
	fmt.Println("\n-------------- test remove ---------------")

	slist1 := LinkedSList{}

	slist1.InsertToLast("xiaodai")
	slist1.InsertToLast("xiaocao")
	slist1.InsertToLast("xiaohua")

	slist1.Remove(0)
	slist1.Remove(1)

	for i := 0; i < slist1.Size(); i++ {
		fmt.Printf("i:[%d], value:[%s]\n", i, slist1.Find(i))
	}
}

func TestLinkedSList_RemoveValue(t *testing.T) {
	fmt.Println("\n-------------- test remove value ---------------")

	stest1 := &STest{3, 5, "xiaodai"}
	stest2 := &STest{2, 4, "xiaocao"}
	stest3 := &STest{1, 6, "xiaohua"}
	stest4 := &STest{7, 9, "wuchen"}
	stest5 := &STest{8, 11, "haotian"}

	slist1 := LinkedSList{}

	slist1.InsertToLast(stest1)
	slist1.InsertToLast(stest2)
	slist1.InsertToLast(stest3)
	slist1.InsertToLast(stest4)
	slist1.InsertToLast(stest5)

	index1 := slist1.IndexOf(stest3)
	fmt.Printf("index1:[%d], slist_len:[%d]!\n", index1, slist1.Size())

	slist1.RemoveValue(stest3)

	index2 := slist1.IndexOf(stest3)
	fmt.Printf("index2:[%d], slist_len:[%d]!\n", index2, slist1.Size())
}

func TestLinkedSList_RemoveFirst(t *testing.T) {
	fmt.Println("\n-------------- test remove to first ---------------")

	slist1 := LinkedSList{}

	slist1.InsertToLast("xiaodai")
	slist1.InsertToLast("xiaocao")
	slist1.InsertToLast("xiaohua")

	slist1.RemoveFirst()
	slist1.RemoveFirst()
	slist1.RemoveFirst()

	for i := 0; i < slist1.Size(); i++ {
		fmt.Printf("i:[%d], value:[%s]\n", i, slist1.Find(i))
	}
}

func TestLinkedSList_RemoveLast(t *testing.T) {
	fmt.Println("\n-------------- test remove to last ---------------")

	slist1 := LinkedSList{}

	slist1.InsertToLast("xiaodai")
	slist1.InsertToLast("xiaocao")
	slist1.InsertToLast("xiaohua")

	slist1.RemoveLast()
	slist1.RemoveLast()
	slist1.RemoveLast()

	for i := 0; i < slist1.Size(); i++ {
		fmt.Printf("i:[%d], value:[%s]\n", i, slist1.Find(i))
	}
}

func TestLinkedSList_RemoveAll(t *testing.T) {
	fmt.Println("\n-------------- test remove all ---------------")

	slist1 := LinkedSList{}

	slist1.InsertToLast("xiaodai")
	slist1.InsertToLast("xiaocao")
	slist1.InsertToLast("xiaohua")

	slist1.RemoveAll()

	for i := 0; i < slist1.Size(); i++ {
		fmt.Printf("i:[%d], value:[%s]\n", i, slist1.Find(i))
	}
}