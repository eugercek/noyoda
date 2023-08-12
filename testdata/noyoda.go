package testdata

func IfNoYodaCondition(n int) {
	if n == 10 {
		return
	}

	return
}

func ElseIfNoYodaCondition(n int) {
	if n == 100 {
		return
	} else if n == 10 {
		return
	}
}

func SwitchNoYodaCondition(n int) {
	switch {
	case n == 10:
		return
	}
}

func ConstNoYodaCondition(n int) {
	const x = 10

	if n == x {
		return
	}
}
