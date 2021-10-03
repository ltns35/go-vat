package vat

import (
	"fmt"
	"testing"

	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/mocks"
)

func TestCheckVAT(t *testing.T) {

	type args struct {
		testName      string
		countriesList []countries.Calculer
		values        []string
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Andorra VALID",
			args: args{
				values: mocks.AndorraValidTests,
				countriesList: []countries.Calculer{
					&countries.Andorra,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Andorra INVALID",
			args: args{
				values: mocks.AndorraInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Andorra,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Austria VALID",
			args: args{
				values: mocks.AustriaValidTests,
				countriesList: []countries.Calculer{
					&countries.Austria,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Austria INVALID",
			args: args{
				values: mocks.AustriaInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Austria,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Brazil VALID",
			args: args{
				values: mocks.BrazilValidTests,
				countriesList: []countries.Calculer{
					&countries.Brazil,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Brazil INVALID",
			args: args{
				values: mocks.BrazilInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Brazil,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Portugal VALID",
			args: args{
				values: mocks.PortugalValidTests,
				countriesList: []countries.Calculer{
					&countries.Portugal,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Portugal INVALID",
			args: args{
				values: mocks.PortugalInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Portugal,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Spain VALID",
			args: args{
				values: mocks.SpainValidTests,
				countriesList: []countries.Calculer{
					&countries.Spain,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Spain INVALID",
			args: args{
				values: mocks.SpainInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Spain,
				},
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		for i, test := range tt.args.values {

			t.Run(
				fmt.Sprintf("%s nº %d: %s", tt.name, i, test), func(t *testing.T) {
					got, err := CheckVAT(test, tt.args.countriesList...)
					if (err != nil) != tt.wantErr {
						t.Errorf("CheckVAT() error = %v, wantErr %v", err, tt.wantErr)
						return
					}
					if got != tt.want {
						t.Errorf("CheckVAT() got = %v, want %v", got, tt.want)
					}
				},
			)

		}

	}
}
