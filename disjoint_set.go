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

func (set *DisjointSet[T]) FindHead(value T) T {
	var stack []element[T]

	ele := element[T]{value}
	// TODO: complex type compare
	for ele != set.fatherMap[ele] {
		stack = append(stack, ele)
		ele = set.fatherMap[ele]
	}

	for len(stack) > 0 {
		elem := stack[0]
		stack = stack[1:]
		set.fatherMap[elem] = ele
	}

	return ele.value
}

func (set *DisjointSet[T]) IsSameSet(a, b T) bool {
	if _, ok := set.elementMap[a]; !ok {
		return false
	}

	if _, ok := set.elementMap[b]; !ok {
		return false
	}

	// TODO: complex type compare
	return set.FindHead(a) == set.FindHead(b)
}

func (set *DisjointSet[T]) Union(a, b T) {
	aF := element[T]{set.FindHead(a)}
	bF := element[T]{set.FindHead(b)}

	if aF == bF {
		return
	}

	big, small := aF, bF
	if set.sizeMap[small] > set.sizeMap[big] {
		small, big = big, small
	}

	set.fatherMap[small] = big
	set.sizeMap[big] = set.sizeMap[small] + set.sizeMap[big]

	delete(set.sizeMap, small)
}
