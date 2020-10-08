package formatters

import (
	"io"

	"github.com/hemanthgk10/tfsec/pkg/app/tfsec/scanner"
)

// Formatter formats scan results into a specific format
type Formatter func(w io.Writer, results []scanner.Result) error
