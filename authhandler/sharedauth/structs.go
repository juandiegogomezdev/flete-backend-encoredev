package sharedauth

type ErrorDetailsToken struct {
	TokenStatus string `json:"token_status"`
}

func (t ErrorDetailsToken) ErrDetails() {}
