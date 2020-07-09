package

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestFoo(t *testing.T) {
	// mockのコントローラー生成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mockの生成
	m := mock_foo.NewMockFoo(ctrl)

	m.
		EXPECT().
		Bar(99).
		Return(101)

	// mockの呼び出し
	m.Bar(99)
}
