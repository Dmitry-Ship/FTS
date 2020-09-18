package main

type index map[string][]int

func (idx index) add(docs []document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				// Don't add same ID twice.
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

func getIntersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	result := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			result = append(result, a[i])
			i++
			j++
		}
	}
	return result
}

func (idx index) search(text string) []int {
	var result []int
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			if result == nil {
				result = ids
			} else {
				result = getIntersection(result, ids)
			}
		} else {
			// Token doesn't exist.
			return nil
		}
	}
	return result
}
