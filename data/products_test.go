package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "raf",
		Price: 3.00,
		SKU:   "as-sa-sa",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
