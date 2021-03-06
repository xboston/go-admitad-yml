package admitadyml

var (
	// AgeMap - allowed ages
	AgeMap = map[string]bool{
		"0":  true,
		"6":  true,
		"12": true,
		"16": true,
		"18": true,
	}

	// MonthMap - allowed month
	MonthMap = map[string]bool{
		"0":  true,
		"1":  true,
		"2":  true,
		"3":  true,
		"4":  true,
		"5":  true,
		"6":  true,
		"7":  true,
		"8":  true,
		"9":  true,
		"10": true,
		"11": true,
		"12": true,
	}

	// CountriesMap - allowed countries
	CountriesMap = map[string]bool{
		"Австралия":   true,
		"Австрия":     true,
		"Азербайджан": true,
		"Албания":     true,
		"Алжир":       true,
		"Американские Виргинские острова": true,
		"Ангилья":                         true,
		"Ангола":                          true,
		"Андорра":                         true,
		"Антигуа и Барбуда":               true,
		"Аргентина":                       true,
		"Армения":                         true,
		"Аруба":                           true,
		"Афганистан":                      true,
		"Багамские острова":               true,
		"Бангладеш":                       true,
		"Барбадос":                        true,
		"Бахрейн":                         true,
		"Беларусь":                        true,
		"Белиз":                           true,
		"Бельгия":                         true,
		"Бенин":                           true,
		"Бермудские Острова":              true,
		"Болгария":                        true,
		"Боливия":                         true,
		"Босния и Герцеговина":            true,
		"Ботсвана":                        true,
		"Бразилия":                        true,
		"Британские Виргинские острова":   true,
		"Бруней":          true,
		"Буркина-Фасо":    true,
		"Бурунди":         true,
		"Бутан":           true,
		"Вануату":         true,
		"Ватикан":         true,
		"Великобритания":  true,
		"Венгрия":         true,
		"Венесуэла":       true,
		"Восточный Тимор": true,
		"Вьетнам":         true,
		"Габон":           true,
		"Гайана":          true,
		"Гаити":           true,
		"Гамбия":          true,
		"Гана":            true,
		"Гваделупа":       true,
		"Гватемала":       true,
		"Гвинея":          true,
		"Гвинея-Бисау":    true,
		"Германия":        true,
		"Гибралтар":       true,
		"Гондурас":        true,
		"Гонконг":         true,
		"Гренада":         true,
		"Гренландия":      true,
		"Греция":          true,
		"Грузия":          true,
		"Дания":           true,
		"Демократическая Республика Конго": true,
		"Джибути":                       true,
		"Доминика":                      true,
		"Доминиканская Республика":      true,
		"Египет":                        true,
		"Замбия":                        true,
		"Западная Сахара":               true,
		"Зимбабве":                      true,
		"Йемен":                         true,
		"Израиль":                       true,
		"Индия":                         true,
		"Индонезия":                     true,
		"Иордания":                      true,
		"Ирак":                          true,
		"Иран":                          true,
		"Ирландия":                      true,
		"Исландия":                      true,
		"Испания":                       true,
		"Италия":                        true,
		"Кабо-Верде":                    true,
		"Казахстан":                     true,
		"Каймановы острова":             true,
		"Камбоджа":                      true,
		"Камерун":                       true,
		"Канада":                        true,
		"Катар":                         true,
		"Кения":                         true,
		"Кипр":                          true,
		"Киргизия":                      true,
		"Кирибати":                      true,
		"Китай":                         true,
		"Колумбия":                      true,
		"Коморские острова":             true,
		"Коста-Рика":                    true,
		"Кот-д’Ивуар":                   true,
		"Куба":                          true,
		"Кувейт":                        true,
		"Лаос":                          true,
		"Латвия":                        true,
		"Лесото":                        true,
		"Либерия":                       true,
		"Ливан":                         true,
		"Ливия":                         true,
		"Литва":                         true,
		"Лихтенштейн":                   true,
		"Люксембург":                    true,
		"Маврикий":                      true,
		"Мавритания":                    true,
		"Мадагаскар":                    true,
		"Майотта":                       true,
		"Макао":                         true,
		"Македония":                     true,
		"Малави":                        true,
		"Малайзия":                      true,
		"Мали":                          true,
		"Мальдивы":                      true,
		"Мальта":                        true,
		"Марокко":                       true,
		"Маршалловы острова":            true,
		"Мексика":                       true,
		"Мозамбик":                      true,
		"Молдова":                       true,
		"Монако":                        true,
		"Монголия":                      true,
		"Мьянма":                        true,
		"Намибия":                       true,
		"Науру":                         true,
		"Непал":                         true,
		"Нигер":                         true,
		"Нигерия":                       true,
		"Нидерланды":                    true,
		"Никарагуа":                     true,
		"Новая Зеландия":                true,
		"Новая Каледония":               true,
		"Норвегия":                      true,
		"Объединённые Арабские Эмираты": true,
		"Оман":                          true,
		"Острова Кука":                  true,
		"Пакистан":                      true,
		"Палау":                         true,
		"Панама":                        true,
		"Папуа - Новая Гвинея":          true,
		"Парагвай":                      true,
		"Перу":                          true,
		"Польша":                        true,
		"Португалия":                    true,
		"Республика Конго":              true,
		"Реюньон":                       true,
		"Россия":                        true,
		"Руанда":                        true,
		"Румыния":                       true,
		"Самоа":                         true,
		"Сан-Марино":                    true,
		"Сан-Томе и Принсипи":           true,
		"Саудовская Аравия":             true,
		"Свазиленд":                     true,
		"Северная Корея":                true,
		"Сейшельские острова":           true,
		"Сенегал":                       true,
		"Сент-Винсент и Гренадины":      true,
		"Сент-Китс и Невис":             true,
		"Сент-Люсия":                    true,
		"Сербия":                        true,
		"Сингапур":                      true,
		"Сирия":                         true,
		"Словакия":                      true,
		"Словения":                      true,
		"Сомали":                        true,
		"Судан":                         true,
		"Суринам":                       true,
		"США":                           true,
		"Сьерра-Леоне":                  true,
		"Таджикистан":                   true,
		"Таиланд":                       true,
		"Танзания":                      true,
		"Тёркс и Кайкос":                true,
		"Того":                          true,
		"Тонга":                         true,
		"Тринидад и Тобаго":             true,
		"Тувалу":                        true,
		"Тунис":                         true,
		"Туркмения":                     true,
		"Турция":                        true,
		"Уганда":                        true,
		"Узбекистан":                    true,
		"Украина":                       true,
		"Уругвай":                       true,
		"Федеративные Штаты Микронезии": true,
		"Фиджи":                             true,
		"Филиппины":                         true,
		"Финляндия":                         true,
		"Франция":                           true,
		"Французская Гвиана":                true,
		"Французская Полинезия":             true,
		"Хорватия":                          true,
		"Центрально-Африканская Республика": true,
		"Чад":                   true,
		"Черногория":            true,
		"Чехия":                 true,
		"Чили":                  true,
		"Швейцария":             true,
		"Швеция":                true,
		"Шри-Ланка":             true,
		"Эквадор":               true,
		"Экваториальная Гвинея": true,
		"Эритрея":               true,
		"Эстония":               true,
		"Эфиопия":               true,
		"ЮАР":                   true,
		"Южная Корея":           true,
		"Ямайка":                true,
		"Япония":                true,
	}
)
