package main

import "fmt"

const (
	defaultName = "World"

	spanish = "Spanish"
	french = "French"
	japanese = "Japanese"

	helloEnglish = "Hello"
	helloSpanish = "Hola"
	helloFrench = "Bonjour"
	helloJapanese = "Kon'nichiwa"
	commaSpace = ", "
)
 

func Hello(name, language string) string {
	if name == "" {
		name = defaultName
	}
return helloIn(language) + commaSpace + name
}

func helloIn(language string) (hello string){
	switch language {
	case french:
		hello = helloFrench
	case spanish:
		hello = helloSpanish
	case japanese:
		hello = helloJapanese
	default:
		hello = helloEnglish
	}

	return hello
}



func main() {
	fmt.Println(Hello("World", "English"))
}
