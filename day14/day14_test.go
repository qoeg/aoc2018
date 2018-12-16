package day14

import (
	"testing"
)

func testAnswer1(skip, take int, t *testing.T) string {
	return Answer1(skip, take)
}

func Test_Answer1Skip9(t *testing.T) {
	exp := "5158916779"
	res := testAnswer1(9, 10, t)
	
	if res != exp {
		t.Errorf("Expected %s, got %s", exp, res)
	}
}

func Test_Answer1Skip5(t *testing.T) {
	exp := "0124515891"
	res := testAnswer1(5, 10, t)
	
	if res != exp {
		t.Errorf("Expected %s, got %s", exp, res)
	}
}

func Test_Answer1Skip18(t *testing.T) {
	exp := "9251071085"
	res := testAnswer1(18, 10, t)
	
	if res != exp {
		t.Errorf("Expected %s, got %s", exp, res)
	}
}

func Test_Answer1Skip2018(t *testing.T) {
	exp := "5941429882"
	res := testAnswer1(2018, 10, t)
	
	if res != exp {
		t.Errorf("Expected %s, got %s", exp, res)
	}
}

func testAnswer2(match string, after int, t *testing.T) int {
	return Answer2(match, after)
}

func Test_Answer2Match51589(t *testing.T) {
	exp := 9
	res := testAnswer2("51589", 14, t)
	
	if exp != res {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

func Test_Answer2Match01245(t *testing.T) {
	exp := 5
	res := testAnswer2("01245", 10, t)
	
	if exp != res {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

func Test_Answer2Match92510(t *testing.T) {
	exp := 18
	res := testAnswer2("92510", 23, t)
	
	if exp != res {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

func Test_Answer2Match59414(t *testing.T) {
	exp := 2018
	res := testAnswer2("59414", 2023, t)
	
	if exp != res {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}
