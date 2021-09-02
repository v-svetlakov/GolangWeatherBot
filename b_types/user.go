package b_types

type UserArea struct {
	ID          uint64     `json:"id"`
	UserID      uint64     `json:"user"`
	Firstname   string     `json:"firstname"`
	Lastname    string     `json:"lastname"`
	PhoneNumber string     `json:"phonenumber"`
}
