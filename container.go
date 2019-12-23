package tcli

type Orientation uint8

const (
	OrientationRow Orientation = iota + 1
	OrientationColumn
)

type Container struct {
	View
	Views       []View
	Ratios      []float32
	x           int
	y           int
	w           int
	h           int
	orientation Orientation
}

func NewContainer(orientation Orientation) *Container {
	c := &Container{}
	c.orientation = orientation
	return c
}

func (c *Container) Position() (x, y int) {
	return c.x, c.y
}
func (c *Container) Move(x, y int) {
	c.x = x
	c.y = y
}

func (c *Container) Size() (w, h int) {
	return c.w, c.h
}
func (c *Container) Resize(w, h int) {
	c.w = w
	c.h = h
}

func (c *Container) Draw() {
	c.Clear()
	offset := 0
	lv := len(c.Views)
	for i, v := range c.Views {
		ratio := c.Ratios[i]
		switch c.orientation {
		case OrientationRow:
			w := int(float32(c.w-(lv-1)) * ratio)
			v.Move(c.x+offset, c.y)
			v.Resize(w, c.h)
			offset += w
			if i != lv-1 {
				for i := c.y; i < c.h; i++ {
					tSetCell(c.x+offset+1, i, 0x7c, DefaultStyle())
				}
				offset += 2
			}
		case OrientationColumn:
			h := int(float32(c.h-(lv-1)) * ratio)
			v.Move(c.x, c.y+offset)
			v.Resize(c.w, h)
			offset += h
			if i != lv-1 {
				for i := c.x; i < c.w; i++ {
					tSetCell(i, c.y+offset+1, 0x2d, DefaultStyle())
				}
				offset += 2
			}
		}
		v.Draw()
	}
}

// Clear - implement View
func (c *Container) Clear() {
	for _, v := range c.Views {
		v.Clear()
	}
}

func (c *Container) Handle(ev Event) {
	for _, v := range c.Views {
		v.Handle(ev)
	}
}

func (c *Container) AddView(view View, ratio float32) {
	c.Views = append(c.Views, view)
	c.Ratios = append(c.Ratios, ratio)
}

func (c *Container) ContentSize() (w, h int) {
	width := 0
	height := 0

	for _, v := range c.Views {
		w, h := v.Size()
		width += w
		height += h
	}
	return w, h
}
