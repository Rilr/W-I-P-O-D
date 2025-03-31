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

type TimeEntries []struct {
	ID      int `json:"id"`
	Company struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			CompanyHref string `json:"company_href"`
			MobileGUID  string `json:"mobileGuid"`
		} `json:"_info"`
	} `json:"company"`
	CompanyType  string `json:"companyType"`
	ChargeToID   int    `json:"chargeToId"`
	ChargeToType string `json:"chargeToType"`
	Member       struct {
		ID            int     `json:"id"`
		Identifier    string  `json:"identifier"`
		Name          string  `json:"name"`
		DailyCapacity float64 `json:"dailyCapacity"`
		Info          struct {
			MemberHref string    `json:"member_href"`
			ImageHref  string `json:"image_href"`
		} `json:"_info"`
	} `json:"member"`
	LocationID        int    `json:"locationId"`
	BusinessUnitID    int    `json:"businessUnitId"`
	BusinessGroupDesc string `json:"businessGroupDesc"`
	Location          struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			LocationHref string `json:"location_href"`
		} `json:"_info"`
	} `json:"location"`
	Department struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			DepartmentHref string `json:"department_href"`
		} `json:"_info"`
	} `json:"department"`
	WorkType struct {
		ID              int    `json:"id"` // ID of Must Change "62"
		Name            string `json:"name"` // 62 Name is "Must Change"
		UtilizationFlag bool   `json:"utilizationFlag"`
		Info            struct {
			WorkTypeHref string `json:"workType_href"`
		} `json:"_info"`
	} `json:"workType"`
	WorkRole struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			WorkRoleHref string `json:"workRole_href"`
		} `json:"_info"`
	} `json:"workRole"`
	Territory                  string    `json:"territory"`
	TimeStart                  time.Time `json:"timeStart"`
	TimeEnd                    time.Time `json:"timeEnd"`
	HoursDeduct                float64   `json:"hoursDeduct"`
	ActualHours                float64   `json:"actualHours"`
	BillableOption             string    `json:"billableOption"`
	Notes                      string    `json:"notes,omitempty"`
	AddToDetailDescriptionFlag bool      `json:"addToDetailDescriptionFlag"`
	AddToInternalAnalysisFlag  bool      `json:"addToInternalAnalysisFlag"`
	AddToResolutionFlag        bool      `json:"addToResolutionFlag"`
	HoursBilled                float64   `json:"hoursBilled"`
	InvoiceHours               float64   `json:"invoiceHours"`
	HourlyCost                 string    `json:"hourlyCost"`
	EnteredBy                  string    `json:"enteredBy"`
	DateEntered                time.Time `json:"dateEntered"`
	MobileGUID                 string    `json:"mobileGuid"`
	HourlyRate                 float64   `json:"hourlyRate"`
	AgreementHours             float64   `json:"agreementHours"`
	AgreementAmount            float64   `json:"agreementAmount"`
	AgreementAdjustment        float64   `json:"agreementAdjustment"`
	Adjustment                 float64   `json:"adjustment"`
	InvoiceReady               int       `json:"invoiceReady"`
	TimeSheet                  struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			TimeSheetHref string `json:"timeSheet_href"`
		} `json:"_info"`
	} `json:"timeSheet"`
	Status                string  `json:"status"`
	InvoiceFlag           bool    `json:"invoiceFlag"`
	ExtendedInvoiceAmount float64 `json:"extendedInvoiceAmount"`
	LocationName          string  `json:"locationName"`
	TaxCode               struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			TaxCodeHref string `json:"taxCode_href"`
		} `json:"_info"`
	} `json:"taxCode"`
	Info struct {
		LastUpdated time.Time `json:"lastUpdated"`
		UpdatedBy   string    `json:"updatedBy"`
	} `json:"_info"`
	Agreement struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Type           string `json:"type"`
		ChargeFirmFlag bool   `json:"chargeFirmFlag"`
		Info           struct {
			AgreementHref string `json:"agreement_href"`
			TypeID        string `json:"typeId"`
		} `json:"_info"`
	} `json:"agreement,omitempty"`
	AgreementType     string `json:"agreementType,omitempty"`
	EmailResourceFlag bool   `json:"emailResourceFlag,omitempty"`
	EmailContactFlag  bool   `json:"emailContactFlag,omitempty"`
	EmailCcFlag       bool   `json:"emailCcFlag,omitempty"`
	EmailCc           string `json:"emailCc,omitempty"`
	Ticket            struct {
		ID      int    `json:"id"`
		Summary string `json:"summary"`
		Info    struct {
			TicketHref    string `json:"ticket_href"`
			BillingMethod string `json:"billingMethod"`
			Status        string `json:"status"`
		} `json:"_info"`
	} `json:"ticket,omitempty"`
	TicketBoard   string `json:"ticketBoard,omitempty"`
	TicketStatus  string `json:"ticketStatus,omitempty"`
	TicketType    string `json:"ticketType,omitempty"`
	TicketSubType string `json:"ticketSubType,omitempty"`
	Invoice       struct {
		ID             int       `json:"id"`
		Identifier     string    `json:"identifier"`
		InvoiceDate    time.Time `json:"invoiceDate"`
		ChargeFirmFlag bool      `json:"chargeFirmFlag"`
		Info           struct {
			InvoiceHref string `json:"invoice_href"`
		} `json:"_info"`
	} `json:"invoice,omitempty"`
	InternalNotes string `json:"internalNotes,omitempty"`
}