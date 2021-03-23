// Code generated by entc, DO NOT EDIT.

package ent

import (
	"sqlent/ent/schema"
	"sqlent/ent/user"
	"sqlent/ent/usercount"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[1].Descriptor()
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = userDescPhone.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescPID is the schema descriptor for p_id field.
	userDescPID := userFields[3].Descriptor()
	// user.DefaultPID holds the default value on creation for the p_id field.
	user.DefaultPID = userDescPID.Default.(int)
	// userDescInvitedCode is the schema descriptor for invited_code field.
	userDescInvitedCode := userFields[4].Descriptor()
	// user.InvitedCodeValidator is a validator for the "invited_code" field. It is called by the builders before save.
	user.InvitedCodeValidator = userDescInvitedCode.Validators[0].(func(string) error)
	// userDescCreated is the schema descriptor for created field.
	userDescCreated := userFields[5].Descriptor()
	// user.DefaultCreated holds the default value on creation for the created field.
	user.DefaultCreated = userDescCreated.Default.(int64)
	// userDescUpdated is the schema descriptor for updated field.
	userDescUpdated := userFields[6].Descriptor()
	// user.DefaultUpdated holds the default value on creation for the updated field.
	user.DefaultUpdated = userDescUpdated.Default.(int64)
	// userDescDeleted is the schema descriptor for deleted field.
	userDescDeleted := userFields[7].Descriptor()
	// user.DefaultDeleted holds the default value on creation for the deleted field.
	user.DefaultDeleted = userDescDeleted.Default.(int64)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(int64) error)
	usercountFields := schema.UserCount{}.Fields()
	_ = usercountFields
	// usercountDescSelfBuy is the schema descriptor for self_buy field.
	usercountDescSelfBuy := usercountFields[0].Descriptor()
	// usercount.DefaultSelfBuy holds the default value on creation for the self_buy field.
	usercount.DefaultSelfBuy = usercountDescSelfBuy.Default.(float64)
	// usercountDescInviteBuy is the schema descriptor for invite_buy field.
	usercountDescInviteBuy := usercountFields[1].Descriptor()
	// usercount.DefaultInviteBuy holds the default value on creation for the invite_buy field.
	usercount.DefaultInviteBuy = usercountDescInviteBuy.Default.(float64)
	// usercountDescLevel is the schema descriptor for level field.
	usercountDescLevel := usercountFields[2].Descriptor()
	// usercount.DefaultLevel holds the default value on creation for the level field.
	usercount.DefaultLevel = usercountDescLevel.Default.(int)
	// usercountDescCreated is the schema descriptor for created field.
	usercountDescCreated := usercountFields[3].Descriptor()
	// usercount.DefaultCreated holds the default value on creation for the created field.
	usercount.DefaultCreated = usercountDescCreated.Default.(int64)
	// usercountDescUpdated is the schema descriptor for updated field.
	usercountDescUpdated := usercountFields[4].Descriptor()
	// usercount.DefaultUpdated holds the default value on creation for the updated field.
	usercount.DefaultUpdated = usercountDescUpdated.Default.(int64)
	// usercountDescDeleted is the schema descriptor for deleted field.
	usercountDescDeleted := usercountFields[5].Descriptor()
	// usercount.DefaultDeleted holds the default value on creation for the deleted field.
	usercount.DefaultDeleted = usercountDescDeleted.Default.(int64)
}