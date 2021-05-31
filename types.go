package mschartgen

// TODO: currently unused
type People struct {
	Persons []Person `json:"value"`
}


// phase 1: these types are binded to the msgraph api (data fetch stage)
type Organization struct {
	Data []OrgMetaData `json:"value"`
}

type OrgMetaData struct {
	DisplayName string `json:"displayName"`
}

type Person struct {
	Id            string `json:"id"`
	Name          string `json:"displayName"`
	Title         string `json:"jobTitle"`
	DirectReports []Person
}

type DirectReports struct {
	Data []Person `json:"value"`
}


// phase 2: these types are the native data structures
type Organisation struct {
	Name string
	Head Member
}

type Member struct {
	Id            string
	Name          string
	Title         string
	DirectReports []Member
}


// phase 3: these types are for post-processing to orgchart js
type OrgChart struct {
	Data []Child
}

type Child struct {
	Name     string
	Title    string
	Children []Child
}
