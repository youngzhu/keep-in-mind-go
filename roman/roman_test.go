package roman

import "testing"

func TestParseToArabic_success(t *testing.T) {
	arabic, err := ParseToArabic("XLII")
	if err != nil {
		t.Errorf("err should be nil, but %v", err)
	}
	if arabic != 42 {
		t.Errorf("should be 42, but %v", arabic)
	}
}

func TestParseToArabic_fail(t *testing.T) {
	_, err := ParseToArabic("FOO")
	if err == nil {
		t.Error("should be an error!")
	}
}

func TestParseToRoman_success(t *testing.T) {
	roman, err := ParseToRoman(42)
	if err != nil {
		t.Errorf("err should be nil, but %v", err)
	}
	want := "XLII"
	if roman != want {
		t.Errorf("should be %q, but %q", want, roman)
	}
}

func TestParseToRoman_fail(t *testing.T) {
	_, err := ParseToRoman(4000)
	if err == nil {
		t.Error("should be an error!")
	}
}
