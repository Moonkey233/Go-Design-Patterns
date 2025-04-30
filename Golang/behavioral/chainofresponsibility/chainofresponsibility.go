// 责任链模式

package chainofresponsibility

import "fmt"

type Department interface {
	execute(*Do)
	setNext(Department)
}

type aPart struct {
	next Department
}

func (r *aPart) execute(p *Do) {
	if p.aPartDone {
		fmt.Println("aPart done")
		r.next.execute(p)
		return
	}
	fmt.Println("aPart")
	p.aPartDone = true
	r.next.execute(p)
}

func (r *aPart) setNext(next Department) {
	r.next = next
}

type bPart struct {
	next Department
}

func (d *bPart) execute(p *Do) {
	if p.bPartDone {
		fmt.Println("bPart done")
		d.next.execute(p)
		return
	}
	fmt.Println("bPart")
	p.bPartDone = true
	d.next.execute(p)
}

func (d *bPart) setNext(next Department) {
	d.next = next
}

type endPart struct {
	next Department
}

func (c *endPart) execute(p *Do) {
	if p.endPartDone {
		fmt.Println("endPart Done")
	}
	fmt.Println("endPart")
}

func (c *endPart) setNext(next Department) {
	c.next = next
}

type Do struct {
	aPartDone   bool
	bPartDone   bool
	endPartDone bool
}
