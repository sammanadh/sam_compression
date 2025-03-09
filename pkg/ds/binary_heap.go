package ds

// we will use binary max heap to implement priority queue
type Node[T any] struct {
	Value  T
	Weight int
	Left   *Node[T]
	Right  *Node[T]
}

func NewNode[T any](val T, weight int) *Node[T] {
	return &Node[T]{val, weight, nil, nil}
}

type BinaryMaxHeap[T any] struct {
	Heap []*Node[T]
}

func NewBinaryMaxHeap[T any]() *BinaryMaxHeap[T] {
	return &BinaryMaxHeap[T]{[]*Node[T]{}}
}

func (b *BinaryMaxHeap[T]) Insert(val T, weight int) {
	b.Heap = append(b.Heap, &Node[T]{val, weight, nil, nil})
	b.bubbleUp()
}

func (b *BinaryMaxHeap[T]) remove() T {
	i := b.Heap[0].Value
	b.bubbleDown()
	return i
}

// func (b *BinaryMaxHeap[T]) CreateHuffmanTree() *Node[T] {
// 	for len(b.Heap) > 1 {
// 		l := len(b.Heap) - 1
// 		combinedWeight := b.Heap[l].Weight + b.Heap[l-1].Weight
// 		var zero T // zero value for type "T"
// 		newSmallestNode := &Node[T]{Value: zero, Weight: combinedWeight, Left: b.Heap[l], Right: b.Heap[l-1]}
// 		b.Heap = b.Heap[:l]
// 		b.Heap[l-1] = newSmallestNode
// 		b.bubbleUp()
// 	}
// 	return b.Heap[0]
// }

func (b *BinaryMaxHeap[T]) bubbleUp() {
	h := b.Heap
	i := len(h)
	for i > 0 && h[i-1].Weight > h[i].Weight {
		parentIdx := (i - 1) / 2
		temp := h[parentIdx]
		h[parentIdx] = h[i]
		h[i] = temp
		i = parentIdx
	}
}

func (b *BinaryMaxHeap[T]) bubbleDown() {
	h := b.Heap
	lstIdx := len(h) - 1
	lowest := h[lstIdx]
	h = h[:lstIdx]

	h[0] = lowest

	i := 0
	for i < lstIdx-1 {
		left := (i * 2) + 1
		right := (i * 2) + 1

		if !(h[left].Weight < h[i].Weight) || !(h[right].Weight < h[i].Weight) {
			idxToSwap := left
			if h[right].Weight > h[left].Weight {
				idxToSwap = right
			}

			temp := h[idxToSwap]
			h[idxToSwap] = h[i]
			h[i] = temp
		}
	}
}
