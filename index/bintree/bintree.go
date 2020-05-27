package bintree

/*
	min extent 是所以已经插入的item的最小范围
	它使用一个启发式的值来为为零的要素构造一个非零的范围
	以一个非零范围开始，为了防止第一个被插入的要素两个方向上都为零范围
	这个值是 没经过优化的，但是只有一个要素被用这个值插入进来
*/
type BinTree struct {
	root      *Root
	minExtend float64
}

func NewBinTree() *BinTree {
	return &BinTree{
		root: NewRoot(),
	}
}

func ensureExtent(itemInterval Interval, minE float64) Interval {
	min, max := itemInterval.min(), itemInterval.max()
	if min != max {
		return itemInterval
	}
	if min == max {
		min = min - minE/2.0
		max = min + minE/2.0
	}
	return NewInterval(min, max)
}

/*
	查询整个树，找到所有压盖这个查询间隔的候选items
	如果间隔是null，会返回树上所有的items
	min max 需要为同一个值？
*/
func (bt *BinTree) query(x float64) *[]interface{} {
	return bt.queryByInterval(NewInterval(x, x))
}

func (bt *BinTree) queryByInterval(interval Interval) *[]interface{} {
	founds := make([]interface{}, 1)
	bt.queryAndAppend(interval, &founds)
	return &founds
}

func (bt *BinTree) queryAndAppend(interval Interval, founds *[]interface{}) {
	bt.root.addAllOverlapingItems(interval, founds)
}

func (bt *BinTree) insert(itemInterval Interval, item interface{}) {
	bt.collectStats(itemInterval)
	bt.root.insert(ensureExtent(itemInterval, bt.minExtend), item)
}

func (bt *BinTree) remove(itemInterval Interval, item interface{}) bool {
	insertInterval := ensureExtent(itemInterval, bt.minExtend)
	return bt.root.remove(insertInterval, item)
}

func (bt *BinTree) depth() int {
	if !(bt.root == nil) {
		return bt.root.depth()
	}
	return 0

}

func (bt *BinTree) size() int {
	if !(bt.root == nil) {
		return bt.root.size()
	}
	return 0

}

func (bt *BinTree) collectStats(interval Interval) {
	width := interval.width()
	if width < bt.minExtend && width > 0.0 {
		bt.minExtend = width
	}
}
