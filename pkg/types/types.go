package types

type Money int64

type Category string

type Payment struct {
	ID       int
	Category Category
	Amount   Money
	Status   Status
}

type Status string

const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)
