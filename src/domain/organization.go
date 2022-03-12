package domain

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
)

var defaultRoleName *string

type CreateMemberAccountAPI interface {
	CreateAccount(ctx context.Context,
		params *organizations.CreateAccountInput,
		optFns ...func(*organizations.Options)) (*organizations.CreateAccountOutput, error)
}

type AWSOrganization struct {
	AccountName *string // The friendly name of the member account
	Email       *string // The email address of the owner to assign to the new member account
	RoleName    *string
}

func MakeAccount(c context.Context, api CreateMemberAccountAPI, input *organizations.CreateAccountInput) (*organizations.CreateAccountOutput, error) {
	return api.CreateAccount(c, input)
}

func (o *AWSOrganization) CreateInfrastructureAccount() (*organizations.CreateAccountOutput, error) {
	if *o.AccountName == "" || *o.Email == "" {
		log.Fatalf("You must supply account name & email")
	}

	awscfg, e := config.LoadDefaultConfig(context.TODO())
	if e != nil {
		panic("configuration error, " + e.Error())
	}

	client := organizations.NewFromConfig(awscfg)

	if len(*o.RoleName) > 0 {
		defaultRoleName = o.RoleName
	}

	input := &organizations.CreateAccountInput{
		AccountName: o.AccountName,
		Email:       o.Email,
		RoleName:    defaultRoleName,
	}

	result, eorg := MakeAccount(context.TODO(), client, input)

	if eorg != nil {
		log.Fatal(eorg.Error())
	}

	return result, nil
}
