package authentication

import (
	"encore.app/authentication/business"
	"encore.app/authentication/store"
	"encore.app/pkg/myjwt"
	"encore.app/pkg/resendmailer"
	"encore.dev/storage/sqldb"
	"github.com/jmoiron/sqlx"
)

var secrets struct {
	JWT_SECRET_KEY string
	RESEND_API_KEY string
}

var appDB = sqldb.NewDatabase("db_app", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})
var appDBX = sqlx.NewDb(appDB.Stdlib(), "postgres")

type ServiceConfig struct {
	BaseUrl string
}

//encore:service
type Authentication struct {
	b *business.Business
}

func initAuthentication() (*Authentication, error) {
	tokenizer := myjwt.NewJWTTokenizer(secrets.JWT_SECRET_KEY)

	// Initialize the resend mailer
	m := resendmailer.NewResendMailer(secrets.RESEND_API_KEY, "Acme <onboarding@resend.dev>")
	s := store.NewStore(appDB, appDBX)
	b := business.NewBusiness(s, tokenizer, m)

	return &Authentication{b: b}, nil
}
