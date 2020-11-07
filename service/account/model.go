package account

import "github.com/RedisLabs/rediscloud-go-api/internal"

type paymentMethods struct {
	PaymentMethods []*PaymentMethod `json:"paymentMethods,omitempty"`
}

func (o paymentMethods) String() string {
	return internal.ToString(o)
}

type PaymentMethod struct {
	ID                 *int    `json:"id,omitempty"`
	Type               *string `json:"type,omitempty"`
	CreditCardEndsWith *int    `json:"creditCardEndsWith,omitempty"`
	ExpirationMonth    *int    `json:"expirationMonth"`
	ExpirationYear     *int    `json:"expirationYear"`
}

func (o PaymentMethod) String() string {
	return internal.ToString(o)
}

type regions struct {
	Regions []*Region `json:"regions,omitempty"`
}

func (o regions) String() string {
	return internal.ToString(o)
}

type Region struct {
	Name     *string    `json:"name,omitempty"`
	Provider *string `json:"provider,omitempty"`
}
