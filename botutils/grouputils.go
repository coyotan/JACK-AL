package botutils

import "github.com/bwmarrin/discordgo"

var (
	//Byte codes for discord permissions.
	administrator int64 = 0x8
)

//CheckAdminPermissions will DEFINITELY need to be updated in the future. I think we might need to use Presences, which is a privileged intent. This means eventually, we will need to pass session into here.
func CheckAdminPermissions(s *discordgo.Session, m *discordgo.Message) (isAdmin bool, err error) {

	//start false, wait to see if true.
	isAdmin = false

	switch m.Author.ID {
	case "228355771526676480":
		isAdmin = true
		break

	default:
		member, err := s.State.Member(m.GuildID, m.Author.ID)

		if err != nil {
			break
		}

		for _, iD := range member.Roles {
			role, err := s.State.Role(m.GuildID, iD)

			if err != nil {
				break
			}

			//Search if *any role* the user has, has Admin permissions.
			if administrator == (role.Permissions & administrator) {
				isAdmin = true
			}
		}
		break
	}

	return
}
