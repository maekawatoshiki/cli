package cli

import (
	"testing"
)

func TestOptions(t *testing.T) {
	type MyOptions struct {
		username string `short:"u",long:"username"`
	}

	AnalyzeOptions(MyOptions{
		username: "",
	})
}
