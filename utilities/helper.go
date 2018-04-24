package utilities

import ()

type ResponseJSON struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}

/*
//User Login Check Function
func GetSessionUserDetails(w http.ResponseWriter, r *http.Request) (*models.YpAdminUser, error) {
	fmt.Println("=====================LOGIN================================")
	frontend := r.Header.Get("Authorization")
	if cast.ToString(frontend) == "" {
		for _, cookie := range r.Cookies() { // loop to get all cookies
			fmt.Println("Cookie", cookie.Value)
			if cookie.Name == "frontend" { // only required to frontend cookies
				frontend = cast.ToString(cookie.Value) // check frontend cookie value
			}
		}
	} else {
		auth := strings.Split(frontend, "Bearer ")
		fmt.Println("Authorization token  :", auth)
		frontend = cast.ToString(auth[1])
	}
	userObj, err := models.GetYpAdminUserByAuthId(frontend)
	fmt.Println("=====================END================================")
	return userObj, err
}

//User Check which type of permissions function
func CheckCurrentUserAccessPermission(currentUser *models.YpAdminUser, requestapi string) bool {
	fmt.Println("=====================Permission================================")
	fmt.Println("RequestType", requestapi)
	flag := false
	UserRole, _ := models.GetRoleYpuserByID(currentUser.AuthId)
	for _, value := range UserRole {
		fmt.Println("role:", value["action"], value["enable"])
		if requestapi == cast.ToString(value["action"]) && cast.ToInt(value["enable"]) == 1 {
			flag = true
		}
	}
	fmt.Println(flag)
	fmt.Println("=====================END================================")
	return flag
}

*/
