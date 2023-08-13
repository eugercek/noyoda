# `noyoda` Go Linter

`noyoda` is a Go linter, which reports Yoda Conditions.

# What are Yoda Conditions?
Yoda Condition style involves writing conditions expression/statements in a way that resembles the Yoda from Star Wars.
Instead of writing `if x == 10`, a Yoda Condition would be `if 10 == x`.
This approach aims to prevent unintended assignment by breaking the program at compile time.
For instance, it would prevent code like `if 10 = x`, which is invalid and would not compile.

# Why Yoda Conditions are Unnecessary in Go

`if x = 10 { ... } ` is invalid syntax in Go.

# Install and Usage

```bash
go install github.com/eugercek/noyoda/cmd/noyoda
cd mycode # go to your code's main package
noyoda ./...
```

# Roadmap

- [ ] `const` check
- [ ] flag for `const` check
- [ ] Auto fix
- [ ] Run tests for comprehensive set of popular go codebases, if there are many maybe Open a PR to golangci-lint