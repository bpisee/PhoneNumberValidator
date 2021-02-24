package webService

import (
	"testing"
)

//Test cases for verifying the phone number is valid or not
func TestVerifyPhoneNumber(t *testing.T) {
	var TestVerifyPhoneNumberTests = []struct {
		in       string
		expected bool
	}{
		{"23948291020", false},
		{"006498767938", true},
		{"00649a767938", false},
		{"6495476738", true},
		{"6195476738", true},
		{"+6195476738", true},
		{"+6495476738", true},
		{"+64954767345345338", false},
		{"0065338", false},
	}

	for _, tv := range TestVerifyPhoneNumberTests {
		actual := verifyPhoneNumber(tv.in)

		if actual != tv.expected {
			t.Errorf("verifyPhoneNumber(%s) is valid %t; expected %t", tv.in, actual, tv.expected)
		}
	}
}
