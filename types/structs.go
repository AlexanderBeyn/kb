package types

// https://docs.kanboard.org/en/latest/api/index.html
// https://mholt.github.io/json-to-go/

type Project struct {
	ID                  int    `json:"id,string"`
	Name                string `json:"name"`
	IsActive            int    `json:"is_active,string"`
	Token               string `json:"token"`
	LastModified        int    `json:"last_modified,string"`
	IsPublic            int    `json:"is_public,string"`
	IsPrivate           int    `json:"is_private,string"`
	DefaultSwimlane     string `json:"default_swimlane"`
	ShowDefaultSwimlane int    `json:"show_default_swimlane,string"`
	Description         string `json:"description"`
	Identifier          string `json:"identifier"`
	URL                 struct {
		Board    string `json:"board"`
		Calendar string `json:"calendar"`
		List     string `json:"list"`
	} `json:"url"`
}

type Column struct {
	ID        int    `json:"id,string"`
	Title     string `json:"title"`
	Position  int    `json:"position,string"`
	ProjectID int    `json:"project_id,string"`
	TaskLimit int    `json:"task_limit,string"`
}

type Me struct {
	ID                   int    `json:"id"`
	Username             string `json:"username"`
	Role                 string `json:"role"`
	IsLdapUser           bool   `json:"is_ldap_user"`
	Name                 string `json:"name"`
	Email                string `json:"email"`
	GoogleID             string `json:"google_id"`
	GithubID             string `json:"github_id"`
	NotificationsEnabled int    `json:"notifications_enabled,string"`
	Timezone             string `json:"timezone"`
	Language             string `json:"language"`
	DisableLoginForm     int    `json:"disable_login_form,string"`
	TwofactorActivated   bool   `json:"twofactor_activated"`
	TwofactorSecret      string `json:"twofactor_secret"`
	Token                string `json:"token"`
	NotificationsFilter  int    `json:"notifications_filter,string"`
}

type Task struct {
	ID                  int         `json:"id,string"`
	Title               string      `json:"title"`
	Description         string      `json:"description"`
	DateCreation        int         `json:"date_creation,string"`
	ColorID             string      `json:"color_id"`
	ProjectID           int         `json:"project_id,string"`
	ColumnID            int         `json:"column_id,string"`
	OwnerID             int         `json:"owner_id,string"`
	Position            int         `json:"position,string"`
	IsActive            int         `json:"is_active,string"`
	DateCompleted       int         `json:"date_completed,string"`
	Score               int         `json:"score,string"`
	DateDue             int         `json:"date_due,string"`
	CategoryID          int         `json:"category_id,string"`
	CreatorID           int         `json:"creator_id,string"`
	DateModification    int         `json:"date_modification,string"`
	Reference           string      `json:"reference"`
	DateStarted         int         `json:"date_started,string"`
	TimeSpent           int         `json:"time_spent,string"`
	TimeEstimated       int         `json:"time_estimated,string"`
	SwimlaneID          int         `json:"swimlane_id,string"`
	DateMoved           int         `json:"date_moved,string"`
	RecurrenceStatus    string      `json:"recurrence_status"`
	RecurrenceTrigger   string      `json:"recurrence_trigger"`
	RecurrenceFactor    string      `json:"recurrence_factor"`
	RecurrenceTimeframe string      `json:"recurrence_timeframe"`
	RecurrenceBasedate  string      `json:"recurrence_basedate"`
	RecurrenceParent    interface{} `json:"recurrence_parent"`
	RecurrenceChild     interface{} `json:"recurrence_child"`
	URL                 string      `json:"url"`
	Color               struct {
		Name       string `json:"name"`
		Background string `json:"background"`
		Border     string `json:"border"`
	} `json:"color"`
}

type SearchTasksParams struct {
	ProjectID int    `json:"project_id"`
	Query     string `json:"query"`
}

type CreateTaskParams struct {
	OwnerID             int    `json:"owner_id,omitempty"`
	CreatorID           int    `json:"creator_id,omitempty"`
	DateDue             string `json:"date_due,omitempty"`
	Description         string `json:"description,omitempty"`
	CategoryID          int    `json:"category_id,omitempty"`
	Score               int    `json:"score,omitempty"`
	Title               string `json:"title"`
	ProjectID           int    `json:"project_id"`
	ColorID             string `json:"color_id,omitempty"`
	ColumnID            int    `json:"column_id,omitempty"`
	RecurrenceStatus    int    `json:"recurrence_status,omitempty"`
	RecurrenceTrigger   int    `json:"recurrence_trigger,omitempty"`
	RecurrenceFactor    int    `json:"recurrence_factor,omitempty"`
	RecurrenceTimeframe int    `json:"recurrence_timeframe,omitempty"`
	RecurrenceBasedate  int    `json:"recurrence_basedate,omitempty"`
}

type MoveTaskPositionParams struct {
	ProjectID  int `json:"project_id"`
	TaskID     int `json:"task_id"`
	ColumnID   int `json:"column_id"`
	Position   int `json:"position"`
	SwimlaneID int `json:"swimlane_id"`
}
