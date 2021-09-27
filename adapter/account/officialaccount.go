package officialaccount

import (
	"github.com/silenceper/wechat/v2/officialaccount"
)

var account *officialaccount.OfficialAccount

func GetWechatAccount() *officialaccount.OfficialAccount {
	if account != nil {
		return account
	}
	return nil
}

func SetWechatAccount(a *officialaccount.OfficialAccount) {
	if a != nil {
		account = a
	}
}