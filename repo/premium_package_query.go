package repo

var (
	// premium package
	RepoGetListPremiumPackage = `
	SELECT "id", "package_uid", "title", "description", "price", "is_active", "created_at", "updated_at",
	"created_by", "updated_by" FROM premium_package WHERE is_active IS TRUE 
	%s %s %s;`
	RepoGetPremiumPackageByPackageUID = `
	SELECT "id", "package_uid", "title", "description", "price", "is_active", "created_at", "updated_at",
	"created_by", "updated_by" FROM premium_package WHERE package_uid = $1;`

	// premium package user
	RepoGetPremiumPackageUserByAccountID = `
	SELECT "id", "premium_package_id", "account_id", "purchased_date" FROM premium_package_user WHERE account_id = $1;`

	InsertPremiumPackageUser = `
	INSERT INTO premium_package_user ("premium_package_id", "account_id")
	VALUES ($1, $2)
	RETURNING "id", "purchased_date";`
)
