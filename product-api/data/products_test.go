package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "asda",
		Price: 11,
		SKU:   "abs-avd-asda",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
