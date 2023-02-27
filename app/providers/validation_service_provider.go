package providers

import (
	"gopkg.in/go-mixed/framework.v1/contracts/validation"
	"gopkg.in/go-mixed/framework.v1/facades"
	"gopkg.in/go-mixed/framework.v1/facades/log"
)

type ValidationServiceProvider struct {
}

func (receiver *ValidationServiceProvider) Register() {

}

func (receiver *ValidationServiceProvider) Boot() {
	if err := facades.Validation.AddRules(receiver.rules()); err != nil {
		log.Errorf("add rules error: %+v", err)
	}
}

func (receiver *ValidationServiceProvider) rules() []validation.Rule {
	return []validation.Rule{}
}
