package testdata

func IfYodaCondition(n int) {
	if 10 == n { // want "yoda condition: 10 == n should be n == 10"
		return
	}
}

func ElseIfYodaCondition(n int) {
	if n == 100 {
		return
	} else if 10 == n { // want "yoda condition: 10 == n should be n == 10"
		return
	}
}

func SwitchYodaCondition(n int) {
	switch {
	case 10 == n: // want "yoda condition: 10 == n should be n == 10"
		return
	case 20 == n, 30 == n: // want "yoda condition: 20 == n should be n == 20" "yoda condition: 30 == n should be n == 30"
	}
}

func ConstYodaCondition(n int) {
	const x = 10

	if n == x {
		return
	}
}
