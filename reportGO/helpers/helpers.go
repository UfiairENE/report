package helpers

import (
	"strings"
	"time"

	errorLog "github.com/nduson/txn-report/errors"

	"github.com/xeonx/timeago"
)

func CheckUserPass(username, password string) bool {
	userpass := make(map[string]string)
	userpass["hello"] = "itsme"
	userpass["john"] = "doe"

	if val, ok := userpass[username]; ok {
		if val == password {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}

type Config struct {
	DefaultLayout string
}

func ConvertDatetime(datetime string) string {
	s, err := time.Parse("Jan 02, 2006 3:04:05 PM", datetime)
	if err != nil {
		errorLog.LogError(err)
	}
	return timeago.English.Format(s)

}
