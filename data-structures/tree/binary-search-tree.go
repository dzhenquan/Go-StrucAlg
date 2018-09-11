package main

import "fmt"

type BinarySearchTree struct {
	value 		int
	lchild 		*BinarySearchTree
	rchild 		*BinarySearchTree
}

func main() {
	var tree *BinarySearchTree

	nums := []int{11, 7, 5, 9, 100, 90, 150, 130}

	for _, value := range nums {
		tree = Add(tree, value)
	}

	//fmt.Println(tree.LeafNum())
	PreOrderTraverse(tree)
	fmt.Println()

	fmt.Println(PreOrderTraverse1(tree))

	InOrderInverse(tree)
	fmt.Println()

	fmt.Println(InOrderInverse1(tree))

	PostOrderTraverse(tree)
	fmt.Println()

	fmt.Println(PostOrderTraverse1(tree))

	//fmt.Println(LeafNum(tree))

	tree = InvertTree(tree)
	fmt.Println(InOrderInverse1(tree))
}

// 先序递归遍历
func PreOrderTraverse(tree *BinarySearchTree) {
	if tree != nil {
		fmt.Printf("%d ", tree.value)
		PreOrderTraverse(tree.lchild)
		PreOrderTraverse(tree.rchild)
	}
}

// 先序非递归遍历
func PreOrderTraverse1(tree *BinarySearchTree) []int {
	var stack [100]		*BinarySearchTree
	var target_node 	*BinarySearchTree
	var value_list 		[]int

	target_node = tree

	top := 0
	for target_node != nil || top > 0 {
		for target_node != nil {
			value_list = append(value_list, target_node.value)
			stack[top] = target_node
			top++
			target_node = target_node.lchild
		}
		if (top > 0) {
			top--
			target_node = stack[top]
			target_node = target_node.rchild
		}
	}
	target_node = nil
	return value_list
}

// 中序递归遍历
func InOrderInverse(tree *BinarySearchTree) {
	if tree != nil {
		InOrderInverse(tree.lchild)
		fmt.Printf("%d ", tree.value)
		InOrderInverse(tree.rchild)
	}
}

// 中序非递归遍历
func InOrderInverse1(tree *BinarySearchTree) []int {
	var stack [100]		*BinarySearchTree
	var terget_node 	*BinarySearchTree
	var value_list 		[]int

	terget_node = tree

	top := 0
	for terget_node != nil || top > 0 {
		for terget_node != nil {
			stack[top] = terget_node
			top++
			terget_node = terget_node.lchild
		}

		if top > 0 {
			top--
			terget_node = stack[top]
			value_list = append(value_list, terget_node.value)
			terget_node = terget_node.rchild
		}
	}
	terget_node = nil
	return value_list
}

// 后序递归遍历
func PostOrderTraverse(tree *BinarySearchTree) {
	if tree != nil {
		PostOrderTraverse(tree.lchild)
		PostOrderTraverse(tree.rchild)
		fmt.Printf("%d ", tree.value)
	}
}

// 后序非递归遍历
func PostOrderTraverse1(tree *BinarySearchTree) []int {
	var stack 			[100]*BinarySearchTree
	var target_node 	*BinarySearchTree
	var tmp_node   		*BinarySearchTree
	var value_list 		[]int

	target_node = tree
	tmp_node = nil

	top := 0
	for target_node != nil || top > 0 {
		for target_node != nil {
			stack[top] = target_node
			top++
			target_node = target_node.lchild
		}
		if top > 0 {
			target_node = stack[top - 1]
			if target_node.rchild == nil || target_node.rchild == tmp_node {
				value_list = append(value_list, target_node.value)
				tmp_node = target_node
				target_node = nil
				top--
			} else {
				target_node = target_node.rchild
			}
		}
	}
	tmp_node = nil
	target_node = nil

	return value_list
}

// 二叉树的反转
func InvertTree(tree *BinarySearchTree) *BinarySearchTree {
	if tree == nil {
		return nil
	}

	tree.lchild = InvertTree(tree.lchild)
	tree.rchild = InvertTree(tree.rchild)

	target_node := tree.lchild
	tree.lchild = tree.rchild
	tree.rchild = target_node;

	target_node = nil

	return tree
}

func Add(tree *BinarySearchTree, value int) *BinarySearchTree {
	if tree == nil {
		tree = new(BinarySearchTree)
		tree.value = value
	} else {
		if value < tree.value {
			tree.lchild = Add(tree.lchild, value)
		} else {
			tree.rchild = Add(tree.rchild, value)
		}
	}
	return tree
}

func LeafNum(tree *BinarySearchTree) int {
	if tree == nil {
		return 0
	}

	if tree.lchild == nil && tree.rchild == nil {
		return 1
	} else {
		return LeafNum(tree.lchild) + LeafNum(tree.rchild)
	}
}
