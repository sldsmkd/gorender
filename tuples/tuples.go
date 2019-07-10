package tuples

type tuple struct {
	x, y, z, w	float32
}

func (t *tuple) Tuple(x, y, z, w float32) *tuple {
	t.x = x
	t.y = y
	t.z = z
	t.w = w
	return t
}

func (t *tuple) Point(x, y, z float32) *tuple {
	t.x = x
	t.y = y
	t.z = z
	t.w = 1.0
	return t
}

func (t *tuple) Vector(x, y, z float32) *tuple {
	t.x = x
	t.y = y
	t.z = z
	t.w = 0.0
	return t
}

func Tuple(x, y, z, w float32) *tuple {
	t := tuple{}
	t.Tuple(x, y, z, w)
	return &t
}

func Point(x, y, z float32) *tuple {
	t := tuple{}
	t.Point(x, y, z)
	return &t
}

func Vector(x, y, z float32) *tuple {
	t := tuple{}
	t.Vector(x, y, z)
	return &t
}

func (t *tuple) IsAVector() bool {
	if t.w == 0 {
		return true
	}
	return false
}

func (t *tuple) IsAPoint() bool {
	if t.w == 1 {
		return true
	}
	return false
}