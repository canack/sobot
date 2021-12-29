package sobot

import "log"

func (user *AccountInfo) SetFile(data string) *AccountInfo {
	return &AccountInfo{
		Username: user.Username,
		Password: user.Password,
		FilePath: data,
		Pname:    user.Pname,
		Caption:  user.Caption,
	}
}

func (user *AccountInfo) SetCaption(caption string) *AccountInfo {
	return &AccountInfo{
		Username: user.Username,
		Password: user.Password,
		FilePath: user.FilePath,
		Pname:    user.Pname,
		Caption:  caption,
	}
}

func (user *AccountInfo) Share(debug bool) {
	if user.Pname == "instagram" {
		// call shareInstagram function
		log.Println(user.Username, "******", user.Pname)
		user.shareInstagram(debug)

	} else if user.Pname == "twitter" {
		// call shareTwitter function
		log.Println(user.Username, "******", user.Pname)
		user.shareTwitter(debug)

	}

}
