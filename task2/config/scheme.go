package config

// Scheme represents the application configuration scheme.
type Scheme struct {
	// Env is the application environment.
	Env string

	Db   *Db
	Grpc *Grpc
}

// Db is service Data base connection params
type Db struct {
	Address string
}

// Grpc represent Grpc server listener params
type Grpc struct {
	Address string
	Timeout string
}

// Http represent Http server listener params
//type Http struct {
//	Address string
//	Timeout string
//}
