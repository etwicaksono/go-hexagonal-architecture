package valueobject

type SupportedDb string

const (
	SupportedDb_MONGO   SupportedDb = "mongodb"
	SupportedDb_MYSQL   SupportedDb = "mysql"
	SupportedDb_INVALID SupportedDb = "INVALID"
)

func (value SupportedDb) ToString() string {
	switch value {
	case SupportedDb_MONGO:
		return "mongodb"
	case SupportedDb_MYSQL:
		return "mysql"
	default:
		return "INVALID"
	}
}

func SupportedDbFromString(value string) SupportedDb {
	switch value {
	case "mongodb":
		return SupportedDb_MONGO
	case "mysql":
		return SupportedDb_MYSQL
	default:
		return SupportedDb_INVALID
	}
}
