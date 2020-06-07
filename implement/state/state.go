package state

import "fmt"

type VoteState interface {
	Vote(user, voteItem string, manager *VoteManager)
}

type NormalVoteState struct {}

func (n *NormalVoteState) Vote(user, voteItem string, manager *VoteManager) {
	// 正常投票
	// 记录到投票记录中
	manager.MapVote[user] = voteItem
	fmt.Println("恭喜你投票成功！")
}

type RepeatVoteState struct {}

func (n *RepeatVoteState) Vote(user, voteItem string, manager *VoteManager) {
	// 重复投票
	// 暂不做处理
	fmt.Println("请不要重复投票！")
}

type SpiteVoteState struct {}

func (n *SpiteVoteState) Vote(user, voteItem string, manager *VoteManager) {
	// 恶意投票
	// 取消用户投票资格，并取消投票记录
	delete(manager.MapVote, user)
	fmt.Println("你有恶意刷票行为，取消投票资格！")
}

type BlackVoteState struct {}

func (n *BlackVoteState) Vote(user, voteItem string, manager *VoteManager) {
	// 黑名单
	// 禁止登陆系统了
	delete(manager.MapVote, user)
	fmt.Println("进入黑名单，将禁止登陆和使用本系统！")
}

type VoteManager struct {
	// 持有状态处理对象
	state VoteState
	// 记录用户投票结果，用户名称 -> 投票选项
	MapVote map[string]string
	// 记录用户投票次数，用户名称 -> 投票次数
	mapVoteCount map[string]int
}

func NewVoteManager() VoteManager {
	return VoteManager{
		state: nil,
		MapVote: make(map[string]string),
		mapVoteCount: make(map[string]int),
	}
}

func (m *VoteManager) Vote(user, voteItem string) {
	m.mapVoteCount[user]++
	if count := m.mapVoteCount[user]; count == 1 {
		m.state = new(NormalVoteState)
	} else if count > 1 && count < 5 {
		m.state = new(RepeatVoteState)
	} else if count >= 5 && count < 8 {
		m.state = new(SpiteVoteState)
	} else {
		m.state = new(BlackVoteState)
	}

	// 转调状态对象来进行相应的操作
	m.state.Vote(user, voteItem, m)
}


