package usecase

import (
	"go-cli-skeleton/internals/customers/models"
	"reflect"
	"testing"
)

func TestItemUsecase_GetAllItems(t *testing.T) {
	tests := []struct {
		name    string
		uc      *CustomerUseCase
		want    []models.Customers
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("ItemUsecase.GetAllItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ItemUsecase.GetAllItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
