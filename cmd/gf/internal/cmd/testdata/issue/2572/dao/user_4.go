// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"/internal"
)

// internalUser4Dao is an internal type for wrapping the internal DAO implementation.
type internalUser4Dao = *internal.User4Dao

// user4Dao is the data access object for the table user4.
// You can define custom methods on it to extend its functionality as needed.
type user4Dao struct {
	internalUser4Dao
}

var (
	// User4 is a globally accessible object for table user4 operations.
	User4 = user4Dao{
		internal.NewUser4Dao(),
	}
)

// Add your custom methods and functionality below.
