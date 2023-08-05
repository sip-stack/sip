package registry

type CreateAnswers struct {
	Project       string
	Backend       string
	Frontend      string
	Proxy         string
	AgreeCreation bool `survey:"agree"`
}
