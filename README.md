gorden
======

WIP authentication for web applications

Sample Setup
------------

    import "github.com/qwait/gorden"

    gorden.AddStrategy("password", PasswordStrategy{})}

    type PasswordStrategy struct {}
    
    func (PasswordStrategy) Authenticate(arguments interface{}) bool {
        return false
    }

    func (PasswordStrategy) IsAuthenticated() bool {
        return false
    }

    gorden.NewManager("password", gorden.SessionConfig{CookieName: "auth_token", CookieKey: []byte("so secret")})