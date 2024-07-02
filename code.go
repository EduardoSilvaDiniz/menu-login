
//func addUser(user, password string) {
//	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
//	if err != nil {
//		panic(hash)
//	}
//
//	users[user] = hash
//}
//
//func passwordCheck(user, password string) bool {
//	err := bcrypt.CompareHashAndPassword(users[user], []byte(password))
//	if err != nil {
//		return false
//	}
//	return true
//}
