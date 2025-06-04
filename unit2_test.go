package main

import (
	"reflect"
	"testing"
)

func TestValidateMinLength(t *testing.T){
		tests := []struct {
			rule string
			field reflect.Value
			fieldName string 
			wantErr bool
		}{
			{"min=2",reflect.ValueOf("A"), "Name", true},
			{"min=2",reflect.ValueOf("Alice"), "Name", false},
			{"min=0",reflect.ValueOf(""), "Name", false},
		}
		for _, tt := range tests {
			err := validateMinLength(
				tt.rule, 
				tt.field, 
				tt.fieldName, 
			)
			if (err !=nil ) != tt.wantErr {
				 t.Errorf( "validateMinLength() error = %v,wantErr = %v",
				err,
			tt.wantErr)
			}
		}

}


func TestValidateMaxLength(t *testing.T) {
	tests := []struct {
		rule string 
		field reflect.Value
		fieldName string
		wantErr bool

	}{
		{"max=2", reflect.ValueOf("Alice"), "Name",true},
		{"max=5", reflect.ValueOf("Alice"), "Name",false},
		{"max=0", reflect.ValueOf(""), "Name",false},
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
		field reflect.Value
		fieldName string
		wantErr bool
	}
	{
		{reflect.ValueOf(""), "Email", true},
		{reflect.ValueOf("alice@example.com"),"Email",false},
	}
	for _, tt  := range tests {
		err := validateRequired(tt.field, tt.fieldName)
		if (err != nil) != tt.wantErr {
			t.Errorf("validateReuired() error = %v,wantErr %v", err, tt.wantErr)
		}
	}
}