package sport

type Lines struct {
	Lines map[string]string `json:"lines"`
}

type Sport struct {
	Sport       string
	Coefficient float64
}

type Sports interface {
	Find(name []string) ([]*Sport, error)
	Upsert(u *Sport) error
}
