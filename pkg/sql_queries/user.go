package sql_queries

const (
	UserTable           = "users"
	UserIDColumnName    = "id"
	NameColumnName      = "name"
	EmailColumnName     = "email"
	PasswordColumnName  = "password"
	CreatedAtColumnName = "created_at"
	DeletedAtColumnName = "deleted_at"
)

var (
	InsertUserColumns = []string{
		NameColumnName,
		EmailColumnName,
		PasswordColumnName,
		CreatedAtColumnName,
	}
	GetUserColumns = []string{
		UserIDColumnName,
		NameColumnName,
		EmailColumnName,
		PasswordColumnName,
		CreatedAtColumnName,
	}
)
