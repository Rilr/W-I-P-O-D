package manage

import "time"

type TimeSheet []struct {
	ID     int `json:"id"`
	Member struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string `json:"member_href"`
		} `json:"_info"`
	} `json:"member"`
	Year      int       `json:"year"`
	Period    int       `json:"period"`
	DateStart time.Time `json:"dateStart"`
	DateEnd   time.Time `json:"dateEnd"`
	Status    string    `json:"status"`
	Hours     float64   `json:"hours"`
	Deadline  time.Time `json:"deadline"`
	Info      struct {
		LastUpdated time.Time `json:"lastUpdated"`
		UpdatedBy   string    `json:"updatedBy"`
	} `json:"_info"`
}