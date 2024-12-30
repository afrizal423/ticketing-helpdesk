package wa

import (
	"strings"

	"go.mau.fi/whatsmeow/types"
)

func getNomorHP(arg string) (phoneNumber string) {
	if strings.Contains(string(arg), ":") {
		parts := strings.Split(arg, "@")
		phoneParts := strings.Split(parts[0], ":")
		phoneNumber = phoneParts[0]
	} else {
		parts := strings.Split(arg, "@")
		phoneNumber = parts[0]
	}
	return
}

func parseJID(arg string) (types.JID, bool) {
	if arg[0] == '+' {
		arg = arg[1:]
	}
	if !strings.ContainsRune(arg, '@') {
		return types.NewJID(arg, types.DefaultUserServer), true
	} else {
		recipient, err := types.ParseJID(arg)
		if err != nil {
			// log.Error().Err(err).Msg("Invalid JID")
			panic(err)
			return recipient, false
		} else if recipient.User == "" {
			// log.Error().Err(err).Msg("Invalid JID no server specified")
			panic(err)
			return recipient, false
		}
		return recipient, true
	}
}
