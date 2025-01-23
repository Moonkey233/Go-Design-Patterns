package abstractfactory

//goland:noinspection GoTestName
func ExampleAbstractFactory() {
	var factory SaveArticle
	factory = &SaveRedis{}
	CreateAndSave(factory)
	factory = &SaveMysql{}
	CreateAndSave(factory)

	// Output:
	// Redis Save Prose
	// Redis Save Ancient Poetry
	// Mysql Save Prose
	// Mysql Save Ancient Poetry
}
