package persistence

import (
	"fmt"
	"math/rand"
	"strings"

	pb "github.com/carlosabdoamaral/behealthy_api/protodefs/gen/proto"
)

type AuthQueries struct{}

var (
	AuthQueriesBuilder = &AuthQueries{}
)

func (*AuthQueries) SignIn(req *pb.SignInRequest) string {
	sb := strings.Builder{}
	sb.WriteString(AccountQueriesBuilder.Details())
	sb.WriteString(
		fmt.Sprintf(
			`WHERE email = '%s' AND password = '%s'`, req.GetEmail(), req.GetPassword(),
		),
	)

	return sb.String()
}

func (*AuthQueries) ValidateTwoFactorCode(req *pb.TwoFactorCode) string {
	sb := strings.Builder{}
	sb.WriteString("SELECT 1 FROM account_tb ")
	sb.WriteString(
		fmt.Sprintf(
			`WHERE email = '%s' AND two_factor_code = '%s'`, req.GetEmail(), req.GetTwoFactorCode(),
		),
	)

	return sb.String()
}

func generateRandomCode() string {
	twoFaCode := ""

	for i := 1; i <= 6; i++ {
		randomNumber := rand.Intn(9)
		twoFaCode = fmt.Sprintf("%s%d", twoFaCode, randomNumber)
	}

	return twoFaCode
}

func (*AuthQueries) RefreshTwoFactorCode(req *pb.RefreshTwoFactorCodeRequest) string {
	return fmt.Sprintf(`
	UPDATE account_tb
	SET two_factor_code='%s'
	WHERE email='%s'
	`, generateRandomCode(), req.GetEmail(),
	)
}
