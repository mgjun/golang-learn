package jwtauth

type options struct {
}

type JWTAuth struct {
	opts  *options
	store *Storer
}
