
go mod download
go test . -v
go test . -v
go test . -v
go get rsc/sampler
go get github.com/rsc/sampler


go get rsc.io/sampler@v1.3.1

## Importing specific versions

import (
    "rsc.io/quote"
    quoteV3 "rsc.io/quote/v3"
)

