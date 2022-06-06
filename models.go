package gotconf

type AuthSuccess struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type User struct {
	Id          string      `json:"id"`           //User unique identifier.
	Uid         string      `json:"uid"`          //User unique identifier for calls.
	Avatar      *string     `json:"avatar"`       //User avatar address. If there is no avatar, null is returned.
	LoginName   string      `json:"login_name"`   //User username. Size range: 1..32
	Password    string      `json:"password"`     //User password (POST | PUT).	Size range: 1..255
	Email       string      `json:"email"`        // User unique email address. Size range: 3..250
	DisplayName string      `json:"display_name"` //User display name. Size range: 1..64
	FirstName   string      `json:"first_name"`   //User first name. Size range: 0..64
	LastName    string      `json:"last_name"`    // User last name. Size range: 0..64
	Company     string      `json:"company"`      //User company.	Size range: 0..64
	Groups      []GroupMini `json:"groups"`       // List of ObjectGroupMini.
	MobilePhone string      `json:"mobile_phone"` // User mobile phone. Size range: 0..50
	WorkPhone   string      `json:"work_phone"`   // User work phone. Size range: 0..50
	HomePhone   string      `json:"home_phone"`   // User home phone. Size range: 0..50
	Status      int         `json:"status"`       // NOT_ACTIVE: -2, INVALID: -1, OFFLINE: 0, ONLINE: 1, BUSY: 2, MULTIHOST: 5 (The user is in the conference and is the conference owner), UNDEFINED: -127.
	IsActive    int         `json:"is_active"`    // Account status: 1 - enabled, 0 - disabled.
}

type GroupMini struct {
	Id          string `json:"id"`           // Unique group ID.
	DisplayName string `json:"display_name"` // Unique group name.
}
