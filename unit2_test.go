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
			err := ValidateMinLength(
				tt.rule, 
				tt.field, 
				tt.fieldName, 
			)
			if (err !=nil ) != tt.wantErr {
				 t.Errorf( "validateMinLength() error = %v,
				 wantErr = %v",
				err,
			tt.wantErr)
			}
		}

}