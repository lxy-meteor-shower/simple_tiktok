package constants

const (
	PlayURL = "192.168.43.163"

	UserTableName       = "user"
	VideoTableName      = "video"
	FavoriteTableName   = "favorite"
	CommentTableName    = "comment"
	MySQLDefaultDSN     = "tiktok:tiktok@tcp(127.0.0.1:9910)/tiktok?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress         = "127.0.0.1:2379"
	MinioAddress        = "127.0.0.1:9000"
	MinioUserName       = "tiktok"
	MinioPassword       = "tiktok123"
	ApiServiceName      = "demoapi"
	VideoServiceName    = "demovideo"
	InteractServiceName = "demointeract"
	UserServiceName     = "demouser"
	SecretKey           = "secret key"
	IdentityKey         = "id"
)
