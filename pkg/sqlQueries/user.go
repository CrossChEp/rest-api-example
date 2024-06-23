package sqlQueries

const (
	UserTable           = "user_schema.users"
	UserIDColumnName    = "id"
	NameColumnName      = "name"
	EmailColumnName     = "email"
	PasswordColumnName  = "password"
	CreatedAtColumnName = "create_at"
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
