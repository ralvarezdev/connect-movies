package postgres

const (
	// CreateUserReviewProc is the query to call the stored procedure to create user review
	CreateUserReviewProc = "CALL create_user_review($1, $2, $3, $4)"
	
	// UpdateUserReviewProc is the query to call the stored procedure to update user review
	UpdateUserReviewProc = "CALL update_user_review($1, $2, $3, $4)"
	
	// DeleteUserReviewProc is the query to call the stored procedure to delete user review
	DeleteUserReviewProc = "CALL delete_user_review($1, $2)"
)