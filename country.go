package phone

import (
	"fmt"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

const countries = `
"676": 
  country_code: "676"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: TO
  name: Tonga
  international_dialing_prefix: "0"
"54": 
  country_code: "54"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: AR
  name: Argentina
  international_dialing_prefix: "0"
"506": 
  country_code: "506"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: CR
  name: Costa Rica
  international_dialing_prefix: "0"
"251": 
  country_code: "251"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: ET
  name: Ethiopia
  international_dialing_prefix: "0"
"590": 
  country_code: "590"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: GP
  name: Guadeloupe
  international_dialing_prefix: "0"
"82": 
  country_code: "82"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: KR
  name: Korea, Republic of
  international_dialing_prefix: "1"
"223": 
  country_code: "223"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: ML
  name: Mali
  international_dialing_prefix: "0"
"420": 
  country_code: "420"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: CZ
  name: Czech Republic
  international_dialing_prefix: "0"
"252": 
  country_code: "252"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: SO
  name: Somalia
  international_dialing_prefix: "0"
"677": 
  country_code: "677"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: SB
  name: Solomon Islands
  international_dialing_prefix: "0"
"421": 
  country_code: "421"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: SK
  name: Slovakia
  international_dialing_prefix: "0"
"507": 
  country_code: "507"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: PA
  name: Panama
  international_dialing_prefix: "0"
"591": 
  country_code: "591"
  national_dialing_prefix: "10"
  char_2_code: "10"
  char_3_code: BO
  name: Bolivia
  international_dialing_prefix: "10"
"224": 
  country_code: "224"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: GN
  name: Guinea
  international_dialing_prefix: "0"
"84": 
  country_code: "84"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: VN
  name: Viet Nam
  international_dialing_prefix: "0"
"678": 
  country_code: "678"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: VU
  name: Vanuatu
  international_dialing_prefix: "0"
"27": 
  country_code: "27"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: ZA
  name: South Africa
  international_dialing_prefix: "0"
  area_code: "800|86[01]|[1-9]\\d"
  max_num_length: 11
"508":
  country_code: "508"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: PM
  name: Saint Pierre And Miquelon
  international_dialing_prefix: "0"
"55": 
  country_code: "55"
  national_dialing_prefix: "14"
  char_2_code: "14"
  char_3_code: BR
  name: Brazil
  international_dialing_prefix: "14"
"253": 
  country_code: "253"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: DJ
  name: Djibouti
  international_dialing_prefix: "0"
"592": 
  country_code: "592"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: GY
  name: Guyana
  international_dialing_prefix: "0"
"225": 
  country_code: "225"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: CI
  name: "C\xC3\xB4te D'Ivoire"
  international_dialing_prefix: "0"
"56": 
  country_code: "56"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: CL
  name: Chile
  international_dialing_prefix: "0"
"679": 
  country_code: "679"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: FJ
  name: Fiji
  international_dialing_prefix: "0"
"509": 
  country_code: "509"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: HT
  name: Haiti
  international_dialing_prefix: "0"
"593": 
  country_code: "593"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: EC
  name: Ecuador
  international_dialing_prefix: "0"
"254": 
  country_code: "254"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: KE
  name: Kenya
  international_dialing_prefix: "0"
"226": 
  country_code: "226"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: BF
  name: Burkina Faso
  international_dialing_prefix: "0"
"423": 
  country_code: "423"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: LI
  name: Liechtenstein
  international_dialing_prefix: "0"
"255": 
  country_code: "255"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: TZ
  name: Tanzania, United Republic of
  international_dialing_prefix: "0"
"227": 
  country_code: "227"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: NE
  name: Niger
  international_dialing_prefix: "0"
"594": 
  country_code: "594"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: GF
  name: French Guiana
  international_dialing_prefix: "0"
"86": 
  country_code: "86"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: CN
  name: China
  international_dialing_prefix: "0"
"960": 
  country_code: "960"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: MV
  name: Maldives
  international_dialing_prefix: "0"
"57": 
  country_code: "57"
  national_dialing_prefix: "5"
  char_2_code: "5"
  char_3_code: CO
  name: Colombia
  international_dialing_prefix: "5"
"58": 
  country_code: "58"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: VE
  name: Venezuela, Bolivarian Republic of
  international_dialing_prefix: "0"
"256": 
  country_code: "256"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: UG
  name: Uganda
  international_dialing_prefix: "0"
"228": 
  country_code: "228"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: TG
  name: Togo
  international_dialing_prefix: "0"
"595": 
  country_code: "595"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: PY
  name: Paraguay
  international_dialing_prefix: "2"
"961": 
  country_code: "961"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: LB
  name: Lebanon
  international_dialing_prefix: "0"
"596": 
  country_code: "596"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: MQ
  name: Martinique
  international_dialing_prefix: "0"
"257": 
  country_code: "257"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: BI
  name: Burundi
  international_dialing_prefix: "0"
"229": 
  country_code: "229"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: BJ
  name: Benin
  international_dialing_prefix: "0"
"962": 
  country_code: "962"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: JO
  name: Jordan
  international_dialing_prefix: "0"
"963": 
  country_code: "963"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: SY
  name: Syrian Arab Republic
  international_dialing_prefix: "0"
"597": 
  country_code: "597"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: SR
  name: Suriname
  international_dialing_prefix: "0"
"680": 
  country_code: "680"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: PW
  name: Palau
  international_dialing_prefix: "0"
"258": 
  country_code: "258"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: MZ
  name: Mozambique
  international_dialing_prefix: "0"
"30": 
  country_code: "30"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: GR
  name: Greece
  international_dialing_prefix: "0"
"681": 
  country_code: "681"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: WF
  name: Wallis and Futuna
  international_dialing_prefix: "19"
"598": 
  country_code: "598"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: UY
  name: Uruguay
  international_dialing_prefix: "0"
  area_code: "2|42|4364|43[34567]|4452|44[3457]|454[24]|4567?|4586|46[234]|4675|47[237]|4779|9[13456789]"
"992": 
  country_code: "992"
  national_dialing_prefix: "8"
  char_2_code: "8"
  char_3_code: TJ
  name: Tajikistan
  international_dialing_prefix: "810"
"31": 
  country_code: "31"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: NL
  name: Netherlands
  international_dialing_prefix: "0"
  area_code: "6760|66|6|800|878|8[4578]|90[069]|1[035]|2[0346]|3[03568]|4[0356]|5[0358]|7\\d|11[134578]|16[124-8]|17[24]|18[0-467]|22[2-46-9]|25[125]|29[479]|31[3-8]|32[01]|34[1-8]|41[12368]|47[58]|48[15-8]|49[23579]|5[129][1-9]|54[134-8]|56[126]|57[0-3578]"
"850": 
  country_code: "850"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: KP
  name: Korea, Democratic People's Republic Of
  international_dialing_prefix: "0"
"964": 
  country_code: "964"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: IQ
  name: Iraq
  international_dialing_prefix: "0"
"370": 
  country_code: "370"
  national_dialing_prefix: "8"
  char_2_code: "8"
  char_3_code: LT
  name: Lithuania
  international_dialing_prefix: "0"
"993": 
  country_code: "993"
  national_dialing_prefix: "8"
  char_2_code: "8"
  char_3_code: TM
  name: Turkmenistan
  international_dialing_prefix: "810"
"599": 
  country_code: "599"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: AN
  name: Netherlands Antilles
  international_dialing_prefix: "0"
"32": 
  country_code: "32"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: BE
  name: Belgium
  international_dialing_prefix: "0"
  area_code: "800|90\\d|2|3|4|9|1[0-69]|5\\d|6[013-9]|7[01]|8[1-9]"
"965": 
  country_code: "965"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: KW
  name: Kuwait
  international_dialing_prefix: "0"
"371": 
  country_code: "371"
  national_dialing_prefix: "8"
  char_2_code: "8"
  char_3_code: LV
  name: Latvia
  international_dialing_prefix: "0"
"682": 
  country_code: "682"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: CK
  name: Cook Islands
  international_dialing_prefix: "0"
"60": 
  country_code: "60"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: MY
  name: Malaysia
  international_dialing_prefix: "0"
"966": 
  country_code: "966"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: SA
  name: Saudi Arabia
  international_dialing_prefix: "0"
"683": 
  country_code: "683"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: NU
  name: Niue
  international_dialing_prefix: "0"
"230": 
  country_code: "230"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: MU
  name: Mauritius
  international_dialing_prefix: "20"
"994": 
  country_code: "994"
  national_dialing_prefix: "8"
  char_2_code: "8"
  char_3_code: AZ
  name: Azerbaijan
  international_dialing_prefix: "810"
"852": 
  country_code: "852"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: HK
  name: Hong Kong
  international_dialing_prefix: "1"
"372": 
  country_code: "372"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: EE
  name: Estonia
  international_dialing_prefix: "0"
"61": 
  country_code: "61"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: AU
  name: Australia
  international_dialing_prefix: "11"
  area_code: "[234578]"
"880": 
  country_code: "880"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: BD
  name: Bangladesh
  international_dialing_prefix: "0"
"967": 
  country_code: "967"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: YE
  name: Yemen
  international_dialing_prefix: "0"
"90": 
  country_code: "90"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: TR
  name: Turkey
  international_dialing_prefix: "0"
"373": 
  country_code: "373"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: MD
  name: Moldova, Republic of
  international_dialing_prefix: "0"
"33": 
  country_code: "33"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: FR
  name: France
  international_dialing_prefix: "0"
  area_code: "[1-9]"
"995": 
  country_code: "995"
  national_dialing_prefix: 8*
  char_2_code: 8*
  char_3_code: GE
  name: Georgia
  international_dialing_prefix: "810"
"853": 
  country_code: "853"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: MO
  name: Macao
  international_dialing_prefix: "0"
"231": 
  country_code: "231"
  national_dialing_prefix: "22"
  char_2_code: "22"
  char_3_code: LR
  name: Liberia
  international_dialing_prefix: "0"
"62": 
  country_code: "62"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: ID
  name: Indonesia
  international_dialing_prefix: "1"
"260": 
  country_code: "260"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: ZM
  name: Zambia
  international_dialing_prefix: "0"
"34": 
  country_code: "34"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: ES
  name: Spain
  international_dialing_prefix: "0"
  area_code: "6[0-9][0-9]|7[1-9][0-9]|8[0-9][0-9]|9[0-9][0-9]"  
"232": 
  country_code: "232"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: SL
  name: Sierra Leone
  international_dialing_prefix: "0"
"685": 
  country_code: "685"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: WS
  name: Samoa
  international_dialing_prefix: "0"
"63": 
  country_code: "63"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: PH
  name: Philippines
  international_dialing_prefix: "0"
"968": 
  country_code: "968"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: OM
  name: Oman
  international_dialing_prefix: "0"
"996": 
  country_code: "996"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: KG
  name: Kyrgyzstan
  international_dialing_prefix: "0"
"374": 
  country_code: "374"
  national_dialing_prefix: "8"
  char_2_code: "8"
  char_3_code: AM
  name: Armenia
  international_dialing_prefix: "0"
"91": 
  country_code: "91"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: IN
  name: India
  international_dialing_prefix: "0"
"92": 
  country_code: "92"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: PK
  name: Pakistan
  international_dialing_prefix: "0"
"64": 
  country_code: "64"
  national_dialing_prefix: 0 (None fo
  char_2_code: 0 (None fo
  char_3_code: NZ
  name: New Zealand
  international_dialing_prefix: "0"
  area_code: "[1-9]"
"855": 
  country_code: "855"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: KH
  name: Cambodia
  international_dialing_prefix: "0"
"261": 
  country_code: "261"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: MG
  name: Madagascar
  international_dialing_prefix: "0"
"1": 
  country_code: "1"
  national_dialing_prefix: "1"
  char_2_code: "1"
  char_3_code: US
  name: United States
  international_dialing_prefix: "11"
"375": 
  country_code: "375"
  national_dialing_prefix: "8"
  char_2_code: "8"
  char_3_code: BY
  name: Belarus
  international_dialing_prefix: "810"
"233": 
  country_code: "233"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: GH
  name: Ghana
  international_dialing_prefix: "0"
"686": 
  country_code: "686"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: KI
  name: Kiribati
  international_dialing_prefix: "0"
"998": 
  country_code: "998"
  national_dialing_prefix: "8"
  char_2_code: "8"
  char_3_code: UZ
  name: Uzbekistan
  international_dialing_prefix: "810"
"65": 
  country_code: "65"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: SG
  name: Singapore
  international_dialing_prefix: "1"
"290": 
  country_code: "290"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: SH
  name: Saint Helena
  international_dialing_prefix: "0"
"262": 
  country_code: "262"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: RE
  name: "R\xC3\xA9union"
  international_dialing_prefix: "0"
"234": 
  country_code: "234"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: NG
  name: Nigeria
  international_dialing_prefix: "9"
"687": 
  country_code: "687"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: NC
  name: New Caledonia
  international_dialing_prefix: "0"
"856": 
  country_code: "856"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: LA
  name: Lao People's Democratic Republic
  international_dialing_prefix: "0"
"93": 
  country_code: "93"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: AF
  name: Afghanistan
  international_dialing_prefix: "0"
"376": 
  country_code: "376"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: AD
  name: Andorra
  international_dialing_prefix: "0"
"36": 
  country_code: "36"
  national_dialing_prefix: "6"
  char_2_code: "6"
  char_3_code: HU
  name: Hungary
  international_dialing_prefix: "0"
  area_code: "1|[2-9]\\d"
"263": 
  country_code: "263"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: ZW
  name: Zimbabwe
  international_dialing_prefix: "0"
"688": 
  country_code: "688"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: TV
  name: Tuvalu
  international_dialing_prefix: "0"
"94": 
  country_code: "94"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: LK
  name: Sri Lanka
  international_dialing_prefix: "0"
"377": 
  country_code: "377"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: MC
  name: Monaco
  international_dialing_prefix: "0"
"235": 
  country_code: "235"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: TD
  name: Chad
  international_dialing_prefix: "15"
"291": 
  country_code: "291"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: ER
  name: Eritrea
  international_dialing_prefix: "0"
"66": 
  country_code: "66"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: TH
  name: Thailand
  international_dialing_prefix: "1"
"886": 
  country_code: "886"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: TW
  name: Taiwan, Province Of China
  international_dialing_prefix: "2"
"378": 
  country_code: "378"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: SM
  name: San Marino
  international_dialing_prefix: "0"
"264": 
  country_code: "264"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: NA
  name: Namibia
  international_dialing_prefix: "0"
"95": 
  country_code: "95"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: MM
  name: Myanmar
  international_dialing_prefix: "0"
"236": 
  country_code: "236"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: CF
  name: Central African Republic
  international_dialing_prefix: "0"
"689": 
  country_code: "689"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: PF
  name: French Polynesia
  international_dialing_prefix: "0"
"970": 
  country_code: "970"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: PS
  name: Palestinian Territory, Occupied
  international_dialing_prefix: "0"
"237": 
  country_code: "237"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: CM
  name: Cameroon
  international_dialing_prefix: "0"
"39": 
  country_code: "39"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: IT
  name: Italy
  international_dialing_prefix: "0"
"265": 
  country_code: "265"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: MW
  name: Malawi
  international_dialing_prefix: "0"
"971": 
  country_code: "971"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: AE
  name: United Arab Emirates
  international_dialing_prefix: "0"
"238": 
  country_code: "238"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: CV
  name: Cape Verde
  international_dialing_prefix: "0"
"266": 
  country_code: "266"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: LS
  name: Lesotho
  international_dialing_prefix: "0"
"239": 
  country_code: "239"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: ST
  name: Sao Tome and Principe
  international_dialing_prefix: "0"
"7": 
  country_code: "7"
  national_dialing_prefix: "8"
  char_2_code: "8"
  char_3_code: RU
  name: Russian Federation
  international_dialing_prefix: "810"
"98": 
  country_code: "98"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: IR
  name: Iran, Islamic Republic Of
  international_dialing_prefix: "0"
"972": 
  country_code: "972"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: IL
  name: Israel
  international_dialing_prefix: "0"
"350": 
  country_code: "350"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: GI
  name: Gibraltar
  international_dialing_prefix: "0"
"267": 
  country_code: "267"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: BW
  name: Botswana
  international_dialing_prefix: "0"
"690": 
  country_code: "690"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: TK
  name: Tokelau
  international_dialing_prefix: "0"
"268": 
  country_code: "268"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: SZ
  name: Swaziland
  international_dialing_prefix: "0"
"40": 
  country_code: "40"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: RO
  name: Romania
  international_dialing_prefix: "0"
"351": 
  country_code: "351"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: PT
  name: Portugal
  international_dialing_prefix: "0"
  area_code: "2[12]|2[3-9][1-9]|70[78]|80[089]|9[136]|92[1-9]"
"973": 
  country_code: "973"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: BH
  name: Bahrain
  international_dialing_prefix: "0"
"380": 
  country_code: "380"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: UA
  name: Ukraine
  international_dialing_prefix: "00"
  area_code: "[1-9]\\d"
"41": 
  country_code: "41"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: CH
  name: Switzerland
  international_dialing_prefix: "0"
"974": 
  country_code: "974"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: QA
  name: Qatar
  international_dialing_prefix: "0"
"691": 
  country_code: "691"
  national_dialing_prefix: "1"
  char_2_code: "1"
  char_3_code: FM
  name: Micronesia, Federated States Of
  international_dialing_prefix: "11"
"297": 
  country_code: "297"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: AW
  name: Aruba
  international_dialing_prefix: "0"
"352": 
  country_code: "352"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: LU
  name: Luxembourg
  international_dialing_prefix: "0"
"269": 
  country_code: "269"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: YT
  name: Mayotte
  international_dialing_prefix: "0"
"381": 
  country_code: "381"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: RS
  name: Serbia
  international_dialing_prefix: "99"
  area_code: "[1-9]\\d"
"975": 
  country_code: "975"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: BT
  name: Bhutan
  international_dialing_prefix: "0"
"298": 
  country_code: "298"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: FO
  name: Faroe Islands
  international_dialing_prefix: "0"
"353": 
  country_code: "353"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: IE
  name: Ireland
  international_dialing_prefix: "0"
  area_code: "1|[2,4-7,9][0-9]|8[0,3-9]|822|818"  
"692": 
  country_code: "692"
  national_dialing_prefix: "1"
  char_2_code: "1"
  char_3_code: MH
  name: Marshall Islands
  international_dialing_prefix: "0"
"212": 
  country_code: "212"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: MA
  name: Morocco
  international_dialing_prefix: "0"
"382": 
  country_code: "382"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: ME
  name: Montenegro
  international_dialing_prefix: "99"
  area_code: "[2-6][0-9]"
"976": 
  country_code: "976"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: MN
  name: Mongolia
  international_dialing_prefix: "1"
"240": 
  country_code: "240"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: GQ
  name: Equatorial Guinea
  international_dialing_prefix: "0"
"299": 
  country_code: "299"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: GL
  name: Greenland
  international_dialing_prefix: "9"
"354": 
  country_code: "354"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: IS
  name: Iceland
  international_dialing_prefix: "0"
"43": 
  country_code: "43"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: AT
  name: Austria
  international_dialing_prefix: "0"
"977": 
  country_code: "977"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: NP
  name: Nepal
  international_dialing_prefix: "0"
"241": 
  country_code: "241"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: GA
  name: Gabon
  international_dialing_prefix: "0"
"355": 
  country_code: "355"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: AL
  name: Albania
  international_dialing_prefix: "0"
"213": 
  country_code: "213"
  national_dialing_prefix: "7"
  char_2_code: "7"
  char_3_code: DZ
  name: Algeria
  international_dialing_prefix: "0"
"44": 
  country_code: "44"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: GB
  name: United Kingdom
  international_dialing_prefix: "0"
  area_code: "2[03489]|11[3-8]|1[2-69]1|1[2-9][0-9]{2}|70|7[0-9]{3}|[8|9][0-9]{2}|3[0-9]{2}"
"242": 
  country_code: "242"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: CG
  name: Congo
  international_dialing_prefix: "0"
"356": 
  country_code: "356"
  national_dialing_prefix: "21"
  char_2_code: "21"
  char_3_code: MT
  name: Malta
  international_dialing_prefix: "0"
"357": 
  country_code: "357"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: CY
  name: Cyprus
  international_dialing_prefix: "0"
"45": 
  country_code: "45"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: DK
  name: Denmark
  international_dialing_prefix: "0"
"385": 
  country_code: "385"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: HR
  name: Croatia
  international_dialing_prefix: "0"
  area_code: "1|[2-9]\\d"
"243": 
  country_code: "243"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: CD
  name: Congo, The Democratic Republic Of The
  international_dialing_prefix: "0"
"216": 
  country_code: "216"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: TN
  name: Tunisia
  international_dialing_prefix: "0"
"46": 
  country_code: "46"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: SE
  name: Sweden
  international_dialing_prefix: "0"
  area_code: "900|1[013689]|2[0136]|3[1356]|4[0246]|54|6[03]|7[01236]|8|9[09]|1[2457]\\d|2[2457-9]\\d|3[0247-9]\\d|4[1357-9]\\d|5[0-35-9]\\d|6[124-9]\\d|74\\d|9[1-8]\\d"
"386": 
  country_code: "386"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: SI
  name: Slovenia
  international_dialing_prefix: "0"
  area_code: "3[01]|4[01]|51|7[01]|64|59|1|2|3|4|5|6|7"
"358": 
  country_code: "358"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: FI
  name: Finland
  international_dialing_prefix: "0"
"244": 
  country_code: "244"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: AO
  name: Angola
  international_dialing_prefix: "0"
"47": 
  country_code: "47"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: SJ
  name: Svalbard And Jan Mayen
  international_dialing_prefix: "0"
"359": 
  country_code: "359"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: BG
  name: Bulgaria
  international_dialing_prefix: "0"
"387": 
  country_code: "387"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: BA
  name: Bosnia and Herzegovina
  international_dialing_prefix: "0"
  area_code: "6|[0-57-9]\\d"
"245": 
  country_code: "245"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: GW
  name: Guinea-Bissau
  international_dialing_prefix: "0"
"48": 
  country_code: "48"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: PL
  name: Poland
  international_dialing_prefix: "0"
"218": 
  country_code: "218"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: LY
  name: Libyan Arab Jamahiriya
  international_dialing_prefix: "0"
"49": 
  country_code: "49"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: DE
  name: Germany
  international_dialing_prefix: "0"
  area_code: "[0-9]{3}"
"389": 
  country_code: "389"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: MK
  name: Macedonia, the Former Yugoslav Republic Of
  international_dialing_prefix: "0"
"670": 
  country_code: "670"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: TL
  name: Timor-Leste
  international_dialing_prefix: None
"248": 
  country_code: "248"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: SC
  name: Seychelles
  international_dialing_prefix: "0"
"20": 
  country_code: "20"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: EG
  name: Egypt
  international_dialing_prefix: "0"
"500": 
  country_code: "500"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: FK
  name: Falkland Islands (Malvinas)
  international_dialing_prefix: "0"
"249": 
  country_code: "249"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: SD
  name: Sudan
  international_dialing_prefix: "0"
"501": 
  country_code: "501"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: BZ
  name: Belize
  international_dialing_prefix: "0"
"672": 
  country_code: "672"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: NF
  name: Norfolk Island
  international_dialing_prefix: "0"
"502": 
  country_code: "502"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: GT
  name: Guatemala
  international_dialing_prefix: "0"
"51": 
  country_code: "51"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: PE
  name: Peru
  international_dialing_prefix: "0"
"220": 
  country_code: "220"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: GM
  name: Gambia
  international_dialing_prefix: "0"
"673": 
  country_code: "673"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: BN
  name: Brunei Darussalam
  international_dialing_prefix: "0"
"503": 
  country_code: "503"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: SV
  name: El Salvador
  international_dialing_prefix: "0"
"221": 
  country_code: "221"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: SN
  name: Senegal
  international_dialing_prefix: "0"
"674": 
  country_code: "674"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: NR
  name: Nauru
  international_dialing_prefix: "0"
"52": 
  country_code: "52"
  national_dialing_prefix: "1"
  char_2_code: "1"
  char_3_code: MX
  name: Mexico
  international_dialing_prefix: "0"
"504": 
  country_code: "504"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: HN
  name: Honduras
  international_dialing_prefix: "0"
"250": 
  country_code: "250"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: RW
  name: Rwanda
  international_dialing_prefix: "0"
"872": 
  country_code: "872"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: PN
  name: Pitcairn
  international_dialing_prefix: "0"
"675": 
  country_code: "675"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: PG
  name: Papua New Guinea
  international_dialing_prefix: "5"
"505": 
  country_code: "505"
  national_dialing_prefix: None
  char_2_code: None
  char_3_code: NI
  name: Nicaragua
  international_dialing_prefix: "0"
"222": 
  country_code: "222"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: MR
  name: Mauritania
  international_dialing_prefix: "0"
"53": 
  country_code: "53"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: CU
  name: Cuba
  international_dialing_prefix: "119"
"81": 
  country_code: "81"
  national_dialing_prefix: "0"
  char_2_code: "0"
  char_3_code: JP
  name: Japan
  international_dialing_prefix: "10"
`

// Country holds country information.
type Country struct {
	Number                     string `yaml:"number"`
	Name                       string `yaml:"name"`
	CountryCode                string `yaml:"country_code"`
	Char2Code                  string `yaml:"char_2_code"`
	Char3Code                  string `yaml:"char_3_code"`
	AreaCode                   string `yaml:"area_code"`
	MaxNumLength               string `yaml:"max_num_length"`
	NationalDialingPrefix      string `yaml:"national_dialing_prefix"`
	InternationalDialingPrefix string `yaml:"international_dialing_prefix"`
	Extension                  string `yaml:"extension"`
	N1Length                   string
}

var Countries map[string]Country

func init() {
	Countries = loadCountries()
}

// FindByCountryCode finds country by dialing code.
func FindByCountryCode(code string) *Country {
	if country, found := Countries[code]; found {
		return &country
	}
	return nil
}

// FindByCountryIsoCode finds country by ISO code (case insensitive).
func FindByCountryIsoCode(isoCode string) (c *Country) {
	for _, v := range Countries {
		if strings.EqualFold(isoCode, v.Char3Code) {
			return &v
		}
	}
	return nil
}

func (c *Country) CountryCodeRegexp() *regexp.Regexp {
	exp := fmt.Sprintf("^[+]%s", c.CountryCode)
	re, _ := regexp.Compile(exp)

	return re
}

func (c *Country) Formats() (*regexp.Regexp, *regexp.Regexp) {
	numberRegex := fmt.Sprintf("([0-9]{1,%s})$", c.MaxNumLength)
	short := regexp.MustCompile(fmt.Sprintf("^0?(%s)%s", c.AreaCode, numberRegex))
	reallyShort := regexp.MustCompile(fmt.Sprintf("^%s", numberRegex))
	return short, reallyShort
}

func (c *Country) DetectFormat(stringWithNumber string) string {
	sh, real := c.Formats()

	var arr []string
	if sh.MatchString(stringWithNumber) {
		arr = append(arr, "short")
	}

	if real.MatchString(stringWithNumber) {
		arr = append(arr, "really_short")
	}

	switch {
	case len(arr) > 1:
		return "really_short"
	case len(arr) == 0:
		return "short"
	default:
		return arr[0]
	}
}

func loadCountries() map[string]Country {
	var c map[string]Country

	err := yaml.Unmarshal([]byte(countries), &c)
	if err != nil {
		panic(err)
	}
	return c
}

func detectCountry(s, defaultCode string) *Country {
	for k, v := range Countries {
		re := fmt.Sprintf("^[+]%s", k)
		matched, _ := regexp.MatchString(re, s)
		if matched {
			return &v
		}
	}

	_c := Country{}
	if _c.CountryCode == "" {
		_c = Countries[defaultCode]
	}
	return &_c
}
