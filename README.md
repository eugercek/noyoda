`noyoda` is a Go linter, which reports Yoda Conditions.

Yoda Condition style involves writing conditions expression/statements in a way that resembles the Yoda from Star Wars.
Instead of writing `if x == 10`, a Yoda Condition would be `if 10 == x`.
This approach aims to prevent unintended assignment by breaking the program at compile time.
For instance, it would prevent code like `if 10 = x`, which is invalid and would not compile.

Because `if x = 10 { ... } ` is invalid syntax in Go, there's no need for yoda conditions.

# Install and Usage

```bash
go install github.com/eugercek/noyoda/cmd/noyoda
cd mycode # go to your code's main package
noyoda ./...
```

# Roadmap

- [x] `if` check
- [x] `switch` check
- [x] `const` check
- [x] Flag for `const`
- [x] Recursive check
- [x] Auto fix
- [ ] Run tests for comprehensive set of popular go codebases, if there are many maybe Open a PR to golangci-lint