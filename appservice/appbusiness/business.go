package appbusiness

import (
	"encore.app/appservice/appstore"
	"encore.app/pkg/resendmailer"
)

type BusinessApp struct {
	store  *appstore.StoreApp
	mailer resendmailer.ResendMailer
}

func NewAppBusiness(store *appstore.StoreApp, mailer resendmailer.ResendMailer) *BusinessApp {
	return &BusinessApp{store: store, mailer: mailer}
}
