package calc

import "testing"

// TestSuｍ　加算テスト
// 引数：testing.Tを渡す
//Testから始まる名称にするとgo testでの実行対象になる

func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		t.Fatal("sum(1,2) should be 3, but doesn't match")
	}
}
