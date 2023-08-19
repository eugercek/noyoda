`noyoda` is a Go linter, which reports yoda style conditionals.

# Flags

| Flag             | Description                                        | Default |
|------------------|----------------------------------------------------|---------|
| `-fix`           | Auto fix yoda conditions                           | no      |
| `-include-const` | Treat const as literal (`const x = 10; if x == a`) | no      |
| `-skip-range`    | Skip number range checks `(10 > a && a > 5)`       | yes     |

# Installation and Usage

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
- [x] Skip number range checks  (`10 > a  && a > 10`...)
- [ ] Run tests for comprehensive set of popular go codebases, if there are many maybe Open a PR to golangci-lint

# Useful Resources

- https://arslan.io/2017/09/14/the-ultimate-guide-to-writing-a-go-tool
- https://arslan.io/2019/06/13/using-go-analysis-to-write-a-custom-linter
- https://arslan.io/2020/07/07/using-go-analysis-to-fix-your-source-code/
- https://disaev.me/p/writing-useful-go-analysis-linter
- https://www.youtube.com/watch?v=k23xhJoTbI4
- https://www.youtube.com/watch?v=YRWCa84pykM

