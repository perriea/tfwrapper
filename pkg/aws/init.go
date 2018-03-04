package authAWS

var (
	err error
)

// AWSConfig file struct
type AWSConfig struct {
	profile string
	role    string
	mfa     string
}
