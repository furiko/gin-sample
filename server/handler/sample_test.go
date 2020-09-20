package handler

import "testing"

func TestSample(t *testing.T) {
	Sample("gin-sample")
	t.Errorf("inteded error: %s", "Actionsでテスト通らなかったら次実行されないかを確認")
}
