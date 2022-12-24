package user

type ResponseJson struct {
	ID                 int    `json:"id"`
	Email              string `json:"email"`
	Name               string `json:"name"`
	SubscriptionStatus string `json:"subscription_status"`
}
