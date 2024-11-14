package valueobject

type SuportedDb string

const (
	SuportedDb_MONGO   SuportedDb = "mongodb"
	SuportedDb_MYSQL   SuportedDb = "mysql"
	SuportedDb_INVALID SuportedDb = "INVALID"
)

func (value SuportedDb) ToString() string {
	switch value {
	case SuportedDb_MONGO:
		return "mongodb"
	case SuportedDb_MYSQL:
		return "mysql"
	default:
		return "INVALID"
	}
}

func SuportedDbFromString(value string) SuportedDb {
	switch value {
	case "mongodb":
		return SuportedDb_MONGO
	case "mysql":
		return SuportedDb_MYSQL
	default:
		return SuportedDb_INVALID
	}
}
