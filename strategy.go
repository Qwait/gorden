package gorden

type Strategy interface {
    Authenticate(arguments interface{}) bool
    IsAuthenticated() bool
}