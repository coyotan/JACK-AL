package botutils

//This function will DEFINITELY need to be updated in the future. I think we might need to use Presences, which is a privileged intent. This means eventually, we will need to pass session into here.
func CheckAdminPermissions(authorID string, guildID string) (isAdmin bool, err error){
	if authorID == "228355771526676480" && len(guildID) > 0 {
		isAdmin = true
	} else {

	}
	return
}

