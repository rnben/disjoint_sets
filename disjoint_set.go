package graph

type DisjointSet[T comparable] struct {
	elementMap map[T]element[T]
	fatherMap  map[element[T]]element[T]
	sizeMap    map[element[T]]int
}

type element[T comparable] struct {
	value T
}

func NewDisjointSet[T comparable](list []T) *DisjointSet[T] {
	elementMap := make(map[T]element[T])
	fatherMap := make(map[element[T]]element[T])
	sizeMap := make(map[element[T]]int)

	for _, value := range list {
		elem := element[T]{value}
		elementMap[value] = elem
		fatherMap[elem] = elem
		sizeMap[elem] = 1
	}

	return &DisjointSet[T]{
		elementMap: elementMap,
		fatherMap:  fatherMap,
		sizeMap:    sizeMap,
	}
}

func (unset *DisjointSet[T]) FindHead(value T) T {
	var stack []element[T]

	ele := element[T]{value}
	for ele != unset.fatherMap[ele] {
		stack = append(stack, ele)
		ele = unset.fatherMap[ele]
	}

	for len(stack) > 0 {
		elem := stack[0]
		stack = stack[1:]
		unset.fatherMap[elem] = ele
	}

	return ele.value
}

func (unset *DisjointSet[T]) IsSameSet(a, b T) bool {
	if _, ok := unset.elementMap[a]; !ok {
		return false
	}

	if _, ok := unset.elementMap[b]; !ok {
		return false
	}

	return unset.FindHead(a) == unset.FindHead(b)
}

func (unset *DisjointSet[T]) Union(a, b T) {
	aF := element[T]{unset.FindHead(a)}
	bF := element[T]{unset.FindHead(b)}

	if aF == bF {
		return
	}

	big, small := aF, bF
	if unset.sizeMap[small] > unset.sizeMap[big] {
		small, big = big, small
	}

	unset.fatherMap[small] = big
	unset.sizeMap[big] = unset.sizeMap[small] + unset.sizeMap[big]
	delete(unset.sizeMap, small)
}
