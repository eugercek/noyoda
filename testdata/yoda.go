package testdata

// Generally first yoda conditions tested afterwards non yoda conditions tested (to not create wrong diagnostics)

// 10, 20, 30 ... -> yoda condition
// 100, 200, 300 ... -> no yoda condition

// In recursive tests we check 4 condition (n = noyoda y = yoda)
// n y
// y n
// n n
// y y

// In case test we check
// y
// y,y
// n
// n,n

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

func ConstIf(n int) {
	const x = 10

	if x == n { // want "yoda condition: x == n should be n == x"
	}

	if n == x {
	}
}

func ConstSwitch(n int) {
	const x, y, z, a, b, c = 10, 20, 30, 40, 50, 60

	switch {
	case x == n: // want "yoda condition: x == n should be n == x"
	case y == n, z == n: // want "yoda condition: y == n should be n == y" "yoda condition: z == n should be n == z"
	case n == a:
	case n == b, n == c:
	}
}

func RecursiveConstIf(n int) {
	const x, y, z = 10, 20, 30
	if n == 100 || x == n { // want "yoda condition: x == n should be n == x"
	}

	if x == n || n == 100 { // want "yoda condition: x == n should be n == x"
	}

	if n == 100 || n == 200 {
	}

	if x == n || y == n { // want "yoda condition: x == n should be n == x" "yoda condition: y == n should be n == y"
	}

	if x == n || y == n || z == n { // want "yoda condition: x == n should be n == x" "yoda condition: y == n should be n == y" "yoda condition: z == n should be n == z"
	}
}

func RecursiveConstSwitch(n int) {
	const a, b, c, d = 10, 20, 30, 40
	switch {
	case n == 100 || a == n: // want "yoda condition: a == n should be n == a"
	case a == n || n == 200: // want "yoda condition: a == n should be n == a"
	case n == a || n == b:
	case c == n || d == n: // want "yoda condition: c == n should be n == c" "yoda condition: d == n should be n == d"
	}
}
