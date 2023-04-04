package internal

import "go.mongodb.org/mongo-driver/bson/primitive"

type HomeSchoolRegisterIn struct {
	ParentGuardians []*ParentGuardian `json:"parent_guardians"`
	Students        []*Student        `json:"students"`
}

type HomeSchoolRegisterOut struct {
	ParentGuardians []*ParentGuardianOut `json:"parent_guardians"`
}

type ParentGuardianOut struct {
	AccountId string `json:"account_id" validate:"required"`
	UserId    string `json:"user_id" validate:"required"`
	Username  string `json:"username" validate:"required"`
}

type StudentOut struct {
	ParentAccountId string `json:"parent_account_id" validate:"required"`
}

type ParentGuardian struct {
	AccountType              string             `json:"account_type" bson:"account_type"`
	County                   string             `json:"county" bson:"county" validate:"required"`
	CreatedAt                primitive.DateTime `json:"created_at" bson:"created_at"`
	FirstName                string             `json:"first_name" bson:"first_name" validate:"required"`
	LastName                 string             `json:"last_name" bson:"last_name" validate:"required"`
	HomeSchoolRegistrationId string             `json:"home_school_registration_id" bson:"home_school_registration_id" validate:"required"`
	IsActive                 bool               `json:"is_active" bson:"is_active"`
	IsMarkedForDeletion      bool               `json:"is_marked_for_deletion" bson:"is_marked_for_deletion"`
	IsPending                bool               `json:"is_pending" bson:"is_pending"`
	Phone                    string             `json:"phone" bson:"phone" validate:"required"`
	Pin                      string             `json:"pin" bson:"pin" validate:"required"`
	UpdatedAt                primitive.DateTime `json:"updated_at" bson:"updated_at"`
	Username                 string             `json:"username" bson:"username" validate:"required,email"`
	Zipcode                  string             `json:"zipcode" bson:"zipcode"`
}

type Student struct {
	AccountType         string             `json:"account_type"`
	County              string             `json:"county" validate:"required"`
	CreatedAt           primitive.DateTime `json:"created_at"`
	FirstName           string             `json:"first_name" validate:"required"`
	GradeTypeRequested  uint8              `json:"grade_type_requested" validate:"required"`
	LastName            string             `json:"last_name" validate:"required"`
	IsActive            bool               `json:"is_active"`
	IsMarkedForDeletion bool               `json:"is_marked_for_deletion"`
	IsPending           bool               `json:"is_pending"`
	Pin                 string             `json:"pin" validate:"required"`
	PrincipalId         string             `json:"principal_id"`
	UpdatedAt           primitive.DateTime `json:"updated_at"`
	Username            string             `json:"username"`
}

type AccountTypeIn struct {
	Id string `json:"id"`
}

type AccountTypeOut struct {
	AccountType string `json:"account_type" bson:"account_type"`
	Id          string `json:"id" bson:"_id"`
}
