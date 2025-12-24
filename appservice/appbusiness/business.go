package appbusiness

import (
	"encore.app/appservice/appstore"
)

type BusinessApp struct {
	store *appstore.StoreApp
}

func NewAppBusiness(store *appstore.StoreApp) *BusinessApp {
	return &BusinessApp{store: store}
}
