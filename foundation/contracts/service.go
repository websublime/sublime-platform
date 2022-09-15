package contracts

type ServiceItem struct {
	Name     string
	Provider interface{}
}

type Services []*ServiceItem
