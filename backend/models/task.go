package models

type User struct {
	ID          uint   `json:"id"`
	FirebaseUID string `json:"firebase_uid"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	CreatedAt   string `json:"created_at"`
}

type Task struct {
	TaskID     uint   `json:"task_id"`
	UserID     uint   `json:"user_id"`
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
	CreatedAt  string `json:"created_at"`
}

// type EmployeeSalary struct {
// 	ID         uint          `json:"id" sql:"primary_key"`
// 	FirstName  string        `json:"first_name"`
// 	LastName   string        `json:"last_name"`
// 	Occupation string        `json:"occupation"`
// 	Salary     int64         `json:"salary"`
// 	DeptID     sql.NullInt16 `json:"dept_id"`
// }
