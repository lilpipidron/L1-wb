package tasks

/*var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}*/
// имеем несколько потенкциальных проблемы
// 1) Строка может быть меньше 100 симвлов
// 2) строка justString хоть и ссылается на подстроку v[:100], но она в любом случае хранит ссылку на весь объект v и имеем утечку памяти
// 3) если у нас не только латиница используется, то у нас может получится неккоренткный срез

// используем проверку на количество символов и преобразование в срез рун, для коректной обработки не только латиницы
// https://go.dev/play/p/h5H4BoaQCCX - пример, когда без среза и с срезом рун
var justString string

// просто заглушка для удобства
func createHugeString(size int) string {
	return ""
}

func someFunc() {
	v := createHugeString(1 << 10)

	if len([]rune(v)) >= 100 {
		justString = string([]rune(v)[:100])
	} else {
		justString = string([]rune(v)[:len([]rune(v))])
	}
}
