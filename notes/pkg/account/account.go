package account

type Email string

type StoreIntegration struct {
	PlatformName string
	Integration  interface{}
}

type GithubIntegration struct {
	Repo  string
	Token string
}

type AwsS3Integration struct {
	Bucket  string
	IamRole string
}

type User struct {
	Name  string
	Email Email
}

type Account struct {
	Email       Email
	Integration StoreIntegration
}
