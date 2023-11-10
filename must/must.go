package must

func Fn[Params, Res any](resOrError func(Params) (Res, error)) func(Params) Res {
	resOrPanic := func(p Params) Res {
		if v, err := resOrError(p); err != nil {
			panic(err)
		} else {
			return v
		}
	}
	return resOrPanic
}
