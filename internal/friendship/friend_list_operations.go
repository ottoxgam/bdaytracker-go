package friendship

import (
	"sort"
)

func SortFriends(friends []Friend) []Friend {
	sort.Slice(friends, func(i, j int) bool {
		if friends[i].BMonth != friends[j].BMonth {
			return friends[i].BMonth < friends[j].BMonth
		}
		if friends[i].BDay != friends[j].BDay {
			return friends[i].BDay < friends[j].BDay
		}
		return friends[i].UUID < friends[j].UUID
	})
	return friends
}

func RemoveTelegramFriends(friends []Friend) []Friend {
	var result []Friend
	for _, friend := range friends {
		if friend.TelegramUsername == nil {
			result = append(result, friend)
		}
	}
	return result
}

func HasTelegramFriends(friends []Friend) bool {
	for _, friend := range friends {
		if friend.TelegramUsername != nil {
			return true
		}
	}
	return false
}
