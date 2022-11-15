package main

import "testing"

func TestDecodeDigit(t *testing.T) {
	testCases := []struct {
		Name   string
		C1     string
		C2     string
		C3     string
		Expect string
	}{
		{
			Name:   "decodeDigit returns a digit.",
			C1:     " _ ",
			C2:     " _|",
			C3:     " _|",
			Expect: "3",
		},
		{
			Name:   "decodeDigit returns a undecode digit.",
			C1:     "   ",
			C2:     " _|",
			C3:     " _|",
			Expect: "?",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			result := decodeDigit(tc.C1, tc.C2, tc.C3)
			if result != tc.Expect {
				t.Errorf("decodeDigit was incorrect, got: %s, want: %s", result, tc.Expect)
			}
		})
	}
}

func TestCheckSum(t *testing.T) {
	testCases := []struct {
		Name     string
		Possible string
		Expect   string
	}{
		{
			Name:     "checkSum returns OK flag.",
			Possible: "000000000",
			Expect:   "OK",
		},
		{
			Name:     "checkSum returns ERR flag.",
			Possible: "111111111",
			Expect:   "ERR",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			result := checkSum(tc.Possible)
			if result != tc.Expect {
				t.Errorf("checkSum was incorrect, got: %s, want: %s", result, tc.Expect)
			}
		})
	}
}
