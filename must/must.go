package must

func Fn[Arg, Res any](fn func(Arg) (Res, error)) func(Arg) Res {
	return DereFn(func(p Arg) (*Res, error) {
		res, err := fn(p)
		return &res, err
	})
}

func DereFn[Arg, Res any](fn func(Arg) (*Res, error)) func(Arg) Res {
	panickyFn := func(p Arg) Res {
		if v, err := fn(p); err != nil {
			panic(err)
		} else {
			return *v
		}
	}
	return panickyFn
}

func Work[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func WorkPtr[T any](v *T, err error) T {
	if err != nil {
		panic(err)
	}
	return *v
}
