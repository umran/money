package money

import (
	"testing"
)

func TestRW(t *testing.T) {
	a := New(150000, MVR)
	b := New(126444, MVR)
	c := New(178301, MVR)
	d := New(508783, MVR)
	e := New(220599, MVR)
	f := New(310250, MVR)
	g := New(71025, MVR)
	h := New(225718, MVR)
	i := New(259164, MVR)

	ab, _ := a.Add(b)
	abc, _ := ab.Add(c)
	abcd, _ := abc.Add(d)
	abcde, _ := abcd.Add(e)
	abcdef, _ := abcde.Add(f)
	abcdefg, _ := abcdef.Add(g)
	abcdefgh, _ := abcdefg.Add(h)
	abcdefghi, _ := abcdefgh.Add(i)

	if abcdefghi.Amount() != int64(2050284) {
		t.Error("unexpected amount")
		t.Log(abcdefghi.Amount())
	}
}
