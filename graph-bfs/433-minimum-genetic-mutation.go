package graphbfs

import (
	"container/list"
	"strings"
)

func minMutation(startGene string, endGene string, bank []string) int {
	bankSet := make(map[string]bool, len(bank))

	for _, g := range bank {
		bankSet[g] = true
	}

	if _, ok := bankSet[endGene]; !ok {
		return -1
	}

	if startGene == endGene {
		return 1
	}

	genes := []rune{'A', 'C', 'G', 'T'}

	q := list.New()
	type Gen struct {
		g       string
		changes int
	}
	q.PushBack(Gen{startGene, 0})

	visited := make(map[string]bool)
	visited[startGene] = true

	for q.Len() > 0 {
		curr := q.Remove(q.Front()).(Gen)

		for i, ch := range curr.g {
			for _, g := range genes {
				if ch == g {
					continue
				}
				newGen := strings.Clone(curr.g)
				runes := []rune(newGen)
				runes[i] = g
				newGen = string(runes)

				if _, ok := visited[newGen]; ok {
					continue
				}
				if _, ok := bankSet[newGen]; !ok {
					continue
				}

				if strings.Compare(newGen, endGene) == 0 {
					return curr.changes + 1
				}

				q.PushBack(Gen{newGen, curr.changes + 1})
				visited[newGen] = true
			}
		}
	}

	return -1
}
