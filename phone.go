package phone

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const (
	commonExtensions = `(ext|ex|x|xt|#|:)+[^0-9]*([-0-9]{1,})*#?$`
	commonNumber     = `[0-9]{1,}$`
	commonExtras     = `(\(0\)|[^0-9+]|^\+?00?)`
	formatTokens     = `(%[caAnflx])`
)

var (
	mu          sync.Mutex
	fmtEnum     = []string{"default", "default_with_extension", "europe", "us"}
	namedFormat = map[string]string{
		"default":                "+%c%a%n",
		"default_with_extension": "+%c%a%n%x",
		"europe":                 "+%c (0) %a %f %l",
		"us":                     "(%a) %f-%l",
	}
	commonExtraReplacements = map[string]string{
		"(0)": "+",
		"00":  "+",
		"+00": "+",
		"+0":  "+",
	}
	defaultCountryCode string
	defaultAreaCode    string
)

type Phone struct {
	NamedFormats       string
	N1Length           string
	Number             string `yaml:"number"`
	CountryCode        string `yaml:"country_code"`
	AreaCode           string `yaml:"area_code"`
	Extension          string `yaml:"extension"`
	DefaultCountryCode string
	DefaultAreaCode    string
}

func Parse(s string) (*Phone, error) {
	if s == "" {
		return nil, nil
	}
	sub, e := extractExtension(s)
	sub = normalize(sub)
	args, err := splitToParts(sub)
	if err != nil {
		return nil, err
	}
	c, err := New(args)
	if err != nil {
		return nil, err
	}
	c.Extension = e
	return c, nil
}

func IsValid(s string) bool {
	_, err := Parse(s)
	return err == nil
}

func New(args []string) (input *Phone, err error) {
	input = ArgsToCountry(args...)

	if input.N1Length == "" {
		input.N1Length = "3"
	}

	if input.CountryCode == "" {
		input.CountryCode = defaultCountryCode
	}

	if input.AreaCode == "" {
		input.AreaCode = defaultAreaCode
	}

	if strings.Trim(input.Number, "\t \n") == "" {
		err = errors.New("must enter number")
	}
	if strings.Trim(input.AreaCode, "\t \n") == "" {
		err = errors.New("must enter area code or set default")
	}
	if strings.Trim(input.CountryCode, "\t \n") == "" {
		err = errors.New("must enter country code or set default")
	}

	return input, err
}

func ArgsToCountry(args ...string) *Phone {
	c := &Phone{}
	switch len(args) {
	case 1:
		c.Number = args[0]
	case 2:
		c.Number = args[0]
		c.AreaCode = args[1]
	case 3:
		c.Number = args[0]
		c.AreaCode = args[1]
		c.CountryCode = args[2]
	case 4:
		c.Number = args[0]
		c.AreaCode = args[1]
		c.CountryCode = args[2]
		c.Extension = args[3]
	}
	return c
}

func (c *Phone) String() string {
	return c.Format("default")
}

func (c *Phone) Number1() string {
	data := []byte(c.Number)
	i, err := strconv.Atoi(c.N1Length)
	if err != nil {
		panic(err)
	}
	str := string(data[0:i])

	return str
}
func (c *Phone) Number2() string {
	data := []byte(c.Number)
	i, err := strconv.Atoi(c.N1Length)
	if err != nil {
		panic(err)
	}
	l := len(data) - i - 1
	str := string(data[l:])

	return str
}

func (c *Phone) Format(fmt string) string {
	if contains(fmtEnum, fmt) {
		return c.FormatNumber(namedFormat[fmt])
	}
	return c.FormatNumber(fmt)
}

func (c *Phone) AreaCodeLong() string {
	if c.AreaCode != "" {
		return fmt.Sprintf("0%s", c.AreaCode)
	}
	return ""
}

func (c *Phone) FormatNumber(fm string) string {
	var replacements = map[string]string{
		"%c": c.CountryCode,
		"%a": c.AreaCode,
		"%A": c.AreaCodeLong(),
		"%n": c.Number,
		"%f": c.Number1(),
		"%l": c.Number2(),
		"%x": c.Extension,
	}

	re := regexp.MustCompile(formatTokens)
	match := re.FindAllString(fm, -1)
	for _, m := range match {
		_s := replacements[m]
		fm = strings.Replace(fm, m, _s, 1)
	}

	fm = removeUselessPlus(fm)
	return fm
}

func SetDefaultCountryCode(code string) string {
	mu.Lock()
	defer mu.Unlock()
	defaultCountryCode = code
	return code
}

func SetDefaultAreaCode(code string) string {
	mu.Lock()
	defer mu.Unlock()
	defaultAreaCode = code
	return code
}

func extractExtension(s string) (string, string) {
	re := regexp.MustCompile(commonExtensions)
	subbed := re.FindString(s)
	if subbed != "" {
		re = regexp.MustCompile(commonExtensions)
		s = re.ReplaceAllString(s, "")
		return s, subbed
	} else {
		return s, ""
	}
}

func normalize(stringWithNumber string) string {
	re := regexp.MustCompile(commonExtras)
	match := re.FindAllString(stringWithNumber, -1)
	var s string
	for _, m := range match {
		s = commonExtraReplacements[m]
		stringWithNumber = re.ReplaceAllString(stringWithNumber, s)
	}
	return stringWithNumber
}

func splitToParts(s string) (args []string, err error) {
	c := detectCountry(s, defaultCountryCode)

	if c != nil {
		re := c.CountryCodeRegexp()
		s = re.ReplaceAllString(s, "0")
		c.CountryCode = "+" + c.CountryCode
	}

	if c == nil {
		err = errors.New("must specify country code")
		return nil, err
	}

	format := c.DetectFormat(s)
	if format == "" {
		return nil, err
	}

	r, _ := regexp.Compile(c.AreaCode)
	areaCode := r.FindString(s)

	n, _ := regexp.Compile(fmt.Sprintf("^0*(%s)", c.AreaCode))
	number := n.ReplaceAllString(s, "")

	args = append(args, number)
	args = append(args, areaCode)
	args = append(args, c.CountryCode)
	return args, nil
}

func removeUselessPlus(s string) string {
	re := regexp.MustCompile(`^(\+ \+)|^(\+\+)`)
	s = re.ReplaceAllString(s, "+")
	return s
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
