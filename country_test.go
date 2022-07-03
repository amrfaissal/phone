package phone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	valid := IsValid("+385915125486")
	assert.True(t, valid)
}

func TestParse(t *testing.T) {
	phone, err := Parse("+00385915125486")
	assert.Nil(t, err)
	assert.Equal(t, phone.CountryCode, "+385")
	assert.Equal(t, phone.String(), "+385915125486")
}

func TestNormalize(t *testing.T) {
	normalizedPhone := normalize("+00385915125486")
	assert.Equal(t, normalizedPhone, "+385915125486")
}

func TestNew(t *testing.T) {
	args := []string{"5125486", "91", "385", "143"}
	phone, err := New(args)
	assert.Nil(t, err)
	assert.Equal(t, phone.String(), "+385915125486")
	assert.Equal(t, phone, &Phone{N1Length: "3", Number: "5125486", CountryCode: "385", AreaCode: "91", Extension: "143"})
}

func TestFormat(t *testing.T) {
	phone, err := Parse("+00385915125486x148")
	assert.Nil(t, err)

	f := phone.Format("%A/%f-%l")
	assert.Equal(t, f, "091/512-5486")

	n := phone.Format("+ %c (%a) %n")
	assert.Equal(t, n, "+385 (91) 5125486")

	europe := phone.Format("europe")
	assert.Equal(t, europe, "+385 (0) 91 512 5486")

	us := phone.Format("us")
	assert.Equal(t, us, "(91) 512-5486")

	ex := phone.Format("default_with_extension")
	assert.Equal(t, ex, "+385915125486x148")
}

func TestSetDefaultCountryCodeAndAreaCode(t *testing.T) {
	SetDefaultAreaCode("4")
	SetDefaultCountryCode("32")

	phone, err := Parse("451-588")
	assert.Nil(t, err)

	f := phone.Format("%A/%f-%l")
	assert.Equal(t, f, "04/515-1588")

	n := phone.Format("+ %c (%a) %n")
	assert.Equal(t, n, "+32 (4) 51588")

	europe := phone.Format("europe")
	assert.Equal(t, europe, "+32 (0) 4 515 1588")

	us := phone.Format("us")
	assert.Equal(t, us, "(4) 515-1588")

	ex := phone.Format("default_with_extension")
	assert.Equal(t, ex, "+32451588")
}

func TestFindByCountryIsoCode(t *testing.T) {
	country := FindByCountryIsoCode("BE")
	assert.NotNil(t, country)
	assert.Equal(t, country.CountryCode, "32")
}

func TestFindByCountryCode(t *testing.T) {
	country := FindByCountryCode("not_found")
	assert.Nil(t, country)

	country = FindByCountryCode("32")
	assert.NotNil(t, country)
	assert.Equal(t, country.Name, "Belgium")
}
