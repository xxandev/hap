package characteristic

import (
	"net/http"
	"testing"
)

func TestCharacteristicValue(t *testing.T) {

	c := NewBrightness()
	c.Val = 0

	n := 0
	c.ValFunc = func(*http.Request) interface{} {
		n++
		return n
	}

	if is, want := c.Value(), 1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := c.Value(), 2; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestCharacteristicSetValue(t *testing.T) {
	req := &http.Request{}
	c := NewBrightness()
	c.Val = 0

	n := 0
	c.OnValueUpdate(func(new, old int, r *http.Request) {
		if r != req {
			t.Fatal(r)
		}
		n++
	})

	c.SetValueRequest(10, req)
	if is, want := c.Value(), 10; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	c.SetValueRequest(20, req)
	if is, want := c.Value(), 20; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	if is, want := n, 2; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestCharacteristicValueTypeConversion(t *testing.T) {
	c := NewBrightness()
	c.Val = 5
	c.setValue(float64(20.5), nil)

	if is, want := c.Val, 20; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	c.setValue("91", nil)

	if is, want := c.Val, 91; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	c.setValue(true, nil)

	if is, want := c.Val, 1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestCharacteristicOnValueUpdate(t *testing.T) {
	c := NewBrightness()
	c.Val = 5

	d := false
	c.OnValueUpdate(func(new, old int, r *http.Request) {
		if r != nil {
			t.Fatal(r)
		}

		if is, want := old, 5; is != want {
			t.Fatalf("%v != %v", is, want)
		}

		if is, want := new, 6; is != want {
			t.Fatalf("%v != %v", is, want)
		}
		d = true
	})

	c.SetValue(6)

	if is, want := d, true; is != want {
		t.Fatalf("%v != %v", is, want)
	}
}

func TestValueChange(t *testing.T) {
	c := NewProgrammableSwitchEvent()
	c.Val = ProgrammableSwitchEventSinglePress

	changed := false
	c.OnValueUpdate(func(new, old int, r *http.Request) {
		changed = true
	})

	c.SetValue(ProgrammableSwitchEventSinglePress)

	if is, want := changed, true; is != want {
		t.Fatalf("%v != %v", is, want)
	}
}

func TestValueIngoreValueUpdate(t *testing.T) {
	c := NewBrightness()
	c.Val = 5

	c.OnValueUpdate(func(new, old int, r *http.Request) {
		t.Fatalf("Update value from %v to %v is unexpected", old, new)
	})

	c.SetValue(5)
}

func TestReadOnly(t *testing.T) {
	c := NewName()

	c.SetValue("Matthias")
	c.SetValueRequest("Gottfried", &http.Request{})

	if is, want := c.Value(), "Matthias"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}