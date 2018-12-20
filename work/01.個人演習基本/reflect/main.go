package main

import (
	"fmt"
	"reflect"
)

func main() {
	{
		p("/* 数値 */")

		p("// 型の取得")
		p(reflect.TypeOf(0)) //=> int

		p("// 型の種別の比較")
		// 比較対象の定義は https://golang.org/pkg/reflect/#Kind
		p(reflect.TypeOf(0).Kind() == reflect.Int) //=> true

		p("// 型の比較")
		p(reflect.TypeOf(0) == reflect.TypeOf(100)) //=> true

		p("// 変数へ reflect 経由で値をセットする")
		var i int
		// reflect.Value 型に変換
		// ポインタでないと値の変更ができないので &i で変数を渡して .Elem() でポインタの値を返している
		v := reflect.ValueOf(&i).Elem()
		// 値をセットできる
		p(v.CanSet()) //=> true
		// 整数値のセット
		v.SetInt(100)
		p(i) //=> 100
		// Set で汎用的に reflect.Value をセットできる | Type が異なる場合は panic になる
		v.Set(reflect.ValueOf(200))
		p(i) //=> 200

		p("// reflect から Interface に変換")
		if i2, ok := v.Interface().(int); ok {
			p(i2) //=> 200
		}
	}

	{
		p("/* 配列 */")

		p("// 型の取得")
		p(reflect.TypeOf([1]int{0})) //=> [1]int

		p("// 型の種別の比較")
		p(reflect.TypeOf([1]int{0}).Elem()) //=> int

		p("// 型の比較")
		p(reflect.TypeOf([1]int{0}).Kind() == reflect.Array)        //=> true
		p(reflect.TypeOf([1]int{0}) == reflect.TypeOf([1]int{100})) //=> true

		p("// 配列型の作成")
		rt := reflect.ArrayOf(1, reflect.TypeOf(0)) // 引数は要素数と要素の型
		p("Type:", rt, "Kind:", rt.Kind())          //=> Type: [1]int Kind: array

		p("// 要素数の取得")
		p(rt.Len()) //=> 1

		p("// 配列の作成と値のセット")
		// New で与えた型の値が作成されポインタが返却されるので Elem で値を取得
		rv := reflect.New(rt).Elem()
		// Index で任意の位置の要素 (reflect.Value) にアクセスできる
		p(rv.Index(0))
		// SetInt で整数値のセット
		rv.Index(0).SetInt(100)
		p(rv) //=> [100]

		p("// Interface 経由で reflect.Value から [1]int へ変換")
		p(rv.Interface().([1]int)) //=> [100]

		p("// 定義済みの変数へ値をセット")
		ary := [1]int{0}
		reflect.ValueOf(&ary).Elem().Index(0).SetInt(500)
		p(ary) //=> [500]
	}

	{
		p("/* スライス */")

		p("// 型の取得")
		p(reflect.TypeOf([]int{}))      //=> []int
		p(reflect.TypeOf(([]int)(nil))) //=> []int

		p("// 型の種別の比較")
		p(reflect.TypeOf([1]int{0}).Elem()) //=> int

		p("// 型の比較")
		p(reflect.TypeOf([]int{}).Kind() == reflect.Slice)           //=> true
		p(reflect.TypeOf([]int{}) == reflect.TypeOf([]int{1, 2, 3})) //=> true

		p("// スライス型の作成")
		rt := reflect.SliceOf(reflect.TypeOf(0)) // 引数は要素数と要素の型
		p("Type:", rt, "Kind:", rt.Kind())       //=> Type: []int Kind: slice

		p("// スライスの作成")
		// MakeSlice で要素の型、要素数、容量を指定
		rv := reflect.MakeSlice(reflect.TypeOf([]int{}), 1, 1)
		// Index で任意の位置の要素 (reflect.Value) にアクセスできる
		p(rv.Index(0))
		// SetInt で整数値のセット
		rv.Index(0).SetInt(100)
		p(rv) //=> [100]
		// Append での追加もできる
		rv = reflect.Append(rv, reflect.ValueOf(200))
		p(rv) //=> [100 200]
		// AppendSlice でスライス同士の結合もできる
		rv = reflect.AppendSlice(rv, reflect.ValueOf([]int{300}))
		p(rv) //=> [100 200 300]

		p("// Interface 経由で reflect.Value から []int へ変換")
		p(rv.Interface().([]int)) //=> [100 200 300]

		p("// 定義済みの変数へ値をセット (1)")
		s := make([]int, 2, 2)
		rv = reflect.ValueOf(&s).Elem()
		rv.Index(0).SetInt(500)
		rv.Index(1).SetInt(1000)
		p(s) //=> [500]

		p("// 定義済みの変数へ値をセット (2)")
		s = []int{}
		rv = reflect.ValueOf(&s).Elem()
		rv.Set(reflect.Append(rv, reflect.ValueOf(2000))) // Append は新しい reflect.Value を作成するので Set で元の変数を上書きする
		p(s)                                              //=> [2000]
	}

	{
		p("/* マップ */")

		p("// 型の取得")
		p(reflect.TypeOf(map[string]int{}))      //=> map[string]int
		p(reflect.TypeOf((map[string]int)(nil))) //=> map[string]int

		p("// キーの型の取得")
		// p(reflect.TypeOf(map[string]int{}).Key()) //=> string
		type Employee struct {
			Id   int
			Name string
		}
		p(reflect.TypeOf(map[Employee]int{}).Key()) //=> string

		p("// 値の型の取得")
		p(reflect.TypeOf(map[string]int{}).Elem()) //=> int
		p("// 型の比較")
		p(reflect.TypeOf(map[string]int{}).Kind() == reflect.Map)                        //=> true
		p(reflect.TypeOf(map[string]int{}) == reflect.TypeOf(map[string]int{"key1": 1})) //=> true

		p("// マップ型の作成")
		rt := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0)) // 引数はキーの型と値の型
		p("Type:", rt, "Kind:", rt.Kind())                         //=> Type: map[string]int Kind: map

		p("// マップの作成")
		rv := reflect.MakeMap(reflect.TypeOf(map[string]int{}))
		// 値のセット
		rv.SetMapIndex(reflect.ValueOf("key1"), reflect.ValueOf(100))
		p("// キーのリスト取得")

		for _, k := range rv.MapKeys() {
			p(k)
		}
		// 値の取得
		p(rv.MapIndex(reflect.ValueOf("key1"))) //=> 100

		p("// 定義済みの変数へ値をセット")
		m := map[string]int{}
		rv = reflect.ValueOf(&m).Elem()
		rv.SetMapIndex(reflect.ValueOf("key2"), reflect.ValueOf(200))
		rv.SetMapIndex(reflect.ValueOf("key3"), reflect.ValueOf(300))
		p(m) //=> map[key2:200 key3:300]
	}

	{
		p("/* 構造体 */")

		type User struct {
			Name string
			Age  int
		}

		p("// 型の取得")
		p(reflect.TypeOf(User{})) //=> main.User

		p("// 型の比較")
		p(reflect.TypeOf(User{}).Kind() == reflect.Struct)                        //=> true
		p(reflect.TypeOf(User{}) == reflect.TypeOf(User{Name: "user1", Age: 10})) //=> true

		p("// 構造体型の作成")
		t := reflect.StructOf([]reflect.StructField{
			reflect.StructField{Name: "Name", Type: reflect.TypeOf("")},
			reflect.StructField{Name: "Age", Type: reflect.TypeOf(0)},
		})
		p(t) //=> struct { Name string; Age int }

		p("// 構造体の作成")
		rv := reflect.New(reflect.TypeOf(User{})).Elem()

		p("// フィールドの一覧")
		rt := rv.Type()
		p(rv, rt)
		for i := 0; i < rt.NumField(); i++ {
			// フィールドの取得
			f := rt.Field(i)
			// フィールド名
			p(f.Name)
			// 型
			p(f.Type)
			// タグ
			p(f.Tag)
		}

		p("// フィールドの取得")
		if f, ok := rt.FieldByName("Name"); ok {
			p(f.Name, f.Type) //=> Name string
		}

		p("// フィールドの更新")
		rv.Field(0).SetString("user1")
		p(rv.Field(0)) //=> user1

		p("// 定義済みの変数へ値をセット")
		u := User{}
		uv := reflect.ValueOf(&u).Elem()
		uv.Field(0).SetString("user2")
		uv.Field(1).SetInt(20)
		p(u.Name, u.Age) //=> user2 20
	}

	{
		p("/* 関数 */")

		fn := func(s string, i int) string {
			out := ""
			for j := 0; j < i; j++ {
				out += s
			}
			return out
		}

		p("// 型の取得")
		p(reflect.TypeOf(fn)) //=> func(string, int) string

		p("// 型の比較")
		p(reflect.TypeOf(fn).Kind() == reflect.Func)                                        //=> true
		p(reflect.TypeOf(fn) == reflect.TypeOf(func(s string, i int) string { return "" })) //=> true

		p("// 引数の一覧")
		fnt := reflect.TypeOf(fn)
		for i := 0; i < fnt.NumIn(); i++ {
			// 引数の型の取得
			p(fnt.In(i))
		}

		p("// 返り値の一覧")
		for i := 0; i < fnt.NumOut(); i++ {
			// 返り値の型の取得
			p(fnt.Out(i))
		}

		p("// 関数の実行")
		fnv := reflect.ValueOf(fn)
		out := fnv.Call([]reflect.Value{reflect.ValueOf("hello"), reflect.ValueOf(2)})
		if s, ok := out[0].Interface().(string); ok {
			p(s) //=> hellohello
		}
	}
}

func p(a ...interface{}) {
	fmt.Println(a...)
}
