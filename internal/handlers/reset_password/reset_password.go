package resetpassword

import (
	useraccount "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/account_model"
	resetpasswordcode "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/reset_password_model"
	"github.com/Alexander-s-Digital-Marketplace/auth-service/internal/utils/email"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// User's reset password
func ResetPasswordHandle(c *gin.Context) (int, string) {

	var userFront useraccount.UserAccount
	var userDB useraccount.UserAccount
	var err error

	userFront.DecodeFromContext(c)
	userDB.Email = userFront.Email

	err = userDB.GetFromTableByEmail()
	if err != nil {
		logrus.Errorln("Incorrect login")
		return 403, "Incorrect login or email"
	}

	if userDB.Email != userFront.Email {
		logrus.Errorln("Incorrect email")
		return 403, "Incorrect login or email"
	}

	var resetForm resetpasswordcode.ResetCode

	resetForm.GenerateCode()
	resetForm.InitDate(5)

	err = email.SendEmail(userDB.Email, userDB.Email, resetForm.Code)
	if err != nil {
		return 503, "Error send email"
	}

	resetForm.User = userDB
	resetForm.AddToTable()

	logrus.Infoln("Send reset email is successful. User: ", userDB.Id, userDB.Email,
		" EMAIL: ", userDB.Email, "code: ", resetForm.Code, "time: ", resetForm.StartTime, " ", resetForm.ExpTime)

	return 200, "Send reset email is successful"
}
