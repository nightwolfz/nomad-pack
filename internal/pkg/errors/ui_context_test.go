// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package errors

import (
	"testing"

	"github.com/shoenig/test/must"
)

func TestUIErrorContext_Add(t *testing.T) {
	testCases := []struct {
		inputUIErrorContext *UIErrorContext
		inputPrefix         string
		inputVal            string
		expectedOutput      []string
		name                string
	}{
		{
			inputUIErrorContext: NewUIErrorContext(),
			inputPrefix:         UIContextPrefixPackName,
			inputVal:            "foobar",
			expectedOutput:      []string{"Pack Name: foobar"},
			name:                "empty input context",
		},
		{
			inputUIErrorContext: &UIErrorContext{
				contexts: []string{"Pack Path: /go/src/github/why/"},
			},
			inputPrefix:    UIContextPrefixPackName,
			inputVal:       "foobar",
			expectedOutput: []string{"Pack Path: /go/src/github/why/", "Pack Name: foobar"},
			name:           "non-empty input context",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.inputUIErrorContext.Add(tc.inputPrefix, tc.inputVal)
			must.SliceContainsAll(t, tc.expectedOutput, tc.inputUIErrorContext.GetAll())
		})
	}
}

func TestUIErrorContext_Append(t *testing.T) {
	testCases := []struct {
		inputUIErrorContext *UIErrorContext
		inputAppendContext  *UIErrorContext
		expectedOutput      []string
		name                string
	}{
		{
			inputUIErrorContext: NewUIErrorContext(),
			inputAppendContext: &UIErrorContext{
				contexts: []string{"Pack Path: /go/src/github/why/"},
			},
			expectedOutput: []string{"Pack Path: /go/src/github/why/"},
			name:           "empty input context",
		},
		{
			inputUIErrorContext: &UIErrorContext{
				contexts: []string{"Pack Name: what is going on"},
			},
			inputAppendContext: &UIErrorContext{
				contexts: []string{"Pack Path: /go/src/github/why/"},
			},
			expectedOutput: []string{
				"Pack Path: /go/src/github/why/",
				"Pack Name: what is going on",
			},
			name: "empty input context",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.inputUIErrorContext.Append(tc.inputAppendContext)
			must.SliceContainsAll(t, tc.expectedOutput, tc.inputUIErrorContext.GetAll())
		})
	}
}

func TestUIErrorContext_Copy(t *testing.T) {
	testCases := []struct {
		inputUIErrorContext *UIErrorContext
		expectedOutput      *UIErrorContext
		name                string
	}{
		{
			inputUIErrorContext: NewUIErrorContext(),
			expectedOutput:      NewUIErrorContext(),
			name:                "empty input context",
		},
		{
			inputUIErrorContext: &UIErrorContext{
				contexts: []string{"Pack Path: /go/src/github/why/"},
			},
			expectedOutput: &UIErrorContext{
				contexts: []string{"Pack Path: /go/src/github/why/"},
			},
			name: "non-empty input context",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			must.Eq(t, tc.expectedOutput, tc.inputUIErrorContext.Copy())
		})
	}
}

func TestUIErrorContext_GetAll(t *testing.T) {
	testCases := []struct {
		inputUIErrorContext *UIErrorContext
		expectedOutput      []string
		name                string
	}{
		{
			inputUIErrorContext: NewUIErrorContext(),
			expectedOutput:      []string{},
			name:                "empty input context",
		},
		{
			inputUIErrorContext: &UIErrorContext{
				contexts: []string{"Pack Path: /go/src/github/why/"},
			},
			expectedOutput: []string{"Pack Path: /go/src/github/why/"},
			name:           "non-empty input context",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			must.SliceContainsAll(t, tc.expectedOutput, tc.inputUIErrorContext.GetAll())
		})
	}
}
