package main

func Bool(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}

func Float(v float64) *float64 {
	p := new(float64)
	*p = v
	return p
}
