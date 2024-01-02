package usecase

import (
	"go-cli-skeleton/internals/yourmodule/models"
	"reflect"
	"testing"
)

func TestItemUsecase_GetAllItems(t *testing.T) {
	tests := []struct {
		name    string
		uc      *ItemUsecase
		want    []models.Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetAllItems()
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
