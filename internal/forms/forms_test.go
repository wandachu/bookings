package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have have required fields when it does")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid when email field does not exist")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "abc")
	form = New(postedValues)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("shows valid when it is not")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "abc@g.com")
	form = New(postedValues)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("shows not valid when it is")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows minLength for non-existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedValues := url.Values{}
	postedValues.Add("some_field", "some value") // 10 chars long
	form = New(postedValues)
	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("form shows minLength of 100 met when data is shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "some value") // 10 chars long
	form = New(postedValues)
	form.MinLength("another_field", 5)
	if !form.Valid() {
		t.Error("form shows minLength of 5 not met when data is longer")
	}

	isError = form.Errors.Get("x")
	if isError != "" {
		t.Error("should have no error, but shows as an empty string")
	}

}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("shows form does not have field when it should")
	}
}
