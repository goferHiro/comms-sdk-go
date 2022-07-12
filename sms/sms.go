package sms

type provider struct{
  apiKey string
  apiPath string
}

type Request struct {
	From            string `json:"from"`
	To              string `json:"to"`
	TemplateID      string `json:"template_id"`
	Message         string `json:"message"`
	Backup          bool   `json:"backup"`
	ProviderVersion string `json:"provider_version"`
	ProviderName    string `json:"provider_name"`
	Webhook         string `json:"webhook"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (p *provider) Send(){
    u, err := url.Parse(apiPath)
}