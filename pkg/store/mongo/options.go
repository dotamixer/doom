package mongo

type Options struct {
	Hosts       []string
	Username    string
	Password    string
	AuthSource  string
	MaxPoolSize uint64
	ReplicaSet  string
}
