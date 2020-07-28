package condition

import "testing"

func TestSwitch1(t *testing.T) {
	for i := 0; i <= 5; i++ {
		switch i {
		case 0, 2, 4:
			t.Logf("偶数 %d", i)
		case 1, 3:
			t.Logf("奇数 %d", i)
		default:
			t.Log("不在0-4中间的数" + string(i))
		}
	}
}

//这里的switch没有条件  和if-else if-else功能相同
func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i <= 5; i++ {
		switch {
		case i%2 == 0:
			t.Logf("偶数 %d", i)
		case i%2 == 1:
			t.Logf("奇数 %d", i)
		default:
			t.Log("unknow")
		}
	}
}
