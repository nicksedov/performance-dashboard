package jira

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

type Issue struct {
	Key    string `json:"key"`
	Fields struct {
		Aggregateprogress struct {
			Progress int `json:"progress"`
			Total    int `json:"total"`
		} `json:"aggregateprogress"`
		Aggregatetimeestimate         interface{} `json:"aggregatetimeestimate"`
		Aggregatetimeoriginalestimate interface{} `json:"aggregatetimeoriginalestimate"`
		Aggregatetimespent            interface{} `json:"aggregatetimespent"`
		Assignee Account `json:"assignee"`
		Attachment []interface{} `json:"attachment"`
		Comment    struct {
			Comments   []interface{} `json:"comments"`
			MaxResults int           `json:"maxResults"`
			Self       string        `json:"self"`
			StartAt    int           `json:"startAt"`
			Total      int           `json:"total"`
		} `json:"comment"`
		Components []interface{} `json:"components"`
		Created    string        `json:"created"`
		Creator    Account `json:"creator"`
		Description      interface{}   `json:"description"`
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
		Labels     []interface{} `json:"labels"`
		LastViewed string        `json:"lastViewed"`
		Parent     struct {
			Fields struct {
				Issuetype IssueType `json:"issuetype"`
				Priority  Priority `json:"priority"`
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
			AvatarUrls AvatarUrls `json:"avatarUrls"`
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
		Status Status `json:"status"`
		Statuscategorychangedate string        `json:"statuscategorychangedate"`
		Subtasks                 []interface{} `json:"subtasks"`
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
	Issues []Issue `json:"issues"`
}
