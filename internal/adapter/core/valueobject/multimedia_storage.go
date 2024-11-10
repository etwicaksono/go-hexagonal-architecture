package valueobject

type MultimediaStorage int32

const (
	MultimediaStorage_LOCAL   MultimediaStorage = 1
	MultimediaStorage_MINIO   MultimediaStorage = 2
	MultimediaStorage_INVALID MultimediaStorage = 0
)

func (value MultimediaStorage) ToInt32() int32 {
	return int32(value)
}

func (value MultimediaStorage) ToString() string {
	switch value {
	case MultimediaStorage_LOCAL:
		return "LOCAL"
	case MultimediaStorage_MINIO:
		return "MINIO"
	default:
		return "INVALID"
	}
}

func MultimediaStorageFromInt32(value int32) MultimediaStorage {
	switch value {
	case 1:
		return MultimediaStorage_LOCAL
	case 2:
		return MultimediaStorage_MINIO
	default:
		return MultimediaStorage_INVALID
	}

}
func MultimediaStorageFromString(value string) MultimediaStorage {
	switch value {
	case "LOCAL":
		return MultimediaStorage_LOCAL
	case "MINIO":
		return MultimediaStorage_MINIO
	default:
		return MultimediaStorage_INVALID
	}
}
