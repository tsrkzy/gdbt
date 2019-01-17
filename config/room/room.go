package room

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/lepra-tsr/gdbt/api"
	. "github.com/lepra-tsr/gdbt/api/organization"
	. "github.com/lepra-tsr/gdbt/api/room"
	. "github.com/lepra-tsr/gdbt/api/user"
	. "github.com/lepra-tsr/gdbt/config"
)

type ConfigJson struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	IconUrl     string     `json:"iconUrl"`
	CurrentRoom *RoomInfo  `json:"currentRoom"`
	Rooms       []RoomInfo `json:"rooms"`
}

func (u *ConfigJson) ParseServerEntity(
	userJson *UserJson,
	orgJson *OrganizationJson,
	roomJson *RoomJson) error {
	users := userJson.Users
	for i := 0; i < len(users); i++ {
		user := users[i]
		if user.Email == "" {
			continue
		}
		u.Id = user.Id
		u.Name = user.Name
	}
	joins := roomJson.Joins
	joinIdList := []int{}
	for i := 0; i < len(joins); i++ {
		join := joins[i]
		joinId := join.Id
		guyId := join.GuyId
		if guyId == u.Id {
			joinIdList = append(joinIdList, joinId)
		}
	}

	orgs := orgJson.Organizations
	rooms := roomJson.Rooms
	u.Rooms = []RoomInfo{}
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		roomJoinIds := room.JoinIds
		// 自分の持つjoinIdが、当該roomのjoinsに含まれるかチェック
		for j := 0; j < len(joinIdList); j++ {
			joinId := joinIdList[j]
			joinedToRoom := false
			for k := 0; k < len(roomJoinIds); k++ {
				roomJoinId := roomJoinIds[k]
				if joinId != roomJoinId {
					continue
				} else {
					// roomが見つかったら、それに紐づくorganizationを探す
					roomOrgId := room.OrganizationId
					for l := 0; l < len(orgs); l++ {
						org := orgs[l]
						if org.Id != roomOrgId {
							continue
						}
						roomInfo := RoomInfo{}
						roomInfo.ParseRoomOrganization(&room, &org)
						u.Rooms = append(u.Rooms, roomInfo)
						joinedToRoom = true
						break
					}
					break
				}
			}
			if joinedToRoom {
				break
			}
		}
	}

	return nil
}

type RoomInfo struct {
	Id           int               `json:"id"`
	Name         string            `json:"name"`
	MessageUrl   string            `json:"messageUrl"`
	Organization *OrganizationInfo `json:"organization"`
}

func (u *RoomInfo) ParseRoomOrganization(room *Room, org *Organization) error {
	u.Id = room.Id
	u.Name = room.Name
	u.MessageUrl = room.Links.Messages
	orgInfo := OrganizationInfo{}
	orgInfo.Id = org.Id
	orgInfo.Name = org.Name
	orgInfo.Slug = org.Slug
	u.Organization = &orgInfo

	return nil
}

type OrganizationInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (u *ConfigJson) Write() error {
	file, err := os.OpenFile(
		RoomJsonPath,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0774)
	if err != nil {
		return err
	}
	defer file.Close()
	jsonBytes, _ := json.Marshal(&u)
	file.Write(jsonBytes)

	return nil
}

func (u *ConfigJson) Read() error {
	if bytes, err := ioutil.ReadFile(CredentialJsonPath); err != nil {
		return err
	} else {
		if err := json.Unmarshal(bytes, &u); err != nil {
			fmt.Println("Failed to parse credential.json")
			return err
		}
		return nil
	}
}
