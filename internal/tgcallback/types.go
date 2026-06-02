package tgcallback

func Init() {
	addCallback(None{})
	addCallback(OpenMenu{})
	addCallback(AddFriend{})
	addCallback(AddFromTelegram{})
	addCallback(FriendList{})
	addCallback(RemoveFriend{})
	addCallback(RemoveFriendApprove{})
	addCallback(RemoveFriendCancel{})
	addCallback(RemoveFromTelegram{})
}

type None struct {
}

type OpenMenu struct {
	Edit bool // If it's true, callback message has to be edited. Otherwise, a new message is sent.
}

type AddFriend struct {
}

type AddFromTelegram struct {
}

type FriendList struct {
	Offset int // How many friends should be skipped
}

type RemoveFriend struct {
}

type RemoveFriendApprove struct {
	UUID string
}

type RemoveFriendCancel struct {
	UUID string
}

type RemoveFromTelegram struct {
}
