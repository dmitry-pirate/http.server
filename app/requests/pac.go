package requests

type PacRequest struct {
	DeviceToken string `json:"device_token"`
	Country     string `json:"country"`
}
