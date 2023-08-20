// Response Body {"id":"6807E49C-D68F-4D38-A990-6FACCB0E51A3","title":"Buy milk","created_at":"2020-12-20T15:04:05Z","completed_at": "2020-12-20T15:04:05Z"}type
package todo

import "time"

type Todo struct { // struct tag
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
}
