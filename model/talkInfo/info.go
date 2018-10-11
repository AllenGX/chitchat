package talkInfo

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

type ThemeInfoManager struct {
	themeInfos []ThemeInfo
}

type ThemeInfo struct {
	ThemeInfoID  string
	ThemeTitle   string
	UserID       int
	Time         time.Time
	Informations []Information
}

type Information struct {
	InfoID   string
	Time     time.Time
	UserID   int
	ThemeID  string
	InfoBody string
}

func createUUID() (string, error) {
	u1, err := uuid.NewV4()
	if err != nil {
		fmt.Println("uuid error")
		return "", err
	}
	return u1.String(), nil
}

var themeInfoManagers *ThemeInfoManager

func init() {
	var themeInfos []ThemeInfo
	themeInfoManagers = &ThemeInfoManager{themeInfos}
}

//转换时间
func (information *Information) CreatedAtDate() string {
	return information.Time.Format("Jan 2, 2006 at 3:04pm")
}

//转换时间
func (themeInfo *ThemeInfo) CreatedAtDate() string {
	return themeInfo.Time.Format("Jan 2, 2006 at 3:04pm")
}

//得到一个主题的消息数量
func (themeInfo *ThemeInfo) NumReplies() (count int) {
	count = len(themeInfo.Informations)
	return
}

//创建一个帖子
func CreateTheme(themeTitle string, userID int) {
	// no ThemeInfo
	// if themeInfoManagers[userID].themeInfos == nil {
	// 	themeInfoManagers[userID] = ThemeInfoManager{
	// 		themeInfos: []ThemeInfo{
	// 			{
	// 				ThemeTitle: themeTitle,
	// 				UserID:     userID,
	// 				Time:       time.Now(),
	// 			},
	// 		},
	// 	}
	// } else {
	// 	themeInfoManagers[userID].themeInfos = append(
	// 		themeInfoManagers[userID].themeInfos,
	// 		ThemeInfo{
	// 			ThemeTitle: themeTitle,
	// 			UserID:     userID,
	// 			Time:       time.Now(),
	// 		})
	// }
	var informations []Information
	uuid, err := createUUID()
	if err != nil {
		fmt.Println("uudi error:", err.Error())
		panic(err)
	}
	themeInfoManagers.themeInfos = append(themeInfoManagers.themeInfos, ThemeInfo{
		ThemeInfoID:  uuid,
		ThemeTitle:   themeTitle,
		Time:         time.Now(),
		UserID:       userID,
		Informations: informations,
	})
}

//得到某个帖子
func GetTheme(themeInfoID string, userID int) (*ThemeInfo, bool) {
	for k, v := range themeInfoManagers.themeInfos {
		if v.ThemeInfoID == themeInfoID && v.UserID == userID {
			return &themeInfoManagers.themeInfos[k], true
		}
	}
	return &ThemeInfo{}, false
}

// func (themeInfo *ThemeInfo) GetInformationList([]Information) {
// 	return themeInfo.
// }

//得到用户的所有帖子
func GetThemeList(userID int) []ThemeInfo {
	var themeInfoList []ThemeInfo
	for _, v := range themeInfoManagers.themeInfos {
		if v.UserID == userID {
			themeInfoList = append(themeInfoList, v)
		}
	}
	return themeInfoList
}

func DeleteTheme(themeInfoID, userID int) {
	for k, v := range themeInfoManagers.themeInfos {
		if v.UserID == userID {
			themeInfoManagers.themeInfos = append(
				themeInfoManagers.themeInfos[:k],
				themeInfoManagers.themeInfos[k+1:]...,
			)
		}
	}

	// for k, v := range themeInfoManagers[userID].themeInfos {
	// 	if v.ThemeInfoID == themeInfoID {
	// 		themeInfoManagers[userID].themeInfos = append(
	// 			themeInfoManagers[userID].themeInfos[:k],
	// 			themeInfoManagers[userID].themeInfos[k+1:]...,
	// 		)
	// 	}
	// }
}

//为某一个帖子添加信息
func (themeInfo *ThemeInfo) AddInfo(info Information) {
	uuid, err := createUUID()
	if err != nil {
		fmt.Println("uudi error:", err.Error())
		panic(err)
	}

	themeInfo.Informations = append(themeInfo.Informations, Information{
		InfoID:   uuid,
		InfoBody: info.InfoBody,
		ThemeID:  themeInfo.ThemeInfoID,
		Time:     time.Now(),
		UserID:   themeInfo.UserID,
	})

	// //保存修改，否则无法修改成功
	// for k, v := range themeInfoManagers.themeInfos {
	// 	if v.ThemeInfoID == themeInfo.ThemeInfoID && v.UserID == themeInfo.UserID {
	// 		themeInfoManagers.themeInfos[k].Informations = themeInfo.Informations
	// 		break
	// 	}
	// }
}

//为某一个帖子删除信息
func (themeInfo *ThemeInfo) DeleteInfo(infoID string) {
	index := -1
	for k, v := range themeInfo.Informations {
		if v.InfoID == infoID {
			index = k
			break
		}
	}
	if index != -1 {
		themeInfo.Informations = append(themeInfo.Informations[:index], themeInfo.Informations[index+1:]...)
	}
}
