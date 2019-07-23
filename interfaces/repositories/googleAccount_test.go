package repositories

import (
	"fmt"
	"reflect"
	"serverless-functions-go/domain/models/cedar"
	"serverless-functions-go/domain/repositories"
	"serverless-functions-go/infrastructure/config"
	"serverless-functions-go/infrastructure/database"
	"testing"
)

var (
	googleAccountRepo repositories.GoogleAccountRepository
)

func init() {
	appConf, err := config.New("os")
	if err != nil {
		panic(err)
	}
	dbConf := database.PGDatabaseConfig{
		Host:     appConf.SqlHost,
		Port:     appConf.SqlPort,
		User:     appConf.SqlUser,
		Password: appConf.SqlPassword,
		Dbname:   appConf.SqlDb,
	}
	fmt.Printf("dbConf = %+v\n", dbConf)
	db, err := database.ConnectPGDatabase(&dbConf)
	if err != nil {
		panic(err)
	}
	googleAccountRepo = NewGoogleAccountRepository(db)
}

func TestGoogleAccountRepository_FindByClientCustomerIds(t *testing.T) {
	type args struct {
		ids []int
	}
	tests := []struct {
		name    string
		args    args
		want    []cedar.GoogleAccount
		wantErr bool
	}{
		{
			name:    "should get client customer ids correctly",
			args:    args{ids: []int{9258066191}},
			want:    make([]cedar.GoogleAccount, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := googleAccountRepo.FindByClientCustomerIds(tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoogleAccountRepository.FindByClientCustomerIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GoogleAccountRepository.FindByClientCustomerIds() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
