package rangecond

// These are not range statement
func YodaConditions(n, a int) {
	if 10 < n && 10 < a { // want "yoda condition: 10 < n should be n < 10" "yoda condition: 10 < a should be a < 10"

	}
	if 10 < n && a < 10 { // want "yoda condition: 10 < n should be n < 10"
	}

	if n > 10 && 10 < a { // want "yoda condition: 10 < a should be a < 10"
	}

	if n > 10 && a > 10 {
	}
}

func RangeSameOperator(n int) {
	if 20 > n && n > 10 {
	}

	if 10 < n && n < 20 {
	}

	if 20 >= n && n >= 10 {
	}

	if 10 <= n && n <= 20 {
	}
}

func RangeDifferentOperator(n int) {
	if 20 > n && n >= 10 {
	}

	if 10 < n && n <= 20 {
	}

	if 20 >= n && n > 10 {
	}

	if 10 <= n && n < 20 {
	}
}
