package exporter

//SocialMedia Function
type SocialMedia interface {
	Feed() []string
	Fame() int
}

//Facebook exporter
type Facebook struct {
	URL     string
	Name    string
	NoOfFriends int
}

// Feed shows the latest posts on Facebook
func (f *Facebook) Feed() []string {
	return []string{
		"Facebook feeds",
		"Hey, here's my cool new selfie",
	}
}

// Fame describes how popular the user is by the total number of friends
func (f *Facebook) Fame() int {
	return f.NoOfFriends
}




//LinkedIn exporter
type LinkedIn struct {
	URL         string
	Name        string
	NoOfConnections int
}

// Feed shows latest posts on LinkedIn
func (l *LinkedIn) Feed() []string {
	return []string{
		"LinkedIn feeds",
		"Hey, here's my cool new selfie",
	}
}

// Fame describes how popular the user is by the total number of connections
func (l *LinkedIn) Fame() int {
	return l.NoOfConnections
}




//Twitter exporter
type Twitter struct {
	URL       string
	Username  string
	NoOfFollowers int
}

// Feed shows latest posts on Twitter
func (t *Twitter) Feed() []string {
	return []string{
		"Twitter feeds",
		"Hey, here's my cool new selfie",
	}
}

// Fame describes how popular the user is by the total number of followers
func (t *Twitter) Fame() int {
	return t.NoOfFollowers
}
