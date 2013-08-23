package gorden

type SessionConfig struct {
    CookieName string
    CookieKey []byte
    Path string
    Domain string
    MaxAge int
    Secure bool
    HttpOnly bool
}