package mrello

type Board struct {
	Todo  []*Card `json:"todo"`
	Doing []*Card `json:"doing"`
	Done  []*Card `json:"done"`
}

func CreateBoardFromCards(cards []*Card) *Board {
	var board Board

	for _, card := range cards {
		switch card.Column {
		case "todo":
			board.Todo = append(board.Todo, card)
		case "doing":
			board.Doing = append(board.Doing, card)
		case "done":
			board.Done = append(board.Done, card)
		}
	}

	return &board
}
