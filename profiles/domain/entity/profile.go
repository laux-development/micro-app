package entity

// Profile contains basic information about a vendor profile
// Gender and Race are options
// Gender = Male | Female | Other | Prefer not to say
// Race = Caucasian, Black, Asian etc
type Profile struct {
	Address    string
	Company    string
	DOB        string
	FirstName  string
	Gender     string
	GithubName string
	JobRole    string
	LastName   string
}
