package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseMomentGetMomentTask struct {
	*response.ResponseWork

	TaskList   []*power.HashMap `json:"task_list"`
	NextCursor string           `json:"next_cursor"`
}
