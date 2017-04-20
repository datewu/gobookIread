package proto

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtCloner()
	if shirtCache == nil {
		t.Fatal("Received cache was nil")
	}

	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	if item1 == whitePrototype {
		t.Errorf("item1 cannot be equal to the white prototype")
	}

	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt1 couldnot be done successfully")
	}
	shirt1.SKU = "abbcc"

	item2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}

	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt2 couldnot be donw successfully")
	}

	if shirt1.SKU == shirt2.SKU {
		t.Errorf("SKU's of shirt1 and shirt2 must be different")
	}

	if shirt1 == shirt2 {
		t.Errorf("shirt1 cannot be equal to shirt2")
	}

	t.Logf("LOG: %s\n", shirt1.GetInfo())
	t.Logf("LOG: %s\n", shirt2.GetInfo())

	t.Logf("LOG: The memoery positions of the shirts are different %p != %p \n", &shirt1, &shirt2)
}
