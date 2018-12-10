package main

import "testing"

func TestIDCommon(t *testing.T) {
	tests := []struct {
		list         []string
		id           string
		expectNil    bool
		expectString string
	}{
		{[]string{}, "", true, ""},    // empty list should not match blank string
		{[]string{}, "xyz", true, ""}, // empty list should not match arbitrary string

		{[]string{"qwerty"}, "qwerty", true, ""}, // exact matches aren't valid
		{[]string{"qwerty"}, "xxerty", true, ""}, // two differ at the start
		{[]string{"qwerty"}, "qwxxty", true, ""}, // two differ in the middle
		{[]string{"qwerty"}, "qwerxx", true, ""}, // two differ at the end
		{[]string{"qwerty"}, "xwertx", true, ""}, // two differ at both ends
		{[]string{"qwerty"}, "qxerxy", true, ""}, // two differ near the middle

		{[]string{"qwerty", "asdfgh"}, "qwerty", true, ""}, // exact matches aren't valid
		{[]string{"qwerty", "asdfgh"}, "xxerty", true, ""}, // two differ at the start
		{[]string{"qwerty", "asdfgh"}, "qwxxty", true, ""}, // two differ in the middle
		{[]string{"qwerty", "asdfgh"}, "qwerxx", true, ""}, // two differ at the end
		{[]string{"qwerty", "asdfgh"}, "xwertx", true, ""}, // two differ at both ends
		{[]string{"qwerty", "asdfgh"}, "qxerxy", true, ""}, // two differ near the middle

		{[]string{"asdfgh", "qwerty"}, "qwerty", true, ""}, // exact matches aren't valid
		{[]string{"asdfgh", "qwerty"}, "xxerty", true, ""}, // two differ at the start
		{[]string{"asdfgh", "qwerty"}, "qwxxty", true, ""}, // two differ in the middle
		{[]string{"asdfgh", "qwerty"}, "qwerxx", true, ""}, // two differ at the end
		{[]string{"asdfgh", "qwerty"}, "xwertx", true, ""}, // two differ at both ends
		{[]string{"asdfgh", "qwerty"}, "qxerxy", true, ""}, // two differ near the middle

		{[]string{"qwerty"}, "xwerty", false, "werty"},           // one differs at the start
		{[]string{"qwerty"}, "qwxrty", false, "qwrty"},           // one differs in the middle
		{[]string{"qwerty"}, "qwertx", false, "qwert"},           // one differs at the end
		{[]string{"qwerty", "asdfgh"}, "xwerty", false, "werty"}, // one differs at the start
		{[]string{"qwerty", "asdfgh"}, "qwxrty", false, "qwrty"}, // one differs in the middle
		{[]string{"qwerty", "asdfgh"}, "qwertx", false, "qwert"}, // one differs at the end
		{[]string{"asdfgh", "qwerty"}, "xwerty", false, "werty"}, // one differs at the start
		{[]string{"asdfgh", "qwerty"}, "qwxrty", false, "qwrty"}, // one differs in the middle
		{[]string{"asdfgh", "qwerty"}, "qwertx", false, "qwert"}, // one differs at the end
	}

	for _, test := range tests {
		out := isMatch(test.list, test.id)
		if out == nil {
			if !test.expectNil {
				t.Errorf("failed: isMatch(%q, %q) -> nil, expected %q", test.list, test.id, test.expectString)
				continue
			}
		} else {
			if test.expectNil {
				t.Errorf("failed: isMatch(%q, %q) -> %q, expected nil", test.list, test.id, *out)
				continue
			}
			if *out != test.expectString {
				t.Errorf("failed: isMatch(%q, %q) -> %q, expected %q", test.list, test.id, *out, test.expectString)
				continue
			}
		}
	}
}
