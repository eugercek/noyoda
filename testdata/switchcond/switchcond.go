package switchcond

func Switch(n int) {
	switch {
	case 10 == n: // want "yoda condition: 10 == n should be n == 10"
	case 20 == n, 30 == n: // want "yoda condition: 20 == n should be n == 20" "yoda condition: 30 == n should be n == 30"
	case n == 40:
	case n == 50, n == 60:
	}
}

func RecursiveSwitch(n int) {
	switch {
	case n == 100 || 10 == n: // want "yoda condition: 10 == n should be n == 10"
	case 10 == n || n == 200: // want "yoda condition: 10 == n should be n == 10"
	case n == 10 || n == 20:
	case 30 == n || 40 == n: // want "yoda condition: 30 == n should be n == 30" "yoda condition: 40 == n should be n == 40"
	}
}
