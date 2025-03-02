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
	heap []*Node[T]
}

func NewBinaryMaxHeap[T any]() *BinaryMaxHeap[T] {
	return &BinaryMaxHeap[T]{[]*Node[T]{}}
}

func (b *BinaryMaxHeap[T]) insert(val T) {
	b.heap = append(b.heap, &Node[T]{val, len(b.heap), nil, nil})
	b.bubbleUp()
}

func (b *BinaryMaxHeap[T]) remove() T {
	i := b.heap[0].Value
	b.bubbleDown()
	return i
}

func (b *BinaryMaxHeap[T]) bubbleUp() {
	h := b.heap
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
	h := b.heap
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
