package days

type testStringIntPair struct {
	input    string
	expected int
}

type testIntPair struct {
	input    int
	expected int
}

type testStringBoolPair struct {
	input    string
	expected bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
