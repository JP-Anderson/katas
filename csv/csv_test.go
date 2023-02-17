package csv

import (
	"testing"
)

func TestSmallestTemperatureSpread(t *testing.T) {
	result := DayOfSmallestTempSpread("weather.dat")
	expected := "14"
	if result != expected {
		t.Errorf("Want %s, got %s", expected, result)
	} 
}

func TestSmallestGoalSpread(t *testing.T) {
	result := TeamWithMinForAgainstSpread("football.dat")	
	expected := "Aston_Villa"
	if result != expected {
                t.Errorf("Want %s, got %s", expected, result)
        }
}

func TestReadIntFieldRemovingAsterisk(t *testing.T) {
	type tc struct {
		in string
		out int
		err error
	}
	cases := []tc{
		{ in: "31", out: 31, err: nil},
		{ in: "78*", out: 78, err: nil},
	}

	for i, testCase := range cases {
		result, err := readIntFieldRemovingAsterisk(testCase.in)
		if result != testCase.out {
			t.Errorf("Failed case %d, Want %d, got %d", i, testCase.out, result)
		}
		if err != testCase.err {
			t.Errorf("Failed case %d, Want %v got err %v", i, testCase.err, err)
		}
		
	}
}
