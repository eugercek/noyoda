package testdata

// 10, 20, 30 ... -> yoda condition
// 100, 200, 300 ... -> no yoda condition

func SimpleIf(n int) {
	if 10 == n { // want "yoda condition: 10 == n should be n == 10"
		return
	}

	if n == 10 {
		return
	}
}

func SimpleElseIf(n int) {
	if n == 100 {
		return
	} else if 10 == n { // want "yoda condition: 10 == n should be n == 10"
		return
	}

	if n == 100 {
		return
	} else if n == 200 {
		return
	}
}

func RecursiveIf(n int) {
	// 4 Condition May happen (n = noyoda, y = yoda)
	// n y
	// y n
	// n n
	// y y

	if n == 100 || 10 == n { // want "yoda condition: 10 == n should be n == 10"
		return
	}

	if 10 == n || n == 100 { // want "yoda condition: 10 == n should be n == 10"
		return
	}

	if n == 100 || n == 200 {
		return
	}

	if 10 == n || 20 == n { // want "yoda condition: 10 == n should be n == 10" "yoda condition: 20 == n should be n == 20"
		return
	}

	if 10 == n || 20 == n || 30 == n { // want "yoda condition: 10 == n should be n == 10" "yoda condition: 20 == n should be n == 20" "yoda condition: 30 == n should be n == 30"

	}
}

func SimpleSwitch(n int) {
	switch {
	case 10 == n: // want "yoda condition: 10 == n should be n == 10"
		return
	case 20 == n, 30 == n: // want "yoda condition: 20 == n should be n == 20" "yoda condition: 30 == n should be n == 30"
	case n == 40:
		return
	case n == 50, n == 60:
	}
}

func RecursiveSwitch(n int) {
	// 4 Condition May happen (n = noyoda, y = yoda)
	// n y
	// y n
	// n n
	// y y

	switch {
	case n == 100 || 10 == n: // want "yoda condition: 10 == n should be n == 10"
		return
	case 10 == n || n == 200: // want "yoda condition: 10 == n should be n == 10"
		return
	case n == 10 || n == 20:
	case 30 == n || 40 == n: // want "yoda condition: 30 == n should be n == 30" "yoda condition: 40 == n should be n == 40"
	}
}

func ConstYodaCondition(n int) {
	const x = 10

	if n == x {
		return
	}
}
