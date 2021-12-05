package sobot

func Instagram(user, pass string) *AccountInfo {
	//&AccountInfo{Username: user.Username, Password: user.Password}
	return &AccountInfo{Username: user, Password: pass, Pname: "instagram", FilePath: "", Caption: ""}
}

func Twitter(user, pass string) *AccountInfo {
	//&AccountInfo{Username: user.Username, Password: user.Password}
	return &AccountInfo{Username: user, Password: pass, Pname: "twitter", FilePath: "", Caption: ""}

}
