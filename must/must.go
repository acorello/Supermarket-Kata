package must

func Fn[Params, Res any](resOrError func(Params) (Res, error)) func(Params) Res {
	return DereFn(func(p Params) (*Res, error) {
		res, err := resOrError(p)
		return &res, err
	})
}

func DereFn[Params, Res any, ResPtr *Res](resOrError func(Params) (ResPtr, error)) func(Params) Res {
	resOrPanic := func(p Params) Res {
		if v, err := resOrError(p); err != nil {
			panic(err)
		} else {
			return *v
		}
	}
	return resOrPanic
}

func Work[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func WorkPtr[T any, P *T](v P, err error) T {
	if err != nil {
		panic(err)
	}
	return *v
}
