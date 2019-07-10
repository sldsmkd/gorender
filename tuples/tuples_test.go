package tuples

import (
	"fmt"
	"github.com/DATA-DOG/godog"
)

func (t *tuple) aTuple(x, y, z, w float32) error {
	t.Tuple(x, y, z, w)
	return nil
}

func (t *tuple) ax(x float32) error {
	if t.x != x {
		return fmt.Errorf("Expected x == %f, found %f", x, t.x)
	}
	return nil
}

func (t *tuple) ay(y float32) error {
	if t.y != y {
		return fmt.Errorf("Expected y == %f, found %f", y, t.y)
	}
	return nil
}

func (t *tuple) az(z float32) error {
	if t.z != z {
		return fmt.Errorf("Expected z == %f, found %f", z, t.z)
	}
	return nil
}

func (t *tuple) aw(w float32) error {
	if t.w != w {
		return fmt.Errorf("Expected w == %f, found %f", w, t.w)
	}
	return nil
}

func (t *tuple) aIsAPoint() error {
	if !t.IsAPoint() {
		return fmt.Errorf("Expected a Point")
	}
	return nil
}

func (t *tuple) aIsAVector() error {
	if !t.IsAVector() {
		return fmt.Errorf("Expected a Vector")
	}
	return nil
}

func (t *tuple) aIsNotAPoint() error {
	if t.IsAPoint() {
		return fmt.Errorf("Didn't expect a Point")
	}
	return nil
}

func (t *tuple) aIsNotAVector() error {
	if t.IsAVector() {
		return fmt.Errorf("Didn't expect a Vector")
	}
	return nil
}

func (t *tuple) pPoint(x, y, z float32) error {
	t.Point(x, y, z)
	return nil
}

func (t *tuple) vVector(x, y, z float32) error {
	t.Vector(x, y, z)
	return nil
}

func FeatureContext(s *godog.Suite) {
	a := &tuple{}
	p := &tuple{}
	v := &tuple{}
	s.Step(`^a ← tuple\((-?\d+\.\d+), (-?\d+\.\d+), (-?\d+\.\d+), (-?\d+\.\d+)\)$`, a.aTuple)
	s.Step(`^a\.x = (-?\d+\.\d+)$`, a.ax)
	s.Step(`^a\.y = (-?\d+\.\d+)$`, a.ay)
	s.Step(`^a\.z = (-?\d+\.\d+)$`, a.az)
	s.Step(`^a\.w = (-?\d+\.\d+)$`, a.aw)
	s.Step(`^a is a point$`, a.aIsAPoint)
	s.Step(`^a is not a vector$`, a.aIsNotAVector)
	s.Step(`^a is not a point$`, a.aIsNotAPoint)
	s.Step(`^a is a vector$`, a.aIsAVector)
	s.Step(`^p ← Point\((-?\d+), (-?\d+), (-?\d+)\)$`, p.pPoint)
	s.Step(`^p = Tuple\((-?\d+), (-?\d+), (-?\d+), (-?\d+)\)$`, p.aTuple)
	s.Step(`^v ← Vector\((-?\d+), (-?\d+), (-?\d+)\)$`, v.vVector)
	s.Step(`^v = Tuple\((-?\d+), (-?\d+), (-?\d+), (-?\d+)\)$`, v.aTuple)
}

