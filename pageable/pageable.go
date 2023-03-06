package pageable

// Pageable 分页
type Pageable interface {
	Page() int
	Size() int
	Skip() int
	Limit() int
}

type pageable struct {
	page int
	size int
}

func (p *pageable) Skip() int {
	return p.page*p.size - p.Size()
}

func (p *pageable) Limit() int {
	return p.size
}

func (p *pageable) Size() int {
	return p.size
}

func (p *pageable) Page() int {
	return p.page
}

func NewPageable(page, size int) Pageable {
	if size < 1 {
		size = 50
	}
	if page < 1 {
		page = 1
	}
	return &pageable{page: page, size: size}
}

func DefaultPageable() Pageable {
	return &pageable{page: 1, size: 50}
}
