package display

// NOTE: we can't use !+..!- comments to excerpt these tests
// into the book because it defeats the Example mechanism,
// which requires the // Output comment to be at the end
// of the function.

func Example_arrayKey() {
	arr := [2]string{"apple", "banana"} //配列定義
	mapk := map[[2]string]int{arr: 100} //配列をキーに持つmapを定義
	Display("mapk", mapk)
}

func Example_structKey() {
	type User struct {
		No   int
		Name string
	}
	// struct型をキーに持つmapを定義
	mapk := map[User]int{
		User{0, "takita"}: 100,
	}
	Display("mapk", mapk)
}
