
```sh
$ go list -m all
github.com/JayneJacobs/gomodules
github.com/davecgh/go-spew v1.1.0
github.com/pmezard/go-difflib v1.0.0
github.com/stretchr/objx v0.1.0
github.com/stretchr/testify v1.4.0
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
gopkg.in/yaml.v2 v2.2.2
rsc.io/quote v1.5.2
rsc.io/quote/v3 v3.1.0
rsc.io/sampler v1.3.0
:gomodules $ go mod tidy
go: downloading gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
go: extracting gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
:gomodules $ go mod tidy
:gomodules $ go list -m all
github.com/JayneJacobs/gomodules
github.com/davecgh/go-spew v1.1.0
github.com/pmezard/go-difflib v1.0.0
github.com/stretchr/objx v0.1.0
github.com/stretchr/testify v1.4.0
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
gopkg.in/yaml.v2 v2.2.2
rsc.io/quote v1.5.2
rsc.io/quote/v3 v3.1.0
rsc.io/sampler v1.3.0
```