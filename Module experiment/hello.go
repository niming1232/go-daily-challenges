package hello

import (
	"rsc.io/quote"            // Different minor versions and patches can in the same path
	quoteV3 "rsc.io/quote/v3" // Different majoy versions in different paths (e.g. v1.2.3: 1 major, 2 minor, 3 patch)
)

func Hello() string {
	return quote.Hello()
}

func Proverb() string {
	return quoteV3.Concurrency()
}
