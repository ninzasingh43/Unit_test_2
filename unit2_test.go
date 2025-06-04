package main

import (
	"reflect"
	"testing"
)

func TestValidateMinLength(t *testing.T) {
	tests := []struct {
		rule      string
		field     reflect.Value
		fieldName string
		wantErr   bool
	}{
		{"min=2", reflect.ValueOf("A"), "Name", true},
		{"min=2", reflect.ValueOf("Alice"), "Name", false},
		{"min=0", reflect.ValueOf(""), "Name", false},
	}
	for _, tt := range tests {
		err := validateMinLength(
			tt.rule,
			tt.field,
			tt.fieldName,
		)
		if (err != nil) != tt.wantErr {
			t.Errorf("validateMinLength() error = %v,wantErr = %v",
				err,
				tt.wantErr)
		}
	}

}

func TestValidateMaxLength(t *testing.T) {
	tests := []struct {
		rule      string
		field     reflect.Value
		fieldName string
		wantErr   bool
	}{
		{"max=2", reflect.ValueOf("Alice"), "Name", true},
		{"max=5", reflect.ValueOf("Alice"), "Name", false},
		{"max=0", reflect.ValueOf(""), "Name", false},
	}
	for _, tt := range tests {
		err := validateMaxLength(tt.rule, tt.field, tt.fieldName)
		if (err != nil) != tt.wantErr {
			t.Errorf("validateMaxLength() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}

func TestValidateRequired(t *testing.T) {
	tests := []struct {
		field     reflect.Value
		fieldName string
		wantErr   bool
	}{
		{reflect.ValueOf(""), "Email", true},
		{reflect.ValueOf("alice@example.com"), "Email", false},
	}
	for _, tt := range tests {
		err := validateRequired(tt.field, tt.fieldName)
		if (err != nil) != tt.wantErr {
			t.Errorf("validateReuired() error = %v,wantErr %v", err, tt.wantErr)
		}
	}
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		field     reflect.Value
		fieldName string
		wantErr   bool
	}{
		{reflect.ValueOf("aliceexample.com"), "Email", true},
		{reflect.ValueOf("alice@example.com"), "Email", false},
		{reflect.ValueOf(""), "Email", true},
	}
	for _, tt := range tests {
		err := validateEmail(tt.field, tt.fieldName)
		if (err != nil) != tt.wantErr {
			t.Errorf("validateEmail() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}

func TestApplyRule(t *testing.T) {
	tests := []struct {
		rule      string
		field     reflect.Value
		fieldName string
		wantErr   bool
	}{
		{"min=2", reflect.ValueOf("A"), "Name", true},
		{"max=2", reflect.ValueOf("Alice"), "Name", true},
		{"required", reflect.ValueOf(""), "Email", true},
		{"email", reflect.ValueOf("aliceexample.com"), "Email", true},
		{"min=2", reflect.ValueOf("A"), "Name", true},
	}
	for _, tt := range tests {
		err := applyRule(tt.rule, tt.field, tt.fieldName)
		if (err != nil) != tt.wantErr {
			t.Errorf("applyRule() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		user    User
		wantErr bool
	}{
		{User{Name: "Alice", Email: "alice@example.com"}, false},
		{User{Name: "A", Email: "alice@example.com"}, true},
		{User{Name: "Alice", Email: "aliceexample.com"}, true},
		{User{Name: "A", Email: "aliceexample.com"}, true},
		{User{Name: "", Email: ""}, true},
	}

	for _, tt := range tests {
		err := validate(tt.user)
		if (err != nil) != tt.wantErr {
			t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
		}
	}

}

func TestValidateNoTags(t *testing.T) {
	type NoTagUser struct {
		Name  string
		Email string
	}
	tests := []struct {
		user    NoTagUser
		wantErr bool
	}{
		{NoTagUser{Name: "Alice", Email: "alice@example.com"}, false},
		{NoTagUser{Name: "A", Email: "alice@example.com"}, false},
		{NoTagUser{Name: "Alice", Email: "aliceexample.com"}, false},
		{NoTagUser{Name: "A", Email: "aliceexample.com"}, false},
		{NoTagUser{Name: "", Email: ""}, false},
	}
	for _, tt := range tests {
		err := validate(tt.user)
		if (err != nil) != tt.wantErr {
			t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
		}
	}

}
