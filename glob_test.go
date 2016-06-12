package glob

import "testing"

func TestGlob(t *testing.T) {
	test(t, "hello", "hello", false, true)
	test(t, "hello", "hallo", false, false)
	test(t, "h?llo", "hallo", false, true)
	test(t, "h[a-z]llo", "hmllo", false, true)
	test(t, "h[^a-z]llo", "hmllo", false, false)
	test(t, "h[^a-z]llo", "hMllo", false, true)
	test(t, "h[^a-z]ll?", "hMllo", false, true)
	test(t, "h*ll?", "hMllo", false, true)
	test(t, "h*ll", "hMllo", false, false)
	test(t, "h**ll", "hMllo", false, false)
	test(t, "**ll", "hMllo", false, false)
	test(t, "*ll*", "hMllo", false, true)
	test(t, "*l*", "hMllo", false, true)
	test(t, "**", "hMllo", false, true)
	test(t, "*", "hMllo", false, true)
	test(t, "*llo", "hMllo", false, true)
	test(t, "hMll*", "hMllo", false, true)
	test(t, "hMllo*", "hMllo", false, true)
}
func test(t *testing.T, pattern, str string, nocase, expect bool) {
	got := Match(pattern, str, nocase)
	if got != expect {
		t.Fatalf("expected %v, got %v (matching '%v' and '%v')", expect, got, pattern, str)
	}
}
