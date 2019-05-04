package leetcode

import (
	"strings"
)

func openLock(deadends []string, target string) int {
	var deadends_map = map[string]bool{}
	for _, deadend := range deadends {
		deadends_map[deadend] = true
	}
	root_node := ""
	for i := 0; i < len(target); i++ {
		root_node = root_node + "0"
	}
	if strings.EqualFold(root_node, target) {
		return 0
	} else if deadends_map[root_node] {
		return -1
	} else {
		return bfs(root_node, deadends_map, target)
	}
}

func bfs(node string, deadends map[string]bool, target string) int {
	var used = map[string]bool{}
	var cur_nodes = []string{}
	var next_nodes = []string{}
	var n int = 1
	var found bool = false
	used[node] = true
	cur_nodes = getChildNodes(node, used, deadends, []string{})
	for {
		next_nodes = []string{}
		for _, cnode := range cur_nodes {
			if strings.EqualFold(target, cnode) {
				found = true
				break
			}
			next_nodes = getChildNodes(cnode, used, deadends, next_nodes)
		}
		if found {
			break
		}
		n = n + 1
		cur_nodes = next_nodes
		if len(cur_nodes) == 0 {
			break
		}
	}
	if found {
		return n
	}
	return -1
}

func getChildNodes(node string, used map[string]bool, deadends map[string]bool, nodes []string) []string {
	for i, c := range node {
		next_c := getNextRune(c)
		front_c := getFrontRune(c)
		next_node := replaceStringIndex(node, i, next_c)
		front_node := replaceStringIndex(node, i, front_c)
		if !used[next_node] && !deadends[next_node] {
			used[next_node] = true
			nodes = append(nodes, next_node)
		}
		if !used[front_node] && !deadends[front_node] {
			used[front_node] = true
			nodes = append(nodes, front_node)
		}
	}
	return nodes
}

func getNextRune(c rune) rune {
	if c == '9' {
		return '0'
	}
	return c + 1
}

func getFrontRune(c rune) rune {
	if c == '0' {
		return '9'
	}
	return c - 1
}

func replaceStringIndex(str string, index int, val rune) string {
	return strings.Join([]string{str[:index], string(val), str[index+1:]}, "")
}
