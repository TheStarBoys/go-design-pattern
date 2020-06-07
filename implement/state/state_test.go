package state

import "testing"

func TestState(t *testing.T) {
	manager := NewVoteManager()
	for i := 0; i < 10; i++ {
		manager.Vote("user1", "A")
	}
	// Output:
	// 恭喜你投票成功！
	// 请不要重复投票！
	// 请不要重复投票！
	// 请不要重复投票！
	// 你有恶意刷票行为，取消投票资格！
	// 你有恶意刷票行为，取消投票资格！
	// 你有恶意刷票行为，取消投票资格！
	// 进入黑名单，将禁止登陆和使用本系统！
	// 进入黑名单，将禁止登陆和使用本系统！
	// 进入黑名单，将禁止登陆和使用本系统！
}
