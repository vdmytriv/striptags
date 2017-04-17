// Copyright (C) 2017 Bolste, Inc. All Rights Reserved.
// This source is private property of Bolste, Inc.

package sanitation_test

import (
	"testing"

	"github.com/BolsteDev/core/sanitation"
)

func TestStripTags(t *testing.T) {
	data := []struct {
		input          string
		expectedOutput string
	}{
		{input: "<div>alert('you have been pwned')</div>", expectedOutput: "alert('you have been pwned')"},
		{input: "<b>&iexcl;Hi!</b> <script>...</script>", expectedOutput: "&iexcl;Hi! "},
		{input: "<b>&iexcl;Hi!</b> <script>alert('you have been pwned')</script>", expectedOutput: "&iexcl;Hi! "},
		{input: "<b>&iexcl;Hi!</b><script>...</script>", expectedOutput: "&iexcl;Hi!"},
	}

	for _, d := range data {
		t.Run(d.input, func(t *testing.T) {
			res := sanitation.StripTags(d.input)

			if res != d.expectedOutput {
				t.Errorf("Output:%s, Expected:%s\n", res, d.expectedOutput)
			}
		})
	}

}
