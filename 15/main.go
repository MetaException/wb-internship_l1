package main

var justString string

func createHugeString(size int) string {
	return string(make([]byte, size))
}

/*
	Тк. v[:100] вернёт слайс, что будет указывать на исходную строку v, то после завершения функции someFunc строка v не будет удалена.
	И, при дальнешей работе, указатель будет держать всю исходную строку в памяти
	Поэтому нужно создать новую строку, которая будет указывать на новую область памяти с нужной подстрокой, а исходная строка будет удалена после выполнения someFunc()
*/

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string(v[:100])
}

func main() {
	someFunc()
}
