package main

// Returns a gretting to a name in the given language
func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := grettingPrefix(lang)

	return prefix + ", " + name
}

func grettingPrefix(lang string) (prefix string) {
	switch lang {
	case "French":
		prefix = "Bonjour"
	case "Spanish":
		prefix = "Hola"
	case "Japanese":
		prefix = "こんにちは"
	default:
		prefix = "Hello"
	}
	return
}
