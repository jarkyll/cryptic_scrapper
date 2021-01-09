package service

func Login() {

}

func Register() {
	// given the email and password
	// check if the email is in the db
	// then check if the passwrod is good
	// check if
}

func Logout() {

}

func RefreshToken() {

}

//func CreateAuth(userid int64, td *models.TokenDetails) error {
//	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
//	rt := time.Unix(td.RtExpires, 0)
//	now := time.Now()
//
//	errAccess := db.GetRedis().Set(td.AccessUUID, strconv.Itoa(int(userid)), at.Sub(now)).Err()
//	if errAccess != nil {
//		return errAccess
//	}
//	errRefresh := db.GetRedis().Set(td.RefreshUUID, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
//	if errRefresh != nil {
//		return errRefresh
//	}
//	return nil
//}
//
//// check into this
//func FetchAuth(authDetails *models.AccessDetails) (int64, error) {
//	userid, err := db.GetRedis().Get(authDetails.AccessUUID).Result()
//	if err != nil {
//		return 0, err
//	}
//	userID, _ := strconv.ParseInt(userid, 10, 64)
//	return userID, nil
//}
//
//func DeleteAuth(givenUUID string) (int64, error) {
//	deleted, err := db.GetRedis().Del(givenUUID).Result()
//	if err != nil {
//		return 0, err
//	}
//	return deleted, nil
//}