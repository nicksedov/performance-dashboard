package jira

import "time"

type IssueType  struct {
	AvatarID       int    `json:"avatarId"`
	Description    string `json:"description"`
	EntityID       string `json:"entityId"`
	HierarchyLevel int    `json:"hierarchyLevel"`
	IconURL        string `json:"iconUrl"`
	ID             string `json:"id"`
	Name           string `json:"name"`
	Self           string `json:"self"`
	Subtask        bool   `json:"subtask"`
}

type Priority struct {
	IconURL string `json:"iconUrl"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Self    string `json:"self"`
}

type Status struct {
	Description    string `json:"description"`
	IconURL        string `json:"iconUrl"`
	ID             string `json:"id"`
	Name           string `json:"name"`
	Self           string `json:"self"`
	StatusCategory struct {
		ColorName string `json:"colorName"`
		ID        int    `json:"id"`
		Key       string `json:"key"`
		Name      string `json:"name"`
		Self      string `json:"self"`
	} `json:"statusCategory"`
}

type Attachment struct {
	Author Account  `json:"author"`
	Content  string `json:"content"`
	Created  string `json:"created"`
	Filename string `json:"filename"`
	ID       string `json:"id"`
	MimeType string `json:"mimeType"`
	Self     string `json:"self"`
	Size     int    `json:"size"`
}

type Comment struct {
	Author       Account `json:"author"`
	Body         string `json:"body"`
	Created      time.Time `json:"created"`
	ID           string `json:"id"`
	JsdPublic    bool   `json:"jsdPublic"`
	Self         string `json:"self"`
	UpdateAuthor Account `json:"updateAuthor"`
	Updated      time.Time `json:"updated"`
}

type Issue struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
	Expand string `json:"expand"`
	Fields struct {
		Aggregateprogress struct {
			Progress int `json:"progress"`
			Total    int `json:"total"`
		} `json:"aggregateprogress"`
		Aggregatetimeestimate         interface{} `json:"aggregatetimeestimate"`
		Aggregatetimeoriginalestimate interface{} `json:"aggregatetimeoriginalestimate"`
		Aggregatetimespent            interface{} `json:"aggregatetimespent"`
		Assignee Account `json:"assignee"`
		Attachment    []Attachment `json:"attachment"`
		ClosedSprints []Sprint `json:"closedSprints"`
		IssueComment struct {
			Comments   []Comment     `json:"comments"`
			MaxResults int           `json:"maxResults"`
			Self       string        `json:"self"`
			StartAt    int           `json:"startAt"`
			Total      int           `json:"total"`
		} `json:"comment"`
		Components []interface{} `json:"components"`
		Created    string        `json:"created"`
		Creator    Account `json:"creator"`
		Description      string   `json:"description"`
		Duedate          interface{}   `json:"duedate"`
		Environment      interface{}   `json:"environment"`
		Epic             interface{}   `json:"epic"`
		FixVersions      []interface{} `json:"fixVersions"`
		Flagged          bool          `json:"flagged"`
		Issuelinks       []interface{} `json:"issuelinks"`
		Issuerestriction struct {
			Issuerestrictions struct {
			} `json:"issuerestrictions"`
			ShouldDisplay bool `json:"shouldDisplay"`
		} `json:"issuerestriction"`
		Issuetype IssueType `json:"issuetype"`
		Labels     []string `json:"labels"`
		LastViewed string        `json:"lastViewed"`
		Parent     struct {
			Fields struct {
				Issuetype struct {
					AvatarID       int    `json:"avatarId"`
					Description    string `json:"description"`
					EntityID       string `json:"entityId"`
					HierarchyLevel int    `json:"hierarchyLevel"`
					IconURL        string `json:"iconUrl"`
					ID             string `json:"id"`
					Name           string `json:"name"`
					Self           string `json:"self"`
					Subtask        bool   `json:"subtask"`
				} `json:"issuetype"`
				Priority struct {
					IconURL string `json:"iconUrl"`
					ID      string `json:"id"`
					Name    string `json:"name"`
					Self    string `json:"self"`
				} `json:"priority"`
				Status Status `json:"status"`
				Summary string `json:"summary"`
			} `json:"fields"`
			ID   string `json:"id"`
			Key  string `json:"key"`
			Self string `json:"self"`
		} `json:"parent"`
		Priority Priority `json:"priority"`
		Progress struct {
			Progress int `json:"progress"`
			Total    int `json:"total"`
		} `json:"progress"`
		Project struct {
			AvatarUrls struct {
				One6X16   string `json:"16x16"`
				Two4X24   string `json:"24x24"`
				Three2X32 string `json:"32x32"`
				Four8X48  string `json:"48x48"`
			} `json:"avatarUrls"`
			ID             string `json:"id"`
			Key            string `json:"key"`
			Name           string `json:"name"`
			ProjectTypeKey string `json:"projectTypeKey"`
			Self           string `json:"self"`
			Simplified     bool   `json:"simplified"`
		} `json:"project"`
		Reporter Account `json:"reporter"`
		Resolution     interface{} `json:"resolution"`
		Resolutiondate interface{} `json:"resolutiondate"`
		Security       interface{} `json:"security"`
		Sprint         Sprint `json:"sprint"`
		Status         Status `json:"status"`
		Statuscategorychangedate string        `json:"statuscategorychangedate"`
		Subtasks                 []Issue       `json:"subtasks"`
		Summary                  string        `json:"summary"`
		Timeestimate             interface{}   `json:"timeestimate"`
		Timeoriginalestimate     interface{}   `json:"timeoriginalestimate"`
		Timespent                interface{}   `json:"timespent"`
		Timetracking             struct {
		} `json:"timetracking"`
		Updated  string        `json:"updated"`
		Versions []interface{} `json:"versions"`
		Votes    struct {
			HasVoted bool   `json:"hasVoted"`
			Self     string `json:"self"`
			Votes    int    `json:"votes"`
		} `json:"votes"`
		Watches struct {
			IsWatching bool   `json:"isWatching"`
			Self       string `json:"self"`
			WatchCount int    `json:"watchCount"`
		} `json:"watches"`
		Worklog struct {
			MaxResults int           `json:"maxResults"`
			StartAt    int           `json:"startAt"`
			Total      int           `json:"total"`
			Worklogs   []interface{} `json:"worklogs"`
		} `json:"worklog"`
		Workratio int `json:"workratio"`
	} `json:"fields"`
}

type Issues struct {
	Expand      string `json:"expand"`
	Issues     []Issue `json:"issues"`
	MaxResults     int `json:"maxResults"`
	StartAt        int `json:"startAt"`
	Total          int `json:"total"`
}