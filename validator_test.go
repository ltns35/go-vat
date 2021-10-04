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
			name: "Belgium VALID",
			args: args{
				values: mocks.BelgiumValidTests,
				countriesList: []countries.Calculer{
					&countries.Belgium,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Belgium INVALID",
			args: args{
				values: mocks.BelgiumInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Belgium,
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
			name: "Bulgaria VALID",
			args: args{
				values: mocks.BulgariaValidTests,
				countriesList: []countries.Calculer{
					&countries.Bulgaria,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Bulgaria INVALID",
			args: args{
				values: mocks.BulgariaInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Bulgaria,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Croatia VALID",
			args: args{
				values: mocks.CroatiaValidTests,
				countriesList: []countries.Calculer{
					&countries.Croatia,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Croatia INVALID",
			args: args{
				values: mocks.CroatiaInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Croatia,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Cyprus VALID",
			args: args{
				values: mocks.CyprusValidTests,
				countriesList: []countries.Calculer{
					&countries.Cyprus,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Cyprus INVALID",
			args: args{
				values: mocks.CyprusInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Cyprus,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Czech Republic VALID",
			args: args{
				values: mocks.CzechRepublicValidTests,
				countriesList: []countries.Calculer{
					&countries.CzechRepublic,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Czech Republic INVALID",
			args: args{
				values: mocks.CzechRepublicInvalidTests,
				countriesList: []countries.Calculer{
					&countries.CzechRepublic,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Denmark VALID",
			args: args{
				values: mocks.DenmarkValidTests,
				countriesList: []countries.Calculer{
					&countries.Denmark,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Denmark INVALID",
			args: args{
				values: mocks.DenmarkInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Denmark,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Estonia VALID",
			args: args{
				values: mocks.EstoniaValidTests,
				countriesList: []countries.Calculer{
					&countries.Estonia,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Estonia INVALID",
			args: args{
				values: mocks.EstoniaInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Estonia,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Finland VALID",
			args: args{
				values: mocks.FinlandValidTests,
				countriesList: []countries.Calculer{
					&countries.Finland,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Finland INVALID",
			args: args{
				values: mocks.FinlandInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Finland,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "France VALID",
			args: args{
				values: mocks.FranceValidTests,
				countriesList: []countries.Calculer{
					&countries.France,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "France INVALID",
			args: args{
				values: mocks.FranceInvalidTests,
				countriesList: []countries.Calculer{
					&countries.France,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Germany VALID",
			args: args{
				values: mocks.GermanyValidTests,
				countriesList: []countries.Calculer{
					&countries.Germany,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Germany INVALID",
			args: args{
				values: mocks.GermanyInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Germany,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Greece VALID",
			args: args{
				values: mocks.GreeceValidTests,
				countriesList: []countries.Calculer{
					&countries.Greece,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Greece INVALID",
			args: args{
				values: mocks.GreeceInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Greece,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Hungary VALID",
			args: args{
				values: mocks.HungaryValidTests,
				countriesList: []countries.Calculer{
					&countries.Hungary,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Hungary INVALID",
			args: args{
				values: mocks.HungaryInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Hungary,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Italy VALID",
			args: args{
				values: mocks.ItalyValidTests,
				countriesList: []countries.Calculer{
					&countries.Italy,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Italy INVALID",
			args: args{
				values: mocks.ItalyInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Italy,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Latvia VALID",
			args: args{
				values: mocks.LatviaValidTests,
				countriesList: []countries.Calculer{
					&countries.Latvia,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Latvia INVALID",
			args: args{
				values: mocks.LatviaInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Latvia,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Luxembourg VALID",
			args: args{
				values: mocks.LuxembourgValidTests,
				countriesList: []countries.Calculer{
					&countries.Luxembourg,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Luxembourg INVALID",
			args: args{
				values: mocks.LuxembourgInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Luxembourg,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Malta VALID",
			args: args{
				values: mocks.MaltaValidTests,
				countriesList: []countries.Calculer{
					&countries.Malta,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Malta INVALID",
			args: args{
				values: mocks.MaltaInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Malta,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Norway VALID",
			args: args{
				values: mocks.NorwayValidTests,
				countriesList: []countries.Calculer{
					&countries.Norway,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Norway INVALID",
			args: args{
				values: mocks.NorwayInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Norway,
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
			name: "Serbia VALID",
			args: args{
				values: mocks.SerbiaValidTests,
				countriesList: []countries.Calculer{
					&countries.Serbia,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Serbia INVALID",
			args: args{
				values: mocks.SerbiaInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Serbia,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Slovakia VALID",
			args: args{
				values: mocks.SlovakiaValidTests,
				countriesList: []countries.Calculer{
					&countries.Slovakia,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Slovakia INVALID",
			args: args{
				values: mocks.SlovakiaInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Slovakia,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Slovenia VALID",
			args: args{
				values: mocks.SloveniaValidTests,
				countriesList: []countries.Calculer{
					&countries.Slovenia,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Slovenia INVALID",
			args: args{
				values: mocks.SloveniaInvalidTests,
				countriesList: []countries.Calculer{
					&countries.Slovenia,
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
