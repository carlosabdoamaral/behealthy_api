package persistence

import "fmt"

type AccountQueries struct{}

var (
	AccountQueriesBuilder = &AccountQueries{}
)

func (*AccountQueries) Details() string {
	return `
	SELECT
		id,
		fullname,
		username,
		email,
		password,
		birthdate,
		role,
		created_at,
		updated_at,
		soft_deleted,
		is_blocked,
		two_factor_code
	FROM
		account_tb
	`
}

func (*AccountQueries) UpdatePassword(email, password string) string {
	return fmt.Sprintf(`
	UPDATE account_tb
	SET password = '%s'
	WHERE email = '%s'
	`, password, email)
}

func (*AccountQueries) SoftDelete(email string) string {
	return fmt.Sprintf(`
	UPDATE account_tb
	SET soft_deleted = TRUE
	WHERE email = '%s'
	`, email)
}
