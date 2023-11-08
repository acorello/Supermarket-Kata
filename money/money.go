package money

type Cents uint

func (my Cents) Percent(percent uint) Cents {
	amount := (uint(my) / 100) * percent
	return Cents(amount)
}
