package internal

// MessageBody : Available actions: get-pipes, block, unblock
// Available pipes (that go under block or unblock): ...
type MessageBody struct {
	User    string   `json:"user,omitempty"`
	Action  string   `json:"action"`
	Block   []string `json:"block,omitempty"`
	Unblock []string `json:"unblock,omitempty"`
}

// ResponseBody : Map response to this struct
type ResponseBody struct {
	Pipename  string   `json:"pipename"`
	Allowed   bool     `json:"allowed"`
	BlockedBy []string `json:"blockedBy"`
}

// FetchPipesMessage : Returns a MessageBody for getting the pipes and their status
func FetchPipesMessage() MessageBody {
	return MessageBody{Action: "get-pipes"}
}

// BlockPipeMessage : Returns a MessageBody for blocking passed pipes
func BlockPipeMessage(user string, pipes []string) MessageBody {
	return MessageBody{
		User:   user,
		Action: "block",
		Block:  pipes,
	}
}

// UnblockPipeMessage : Returns a MessageBody for unblocking passed pipes
func UnblockPipeMessage(user string, pipes []string) MessageBody {
	return MessageBody{
		User:    user,
		Action:  "unblock",
		Unblock: pipes,
	}
}
