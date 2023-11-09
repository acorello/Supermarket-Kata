package money

type Cents int

func (c Cents) Mul(v int) Cents {
	return c * Cents(v)
}
