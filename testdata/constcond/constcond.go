package constcond

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
