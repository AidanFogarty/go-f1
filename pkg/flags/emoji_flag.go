package flags

var flagMap = map[string]string{
	"Aruban":                                 "🇦🇼",
	"Afghan":                                 "🇦🇫",
	"Angolan":                                "🇦🇴",
	"Anguillian":                             "🇦🇮",
	"Ålandish":                               "🇦🇽",
	"Albanian":                               "🇦🇱",
	"Andorran":                               "🇦🇩",
	"Emirati":                                "🇦🇪",
	"Argentine":                              "🇦🇷",
	"Armenian":                               "🇦🇲",
	"American Samoan":                        "🇦🇸",
	"Antarctican":                            "🇦🇶",
	"Antiguan, Barbudan":                     "🇦🇬",
	"Australian":                             "🇦🇺",
	"Austrian":                               "🇦🇹",
	"Azerbaijani":                            "🇦🇿",
	"Burundian":                              "🇧🇮",
	"Belgian":                                "🇧🇪",
	"Beninese":                               "🇧🇯",
	"Burkinabe":                              "🇧🇫",
	"Bangladeshi":                            "🇧🇩",
	"Bulgarian":                              "🇧🇬",
	"Bahraini":                               "🇧🇭",
	"Bahamian":                               "🇧🇸",
	"Bosnian, Herzegovinian":                 "🇧🇦",
	"Saint Barthélemy Islander":              "🇧🇱",
	"Saint Helenian":                         "🇸🇭",
	"Belarusian":                             "🇧🇾",
	"Belizean":                               "🇧🇿",
	"Bermudian":                              "🇧🇲",
	"Bolivian":                               "🇧🇴",
	"Brazilian":                              "🇧🇷",
	"Barbadian":                              "🇧🇧",
	"Bruneian":                               "🇧🇳",
	"Bhutanese":                              "🇧🇹",
	"Motswana":                               "🇧🇼",
	"Central African":                        "🇨🇫",
	"Canadian":                               "🇨🇦",
	"Cocos Islander":                         "🇨🇨",
	"Swiss":                                  "🇨🇭",
	"Chilean":                                "🇨🇱",
	"Chinese":                                "🇨🇳",
	"Ivorian":                                "🇨🇮",
	"Cameroonian":                            "🇨🇲",
	"Congolese":                              "🇨🇬",
	"Cook Islander":                          "🇨🇰",
	"Colombian":                              "🇨🇴",
	"Comoran":                                "🇰🇲",
	"Cape Verdian":                           "🇨🇻",
	"Costa Rican":                            "🇨🇷",
	"Cuban":                                  "🇨🇺",
	"Curaçaoan":                              "🇨🇼",
	"Christmas Islander":                     "🇨🇽",
	"Caymanian":                              "🇰🇾",
	"Cypriot":                                "🇨🇾",
	"Czech":                                  "🇨🇿",
	"German":                                 "🇩🇪",
	"Djibouti":                               "🇩🇯",
	"Danish":                                 "🇩🇰",
	"Dominican":                              "🇩🇴",
	"Algerian":                               "🇩🇿",
	"Ecuadorean":                             "🇪🇨",
	"Egyptian":                               "🇪🇬",
	"Eritrean":                               "🇪🇷",
	"Sahrawi":                                "🇪🇭",
	"Spanish":                                "🇪🇸",
	"Estonian":                               "🇪🇪",
	"Ethiopian":                              "🇪🇹",
	"Finnish":                                "🇫🇮",
	"Fijian":                                 "🇫🇯",
	"Falkland Islander":                      "🇫🇰",
	"French":                                 "🇫🇷",
	"Faroese":                                "🇫🇴",
	"Micronesian":                            "🇫🇲",
	"Gabonese":                               "🇬🇦",
	"British":                                "🇬🇧",
	"Georgian":                               "🇬🇪",
	"Ghanaian":                               "🇬🇭",
	"Gibraltar":                              "🇬🇮",
	"Guinean":                                "🇬🇳",
	"Guadeloupian":                           "🇬🇵",
	"Gambian":                                "🇬🇲",
	"Guinea-Bissauan":                        "🇬🇼",
	"Equatorial Guinean":                     "🇬🇶",
	"Greek":                                  "🇬🇷",
	"Grenadian":                              "🇬🇩",
	"Greenlandic":                            "🇬🇱",
	"Guatemalan":                             "🇬🇹",
	"Guianan":                                "🇬🇫",
	"Guamanian":                              "🇬🇺",
	"Guyanese":                               "🇬🇾",
	"Hong Konger":                            "🇭🇰",
	"Heard and McDonald Islander":            "🇭🇲",
	"Honduran":                               "🇭🇳",
	"Croatian":                               "🇭🇷",
	"Haitian":                                "🇭🇹",
	"Hungarian":                              "🇭🇺",
	"Indonesian":                             "🇮🇩",
	"Manx":                                   "🇮🇲",
	"Indian":                                 "🇮🇳",
	"Irish":                                  "🇮🇪",
	"Iranian":                                "🇮🇷",
	"Iraqi":                                  "🇮🇶",
	"Icelander":                              "🇮🇸",
	"Israeli":                                "🇮🇱",
	"Italian":                                "🇮🇹",
	"Jamaican":                               "🇯🇲",
	"Channel Islander":                       "🇯🇪",
	"Jordanian":                              "🇯🇴",
	"Japanese":                               "🇯🇵",
	"Kazakhstani":                            "🇰🇿",
	"Kenyan":                                 "🇰🇪",
	"Kirghiz":                                "🇰🇬",
	"Cambodian":                              "🇰🇭",
	"I-Kiribati":                             "🇰🇮",
	"Kittitian or Nevisian":                  "🇰🇳",
	"South Korean":                           "🇰🇷",
	"Kosovar":                                "🇽🇰",
	"Kuwaiti":                                "🇰🇼",
	"Laotian":                                "🇱🇦",
	"Lebanese":                               "🇱🇧",
	"Liberian":                               "🇱🇷",
	"Libyan":                                 "🇱🇾",
	"Saint Lucian":                           "🇱🇨",
	"Liechtensteiner":                        "🇱🇮",
	"Sri Lankan":                             "🇱🇰",
	"Mosotho":                                "🇱🇸",
	"Lithuanian":                             "🇱🇹",
	"Luxembourger":                           "🇱🇺",
	"Latvian":                                "🇱🇻",
	"Macanese":                               "🇲🇴",
	"Saint Martin Islander":                  "🇲🇫",
	"Moroccan":                               "🇲🇦",
	"Monegasque":                             "🇲🇨",
	"Moldovan":                               "🇲🇩",
	"Malagasy":                               "🇲🇬",
	"Maldivan":                               "🇲🇻",
	"Mexican":                                "🇲🇽",
	"Marshallese":                            "🇲🇭",
	"Macedonian":                             "🇲🇰",
	"Malian":                                 "🇲🇱",
	"Maltese":                                "🇲🇹",
	"Burmese":                                "🇲🇲",
	"Montenegrin":                            "🇲🇪",
	"Mongolian":                              "🇲🇳",
	"Mozambican":                             "🇲🇿",
	"Mauritanian":                            "🇲🇷",
	"Montserratian":                          "🇲🇸",
	"Martinican":                             "🇲🇶",
	"Mauritian":                              "🇲🇺",
	"Malawian":                               "🇲🇼",
	"Malaysian":                              "🇲🇾",
	"Mahoran":                                "🇾🇹",
	"Namibian":                               "🇳🇦",
	"New Caledonian":                         "🇳🇨",
	"Nigerien":                               "🇳🇪",
	"Norfolk Islander":                       "🇳🇫",
	"Nigerian":                               "🇳🇬",
	"Nicaraguan":                             "🇳🇮",
	"Niuean":                                 "🇳🇺",
	"Dutch":                                  "🇳🇱",
	"Nepalese":                               "🇳🇵",
	"Nauruan":                                "🇳🇷",
	"New Zealander":                          "🇳🇿",
	"Omani":                                  "🇴🇲",
	"Pakistani":                              "🇵🇰",
	"Panamanian":                             "🇵🇦",
	"Pitcairn Islander":                      "🇵🇳",
	"Peruvian":                               "🇵🇪",
	"Filipino":                               "🇵🇭",
	"Palauan":                                "🇵🇼",
	"Papua New Guinean":                      "🇵🇬",
	"Polish":                                 "🇵🇱",
	"Puerto Rican":                           "🇵🇷",
	"North Korean":                           "🇰🇵",
	"Portuguese":                             "🇵🇹",
	"Paraguayan":                             "🇵🇾",
	"Palestinian":                            "🇵🇸",
	"French Polynesian":                      "🇵🇫",
	"Qatari":                                 "🇶🇦",
	"Réunionese":                             "🇷🇪",
	"Romanian":                               "🇷🇴",
	"Russian":                                "🇷🇺",
	"Rwandan":                                "🇷🇼",
	"Saudi Arabian":                          "🇸🇦",
	"Sudanese":                               "🇸🇩",
	"Senegalese":                             "🇸🇳",
	"Singaporean":                            "🇸🇬",
	"South Georgian South Sandwich Islander": "🇬🇸",
	"Norwegian":                              "🇸🇯",
	"Solomon Islander":                       "🇸🇧",
	"Sierra Leonean":                         "🇸🇱",
	"Salvadoran":                             "🇸🇻",
	"Sammarinese":                            "🇸🇲",
	"Somali":                                 "🇸🇴",
	"Saint-Pierrais, Miquelonnais":           "🇵🇲",
	"Serbian":                                "🇷🇸",
	"South Sudanese":                         "🇸🇸",
	"Sao Tomean":                             "🇸🇹",
	"Surinamer":                              "🇸🇷",
	"Slovak":                                 "🇸🇰",
	"Slovene":                                "🇸🇮",
	"Swedish":                                "🇸🇪",
	"Swazi":                                  "🇸🇿",
	"St. Maartener":                          "🇸🇽",
	"Seychellois":                            "🇸🇨",
	"Syrian":                                 "🇸🇾",
	"Turks and Caicos Islander":              "🇹🇨",
	"Chadian":                                "🇹🇩",
	"Togolese":                               "🇹🇬",
	"Thai":                                   "🇹🇭",
	"Tadzhik":                                "🇹🇯",
	"Tokelauan":                              "🇹🇰",
	"Turkmen":                                "🇹🇲",
	"East Timorese":                          "🇹🇱",
	"Tongan":                                 "🇹🇴",
	"Trinidadian":                            "🇹🇹",
	"Tunisian":                               "🇹🇳",
	"Turkish":                                "🇹🇷",
	"Tuvaluan":                               "🇹🇻",
	"Taiwanese":                              "🇹🇼",
	"Tanzanian":                              "🇹🇿",
	"Ugandan":                                "🇺🇬",
	"Ukrainian":                              "🇺🇦",
	"American Islander":                      "🇺🇲",
	"Uruguayan":                              "🇺🇾",
	"American":                               "🇺🇸",
	"Uzbekistani":                            "🇺🇿",
	"Vatican":                                "🇻🇦",
	"Saint Vincentian":                       "🇻🇨",
	"Venezuelan":                             "🇻🇪",
	"Virgin Islander":                        "🇻🇮",
	"Vietnamese":                             "🇻🇳",
	"Ni-Vanuatu":                             "🇻🇺",
	"Wallis and Futuna Islander":             "🇼🇫",
	"Samoan":                                 "🇼🇸",
	"Yemeni":                                 "🇾🇪",
	"South African":                          "🇿🇦",
	"Zambian":                                "🇿🇲",
	"Zimbabwean":                             "🇿🇼",
}

// GetFlag returns an emoji flag for a given nationality.
func GetFlag(nationality string) string {
	if flag, ok := flagMap[nationality]; ok {
		return flag
	}
	return "🌎"
}
