package entity

type Document struct {
	ID       ID
	Name     string
	Username string
}

func NewDocument(name, username string) (*Document, error) {
	u := &Document{
		ID:       NewID(),
		Name:     name,
		Username: username,
	}
	// err := u.Validate()
	// if err != nil {
	// 	return nil, ErrInvalidEntity
	// }
	return u, nil
}
