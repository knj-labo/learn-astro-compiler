package main

// テストコード
import "testing"

// TestExercise001 テストコード
// * は、ポインタを示す演算子
// *testing.T は、testingパッケージのT型を示すポインタ
// t *testing.Tは、テストコードを定義する際に必要な引数
func TestExercise001(t *testing.T) {
	// Ex001を呼び出す
	got := Ex001(2000, 2020)
	// 期待値
	want := "2002,2009,2016"
	// 結果を比較する
	if got != want {
		// 結果が期待値と異なるとき、エラーを返す
		t.Errorf("Expected %s but got %s", want, got) // t.Errorf()は、テストが失敗したときに表示するメッセージを返す
	}
}
