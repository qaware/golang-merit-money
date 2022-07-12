package business

import "testing"

func TestNewUuidFromStringError(t *testing.T) {
	_, err := NewUuidFromString("abc")
	if err == nil {
		t.Fail()
	}
}

func TestNewUuidFromStringSuccess(t *testing.T) {
	_, err := NewUuidFromString("165161")
	if err != nil {
		print(err.Error())
		t.Fail()
	}
}

func TestNewQACoinError(t *testing.T) {
	_, err := NewQaCoin("abc")
	if err == nil {
		t.Fail()
	}
}

func TestNewQACoinSuccess(t *testing.T) {
	_, err := NewQaCoin("123")
	if err != nil {
		print(err.Error())
		t.Fail()
	}
}

func TestSliceFind(t *testing.T) {
	testSlice := []string{"Alex", "Felix", "Markus"}
	toFind := "Alex"
	result := sliceFind(testSlice, func(str string) bool { return str == toFind })
	if result == nil || *result != toFind {
		t.Fail()
	}
}
