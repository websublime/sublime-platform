package contracts

type ServiceItem struct {
	Name     string
	Provider interface{}
}

type Services []*ServiceItem

type Service struct {
	Item *ServiceItem
}
