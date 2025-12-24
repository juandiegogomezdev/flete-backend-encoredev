package appService

import (
	"encore.app/appservice/appbusiness"
	"encore.app/appservice/appstore"
	"encore.dev/storage/sqldb"
	"github.com/jmoiron/sqlx"
)

var secrets struct {
	JWT_SECRET_KEY string
	RESEND_API_KEY string
}

var primaryDB = sqldb.Named("primary_db")

var appDBX = sqlx.NewDb(primaryDB.Stdlib(), "postgres")

type ServiceConfig struct {
	BaseUrl string
}

// var cfg *ServiceConfig = config.Load[*ServiceConfig]()

//encore:service
type ServiceApp struct {
	b *appbusiness.BusinessApp
}

func initServiceApp() (*ServiceApp, error) {

	// Initialize the resend mailer
	s := appstore.NewStoreApp(primaryDB, appDBX)
	b := appbusiness.NewAppBusiness(s)

	return &ServiceApp{b: b}, nil
}
