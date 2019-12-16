# Go Modules

Documentation go mode help
go mod download (to cache)
go mod vendor vendor copy of dependencies.

 "https://research.swtch.com/vgo-import"

 go help modules
 cat go.mod
 go list -m all

A hash is a one way encryption. When you put something in the output will always be teh same if the input is the same.

# Class References

[Semantic Versioning]()

major.minor.patch

## Upgrading minor depandency

```bash
go list -m all

```sh
PASML-335382:modules jjacob151$ go list -m all
github.com/JayneJacobs/modules
github.com/davecgh/go-spew v1.1.0
github.com/pmezard/go-difflib v1.0.0
github.com/stretchr/objx v0.1.0
github.com/stretchr/testify v1.4.0
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
gopkg.in/yaml.v2 v2.2.7
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
```

go get package

```sh
go get gopkg.in/yaml.v2
```

adds to go.mod

```json
gopkg.in/yaml.v2 v2.2.7 // indirect
```
# Remove Unused Dependencies

go mod tidy