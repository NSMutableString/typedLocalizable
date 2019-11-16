package main

import "testing"

func TestExtractKeyFromLineGivenValidLineData(t *testing.T) {
	line := "\"lorem\" = \"ipsum\";"
	expectedKey := "lorem"
	key, _ := extractKeyFromLine(line)
	if key != expectedKey {
		t.Errorf("Did not get expected result")
	}
}

func TestExtractKeyFromLineGivenInvalidLineData(t *testing.T) {
	line := "lorem ipsum"
	expectedError := "Line did not match regex"
	_, err := extractKeyFromLine(line)
	if err.Error() != expectedError {
		t.Errorf("Did not get expected error")
	}
}