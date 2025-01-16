package Structs

/*
Here are the key pointer operations in Go:

& (Address-of operator):

	goCopyx := 42
	ptr := &x    // Get memory address of x

* has two uses:

	Declaration - Declaring a pointer type:

		goCopyvar ptr *int   // ptr is a pointer to int

	Dereferencing - Getting value at pointer's address:

		goCopyx := 42
		ptr := &x
		value := *ptr  // Get value that ptr points to (42)
		*ptr = 100     // Change value at ptr's address
*/

type ProfileMaker interface {
	UpdateProfilePicture(ProfilePicture)
	CheckDuplicateProfile(SlackProfile) bool
}

type SlackProfile struct {
	Name          string
	Username      string
	Designation   string
	ContactNumber string
	/* Embedding ProfilePicture struct or Anonymous field,
	   ProfilePicture is promoted to SlackProfile hence can be accessed directly
	   e.g. slackProfile.ImageName
	*/
	ProfilePicture
}

type ProfilePicture struct {
	ImageName string
	ImagePath string
}

func (s *SlackProfile) UpdateProfilePicture(pp ProfilePicture) {
	s.ImageName = pp.ImageName
	s.ImagePath = pp.ImagePath
}

func (s *SlackProfile) CheckDuplicateProfile(sp SlackProfile) bool {
	return *s == sp
}

func NewSlackProfile(Name, Username, Designation, ContactNumber, ImageName, ImagePath string) *SlackProfile {
	newSlackProfile := &SlackProfile{
		Name: Name, Username: Username,
		Designation:    Designation,
		ContactNumber:  ContactNumber,
		ProfilePicture: ProfilePicture{ImageName: ImageName, ImagePath: ImagePath},
	}
	return newSlackProfile
}

func (s *SlackProfile) UpdateSlackProfile(Name, Username, Designation, ContactNumber, ImageName, ImagePath string) {
	s.Name = Name
	s.Username = Username
	s.Designation = Designation
	s.ContactNumber = ContactNumber
	s.ImageName = ImageName
	s.ImagePath = ImagePath
}

// func main() {
// 	slackProfile := NewSlackProfile("John Doe", "johndoe", "Software Engineer", "1234567890", "john.png", "/images/john.png")
// 	fmt.Println("New Slack Profile:", slackProfile)

// 	slackProfile.UpdateSlackProfile("Jane Doe", "janedoe", "Software Engineer", "9876543210", "jane.png", "/images/jane.png")
// 	fmt.Println("Updated Slack Profile:", slackProfile)

// }
