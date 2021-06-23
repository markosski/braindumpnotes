package notes

type Email string

const PLATFORM_GITHUB = "GITHUB"
const PLATFORM_AWS = "AWS"

type StoreIntegration struct {
	PlatformName string      `json:"platform"`
	Config       interface{} `json:"config"`
}

type GithubIntegration struct {
	Repo  string `json:"repo"`
	Token string `json:"token"`
}

type AwsS3Integration struct {
	Bucket  string `json:"bucket"`
	IamRole string `json:"role"`
}

type Account struct {
	Id               string           `json:"id"`
	UserName         string           `json:"userName"`
	Email            Email            `json:"email"`
	StoreIntegration StoreIntegration `json:"integration"`
}

type AccountStore interface {
	Get(email Email) (Account, error)
	Create(account Account) error
}
