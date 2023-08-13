package ifcond

func If(n int) {
	if 10 == n { // want "yoda condition: 10 == n should be n == 10"
	}

	if n == 10 {
	}
}

func ElseIf(n int) {
	if n == 100 {
	} else if 10 == n { // want "yoda condition: 10 == n should be n == 10"
	}

	if n == 100 {
	} else if n == 200 {
	}
}

func RecursiveIf(n int) {
	if n == 100 || 10 == n { // want "yoda condition: 10 == n should be n == 10"
	}

	if 10 == n || n == 100 { // want "yoda condition: 10 == n should be n == 10"
	}

	if n == 100 || n == 200 {
	}

	if 10 == n || 20 == n { // want "yoda condition: 10 == n should be n == 10" "yoda condition: 20 == n should be n == 20"
	}

	if 10 == n || 20 == n || 30 == n { // want "yoda condition: 10 == n should be n == 10" "yoda condition: 20 == n should be n == 20" "yoda condition: 30 == n should be n == 30"

	}
}
