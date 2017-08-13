package data

type DataToDisplay struct {
	IsRegistration *bool
	IsTransaction  *bool
	IsPayments     *bool
	OnlyMessage    *bool
	Message        string
	Label          string
	DataHeader     string
	Data           []interface{}
}
