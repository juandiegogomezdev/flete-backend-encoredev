package business

import (
	"encore.app/authentication/store"
	"encore.app/pkg/myjwt"
	"encore.app/pkg/resendmailer"
)

type Business struct {
	store     *store.Store
	tokenizer myjwt.JWTTokenizer
	mailer    resendmailer.ResendMailer
}

func NewBusiness(store *store.Store, tokenizer myjwt.JWTTokenizer, mailer resendmailer.ResendMailer) *Business {
	return &Business{store: store, tokenizer: tokenizer, mailer: mailer}
}
