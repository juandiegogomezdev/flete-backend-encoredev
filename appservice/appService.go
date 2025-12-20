package appService

import (
	"encore.app/appservice/appbusiness"
	"encore.app/appservice/appstore"
	"encore.app/pkg/resendmailer"
	"github.com/jmoiron/sqlx"
)

var secrets struct {
	JWT_SECRET_KEY string
	RESEND_API_KEY string
}

var appDBX = sqlx.NewDb(appDB.Stdlib(), "postgres")

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
	m := resendmailer.NewResendMailer(secrets.RESEND_API_KEY, "Acme <onboarding@resend.dev>")
	s := appstore.NewStoreApp(appDB, appDBX)
	b := appbusiness.NewAppBusiness(s, m)

	return &ServiceApp{b: b}, nil
}
