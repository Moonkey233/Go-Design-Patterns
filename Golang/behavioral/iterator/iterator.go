// 迭代器模式

package iterator

// collection 集合接口
type collection interface {
	CreateIterator() iterator
	CreateReIterator() iterator
}

// part 具体的集合
type part struct {
	title  string
	number int
}

type partCollection struct {
	part
	parts []*part
}

func (u *partCollection) CreateIterator() iterator {
	return &partIterator{
		parts: u.parts,
	}
}

func (u *partCollection) CreateReIterator() iterator {
	return &partReIterator{
		index: len(u.parts) - 1,
		parts: u.parts,
	}
}

// iterator 迭代器
type iterator interface {
	HasNext() bool
	GetNext() *part
}

// partIterator 具体的迭代器
type partIterator struct {
	index int
	parts []*part
}

func (u *partIterator) HasNext() bool {
	if u.index < len(u.parts) {
		return true
	}
	return false
}

func (u *partIterator) GetNext() *part {
	if u.HasNext() {
		part := u.parts[u.index]
		u.index++
		return part
	}
	return nil
}

// partReIterator 逆迭代器
type partReIterator struct {
	index int
	parts []*part
}

func (u *partReIterator) HasNext() bool {
	if u.index >= 0 {
		return true
	}
	return false
}

func (u *partReIterator) GetNext() *part {
	if u.HasNext() {
		part := u.parts[u.index]
		u.index--
		return part
	}
	return nil
}
