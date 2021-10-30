package usecase

import (
	"errors"
	"gomocking/mocks"
	"gomocking/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/undefinedlabs/go-mpatch"
)

func TestUsecase_Addition(t *testing.T) {

	// substitute interface

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocks.NewMockDatabaseInterface(mockCtrl)

	type fields struct {
		Database repository.DatabaseInterface
	}
	type args struct {
		number1 string
		number2 string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
		patch   func()
	}{
		{
			name: "Succesfully Insert to Database",
			fields: fields{
				Database: mockDB,
			},
			args: args{
				number1: "50",
				number2: "20",
			},
			want:    70,
			wantErr: false,
			patch: func() {

				mockDB.
					EXPECT().
					Insert(gomock.Eq(70)).
					Return(nil)

			},
		},
		{
			name: "Failed to convert input to int",
			fields: fields{
				Database: mockDB,
			},
			args: args{
				number1: "0",
				number2: "0",
			},
			want:    0,
			wantErr: true,
			patch: func() {

				// monkey patch

				var monkeyPatch *mpatch.Patch
				var err error

				monkeyPatch, err = mpatch.PatchMethod(convertStringToInt, func(s string) (int, error) {
					defer monkeyPatch.Unpatch()
					return 0, errors.New("")
			
				})
				if err != nil {
					t.Fatal(err)
				}

				// we dont need to mock Insert method because it is unreachable in this case

			},
		},
		{
			name: "Failed Insert to Database",
			fields: fields{
				Database: mockDB,
			},
			args: args{
				number1: "50",
				number2: "20",
			},
			want:    0,
			wantErr: true,
			patch: func() {

				mockDB.
					EXPECT().
					Insert(gomock.Eq(70)).
					Return(errors.New(""))

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Usecase{
				Database: tt.fields.Database,
			}

			tt.patch()

			got, err := u.Addition(tt.args.number1, tt.args.number2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.Addition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Usecase.Addition() = %v, want %v", got, tt.want)
			}
		})
	}
}
