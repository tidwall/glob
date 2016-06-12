package glob

import "testing"

func TestGlob(t *testing.T) {
	test(t, "hello", "hello", true)
	test(t, "hello", "hallo", false)
	test(t, "h?llo", "hallo", true)
	test(t, "h[a-z]llo", "hmllo", true)
	test(t, "h[^a-z]llo", "hmllo", false)
	test(t, "h[^a-z]llo", "hMllo", true)
	test(t, "h[^a-z]ll?", "hMllo", true)
	test(t, "h*ll?", "hMllo", true)
	test(t, "h*ll", "hMllo", false)
	test(t, "h**ll", "hMllo", false)
	test(t, "**ll", "hMllo", false)
	test(t, "*ll*", "hMllo", true)
	test(t, "*l*", "hMllo", true)
	test(t, "**", "hMllo", true)
	test(t, "*", "hMllo", true)
	test(t, "*llo", "hMllo", true)
	test(t, "hMll*", "hMllo", true)
	test(t, "hMllo*", "hMllo", true)
}
func test(t *testing.T, pattern, str string, expect bool) {
	got := Match(pattern, str)
	if got != expect {
		t.Fatalf("expected %v, got %v (matching '%v' and '%v')", expect, got, pattern, str)
	}
}
