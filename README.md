# phone

Go library for phone number parsing, validation, and formatting.

## Install

```shell
go get -u github.com/amrfaissal/phone
```

### Automatic country and area code detection

Phone does its best to automatically detect the country and area code while parsing.

To do this, phone uses data stored in YAML format under `country.go`.

Each country code can have a regular expression named `area_code` that describes what the area code for that particular country looks like.

If an `area_code` regular expression isn't specified, a default value which is considered correct for the US will be used.

If your country has phone numbers longer that 8 digits - excluding country and area code - you can specify that within the country's configuration in `country.go` file.

### Validation

Validating is very relaxed, basically it strips out everything that's not a number or '+' character:

```go
phone.IsValid("random 091/512-5486 random")
```

### Formatting

Formating is done via the `#format` method. The method accepts a `Symbol` or a `String`.

When given a string, it interpolates the string with the following fields:

* %c - country_code (385)
* %a - area_code (91)
* %A - area_code with leading zero (091)
* %n - number (5125486)
* %f - first @@n1_length characters of number (configured through country n1_length), default is 3 (512)
* %l - last characters of number (5486)
* %x - the extension number

```go
parsedPhone = phone.Parse("+385915125486")
parsedPhone.String() // => "+385915125486"
parsedPhone.Format("%A/%f-%l") // => "091/512-5486"
parsedPhone.Format("+ %c (%a) %n") // => "+ 385 (91) 5125486"
```

When given a symbol it is used as a lookup for the format in the `namedFormats` map.

```go
parsedPhone.Format("europe") // => "+385 (0) 91 512 5486"
parsedPhone.Format("us") // => "(234) 123-4567"
parsedPhone.Format("default_with_extension") // => "+3851234567x143"
```

### Finding countries by their ISO code

If you don't have the country code, but you know from other sources what country a phone is from, you can retrieve the country using the country ISO code (such as 'de', 'es', 'us', ...).

### Phone Initialization

To initialize a new phone struct with the number, area code, country code and extension number:

```go
args := []string{"5125486", "91", "385"}
phone.New(args)
```

```go
args := []string{"5125486", "91", "385", "143"}
phone.New(args)
```

### Parsing

Create a new phone object by parsing from a string. The library does it's best to detect the country and area codes:

```go
phone.Parse("+385915125486")
phone.Parse("00385915125486")
```

If the country or area code isn't given in the string, you must set it, otherwise it doesn't work:

```go
args := []string{"091/512-5486", "385"}
phone.New(args)
```

If you feel that it's tedious, set the default country code once:

```go
phone.SetDefaultCountryCode("385")
args := []string{"091/512-5486"}
phone.New(args)
```

Same goes for the area code:

```go
args := []string{"5125486", "91"}
phone.New(args)
```

OR

```go
phone.SetDefaultCountryCode("385")
phone.SetDefaultAreaCode("47")
args := []string{"451-588"}
phone.New(args)
```

## Adding and maintaining countries

From time to time, the specifics about your countries information may change. You can add or update your countries configuration by editing `country.go` file.

These are the supported attributes for countries configuration:

* `country_code`: Required. A string representing your country's international dialling code. e.g. "123"
* `national_dialing_prefix`: Required. A string representing your default dialling prefix for national calls. e.g. "0"
* `char_3_code`: Required. A string representing a country's ISO code. e.g. "US"
* `name`: Required. The name of the country. e.g. "Denmark"
* `international_dialing_prefix`: Required. The dialling prefix a country typically uses when making international calls. e.g. "0"
* `area_code`: Optional. A regular expression detailing valid area codes. Default: "\d{3}" i.e. any 3 digits.
* `max_num_length`: Optional. The maximum length of a phone number after country and area codes have been removed. Default: 8

## License

MIT License
