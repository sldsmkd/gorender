package tuples

type tuple struct {
	x, y, z, w	float32
}

func (t *tuple) Init(x, y, z, w float32) *tuple {
	t.x = x
	t.y = y
	t.z = z
	t.w = w
	return t
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