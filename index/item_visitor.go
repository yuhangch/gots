package index

type ItemVisitor interface {
	VisitItem(item interface{})
}

type ListVisitor struct {
}
