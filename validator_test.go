package vat

import (
	"fmt"
	"testing"

	"github.com/ltns35/go-vat/countries"
	"github.com/ltns35/go-vat/countries/albania"
	"github.com/ltns35/go-vat/countries/andorra"
	"github.com/ltns35/go-vat/countries/argentina"
	"github.com/ltns35/go-vat/countries/austria"
	"github.com/ltns35/go-vat/countries/belarus"
	"github.com/ltns35/go-vat/countries/belgium"
	"github.com/ltns35/go-vat/countries/brazil"
	"github.com/ltns35/go-vat/countries/bulgaria"
	"github.com/ltns35/go-vat/countries/croatia"
	"github.com/ltns35/go-vat/countries/cyprus"
	"github.com/ltns35/go-vat/countries/czech_republic"
	"github.com/ltns35/go-vat/countries/denmark"
	"github.com/ltns35/go-vat/countries/estonia"
	"github.com/ltns35/go-vat/countries/finland"
	"github.com/ltns35/go-vat/countries/france"
	"github.com/ltns35/go-vat/countries/germany"
	"github.com/ltns35/go-vat/countries/greece"
	"github.com/ltns35/go-vat/countries/hungary"
	"github.com/ltns35/go-vat/countries/ireland"
	"github.com/ltns35/go-vat/countries/italy"
	"github.com/ltns35/go-vat/countries/latvia"
	"github.com/ltns35/go-vat/countries/liechtenstein"
	"github.com/ltns35/go-vat/countries/lithuania"
	"github.com/ltns35/go-vat/countries/luxembourg"
	"github.com/ltns35/go-vat/countries/malta"
	"github.com/ltns35/go-vat/countries/mocks"
	"github.com/ltns35/go-vat/countries/netherlands"
	"github.com/ltns35/go-vat/countries/north_macedonia"
	"github.com/ltns35/go-vat/countries/norway"
	"github.com/ltns35/go-vat/countries/peru"
	"github.com/ltns35/go-vat/countries/poland"
	"github.com/ltns35/go-vat/countries/portugal"
	"github.com/ltns35/go-vat/countries/romania"
	"github.com/ltns35/go-vat/countries/russia"
	"github.com/ltns35/go-vat/countries/san_marino"
	"github.com/ltns35/go-vat/countries/serbia"
	"github.com/ltns35/go-vat/countries/slovakia"
	"github.com/ltns35/go-vat/countries/slovenia"
	"github.com/ltns35/go-vat/countries/spain"
	"github.com/ltns35/go-vat/countries/sweden"
	"github.com/ltns35/go-vat/countries/switzerland"
	"github.com/ltns35/go-vat/countries/turkey"
	"github.com/ltns35/go-vat/countries/ukraine"
	"github.com/ltns35/go-vat/countries/united_kingdom"
)

func TestCheckVAT(t *testing.T) {

	type args struct {
		values []string
		want   bool
	}

	tests := []struct {
		name    string
		args    []args
		country countries.Calculer
	}{
		{
			name: "Albania",
			args: []args{
				{
					values: mocks.AlbaniaValidTests,
					want:   true,
				},
				{
					values: mocks.AlbaniaInvalidTests,
					want:   false,
				},
			},
			country: &albania.VAT,
		},
		{
			name: "Andorra",
			args: []args{
				{
					values: mocks.AndorraValidTests,
					want:   true,
				},
				{
					values: mocks.AndorraInvalidTests,
					want:   false,
				},
			},
			country: &andorra.VAT,
		},
		{
			name: "Argentina",
			args: []args{
				{
					values: mocks.ArgentinaValidTests,
					want:   true,
				},
				{
					values: mocks.ArgentinaInvalidTests,
					want:   false,
				},
			},
			country: &argentina.VAT,
		},
		{
			name: "Austria",
			args: []args{
				{
					values: mocks.AustriaValidTests,
					want:   true,
				},
				{
					values: mocks.AustriaInvalidTests,
					want:   false,
				},
			},
			country: &austria.VAT,
		},
		{
			name: "Belarus",
			args: []args{
				{
					values: mocks.BelarusValidTests,
					want:   true,
				},
				{
					values: mocks.BelarusInvalidTests,
					want:   false,
				},
			},
			country: &belarus.VAT,
		},
		{
			name: "Belgium",
			args: []args{
				{
					values: mocks.BelgiumValidTests,
					want:   true,
				},
				{
					values: mocks.BelgiumInvalidTests,
					want:   false,
				},
			},
			country: &belgium.VAT,
		},
		{
			name: "Brazil",
			args: []args{
				{
					values: mocks.BrazilValidTests,
					want:   true,
				},
				{
					values: mocks.BrazilInvalidTests,
					want:   false,
				},
			},
			country: &brazil.VAT,
		},
		{
			name: "Bulgaria",
			args: []args{
				{
					values: mocks.BulgariaValidTests,
					want:   true,
				},
				{
					values: mocks.BulgariaInvalidTests,
					want:   false,
				},
			},
			country: &bulgaria.VAT,
		},
		{
			name: "Croatia",
			args: []args{
				{
					values: mocks.CroatiaValidTests,
					want:   true,
				},
				{
					values: mocks.CroatiaInvalidTests,
					want:   false,
				},
			},
			country: &croatia.VAT,
		},
		{
			name: "Cyprus",
			args: []args{
				{
					values: mocks.CyprusValidTests,
					want:   true,
				},
				{
					values: mocks.CyprusInvalidTests,
					want:   false,
				},
			},
			country: &cyprus.VAT,
		},
		{
			name: "Czech Republic",
			args: []args{
				{
					values: mocks.CzechRepublicValidTests,
					want:   true,
				},
				{
					values: mocks.CzechRepublicInvalidTests,
					want:   false,
				},
			},
			country: &czech_republic.VAT,
		},
		{
			name: "Denmark",
			args: []args{
				{
					values: mocks.DenmarkValidTests,
					want:   true,
				},
				{
					values: mocks.DenmarkInvalidTests,
					want:   false,
				},
			},
			country: &denmark.VAT,
		},
		{
			name: "Estonia",
			args: []args{
				{
					values: mocks.EstoniaValidTests,
					want:   true,
				},
				{
					values: mocks.EstoniaInvalidTests,
					want:   false,
				},
			},
			country: &estonia.VAT,
		},
		{
			name: "Finland",
			args: []args{
				{
					values: mocks.FinlandValidTests,
					want:   true,
				},
				{
					values: mocks.FinlandInvalidTests,
					want:   false,
				},
			},
			country: &finland.VAT,
		},
		{
			name: "France",
			args: []args{
				{
					values: mocks.FranceValidTests,
					want:   true,
				},
				{
					values: mocks.FranceInvalidTests,
					want:   false,
				},
			},
			country: &france.VAT,
		},
		{
			name: "Germany",
			args: []args{
				{
					values: mocks.GermanyValidTests,
					want:   true,
				},
				{
					values: mocks.GermanyInvalidTests,
					want:   false,
				},
			},
			country: &germany.VAT,
		},
		{
			name: "Greece",
			args: []args{
				{
					values: mocks.GreeceValidTests,
					want:   true,
				},
				{
					values: mocks.GreeceInvalidTests,
					want:   false,
				},
			},
			country: &greece.VAT,
		},
		{
			name: "Hungary",
			args: []args{
				{
					values: mocks.HungaryValidTests,
					want:   true,
				},
				{
					values: mocks.HungaryInvalidTests,
					want:   false,
				},
			},
			country: &hungary.VAT,
		},
		{
			name: "Italy",
			args: []args{
				{
					values: mocks.ItalyValidTests,
					want:   true,
				},
				{
					values: mocks.ItalyInvalidTests,
					want:   false,
				},
			},
			country: &italy.VAT,
		},
		{
			name: "Ireland",
			args: []args{
				{
					values: mocks.IrelandValidTests,
					want:   true,
				},
				{
					values: mocks.IrelandInvalidTests,
					want:   false,
				},
			},
			country: &ireland.VAT,
		},
		{
			name: "Latvia",
			args: []args{
				{
					values: mocks.LatviaValidTests,
					want:   true,
				},
				{
					values: mocks.LatviaInvalidTests,
					want:   false,
				},
			},
			country: &latvia.VAT,
		},
		{
			name: "Liechtenstein",
			args: []args{
				{
					values: mocks.LiechtensteinValidTests,
					want:   true,
				},
				{
					values: mocks.LiechtensteinInvalidTests,
					want:   false,
				},
			},
			country: &liechtenstein.VAT,
		},
		{
			name: "Lithuania",
			args: []args{
				{
					values: mocks.LithuaniaValidTests,
					want:   true,
				},
				{
					values: mocks.LithuaniaInvalidTests,
					want:   false,
				},
			},
			country: &lithuania.VAT,
		},
		{
			name: "Luxembourg",
			args: []args{
				{
					values: mocks.LuxembourgValidTests,
					want:   true,
				},
				{
					values: mocks.LuxembourgInvalidTests,
					want:   false,
				},
			},
			country: &luxembourg.VAT,
		},
		{
			name: "Malta",
			args: []args{
				{
					values: mocks.MaltaValidTests,
					want:   true,
				},
				{
					values: mocks.MaltaInvalidTests,
					want:   false,
				},
			},
			country: &malta.VAT,
		},
		{
			name: "Netherlands",
			args: []args{
				{
					values: mocks.NetherlandsValidTests,
					want:   true,
				},
				{
					values: mocks.NetherlandsInvalidTests,
					want:   false,
				},
			},
			country: &netherlands.VAT,
		},
		{
			name: "North Macedonia",
			args: []args{
				{
					values: mocks.NorthMacedoniaValidTests,
					want:   true,
				},
				{
					values: mocks.NorthMacedoniaInvalidTests,
					want:   false,
				},
			},
			country: &north_macedonia.VAT,
		},
		{
			name: "Norway",
			args: []args{
				{
					values: mocks.NorwayValidTests,
					want:   true,
				},
				{
					values: mocks.NorwayInvalidTests,
					want:   false,
				},
			},
			country: &norway.VAT,
		},
		{
			name: "Peru",
			args: []args{
				{
					values: mocks.PeruValidTests,
					want:   true,
				},
				{
					values: mocks.PeruInvalidTests,
					want:   false,
				},
			},
			country: &peru.VAT,
		},
		{
			name: "Poland",
			args: []args{
				{
					values: mocks.PolandValidTests,
					want:   true,
				},
				{
					values: mocks.PolandInvalidTests,
					want:   false,
				},
			},
			country: &poland.VAT,
		},
		{
			name: "Portugal",
			args: []args{
				{
					values: mocks.PortugalValidTests,
					want:   true,
				},
				{
					values: mocks.PortugalInvalidTests,
					want:   false,
				},
			},
			country: &portugal.VAT,
		},
		{
			name: "Romania",
			args: []args{
				{
					values: mocks.RomaniaValidTests,
					want:   true,
				},
				{
					values: mocks.RomaniaInvalidTests,
					want:   false,
				},
			},
			country: &romania.VAT,
		},
		{
			name: "Russia",
			args: []args{
				{
					values: mocks.RussiaValidTests,
					want:   true,
				},
				{
					values: mocks.RussiaInvalidTests,
					want:   false,
				},
			},
			country: &russia.VAT,
		},
		{
			name: "San Marino",
			args: []args{
				{
					values: mocks.SanMarinoValidTests,
					want:   true,
				},
				{
					values: mocks.SanMarinoInvalidTests,
					want:   false,
				},
			},
			country: &san_marino.VAT,
		},
		{
			name: "Serbia",
			args: []args{
				{
					values: mocks.SerbiaValidTests,
					want:   true,
				},
				{
					values: mocks.SerbiaInvalidTests,
					want:   false,
				},
			},
			country: &serbia.VAT,
		},
		{
			name: "Slovakia",
			args: []args{
				{
					values: mocks.SlovakiaValidTests,
					want:   true,
				},
				{
					values: mocks.SlovakiaInvalidTests,
					want:   false,
				},
			},
			country: &slovakia.VAT,
		},
		{
			name: "Slovenia",
			args: []args{
				{
					values: mocks.SloveniaValidTests,
					want:   true,
				},
				{
					values: mocks.SloveniaInvalidTests,
					want:   false,
				},
			},
			country: &slovenia.VAT,
		},
		{
			name: "Spain",
			args: []args{
				{
					values: mocks.SpainValidTests,
					want:   true,
				},
				{
					values: mocks.SpainInvalidTests,
					want:   false,
				},
			},
			country: &spain.VAT,
		},
		{
			name: "Sweden",
			args: []args{
				{
					values: mocks.SwedenValidTests,
					want:   true,
				},
				{
					values: mocks.SwedenInvalidTests,
					want:   false,
				},
			},
			country: &sweden.VAT,
		},
		{
			name: "Switzerland",
			args: []args{
				{
					values: mocks.SwitzerlandValidTests,
					want:   true,
				},
				{
					values: mocks.SwitzerlandInvalidTests,
					want:   false,
				},
			},
			country: &switzerland.VAT,
		},
		{
			name: "Turkey",
			args: []args{
				{
					values: mocks.TurkeyValidTests,
					want:   true,
				},
				{
					values: mocks.TurkeyInvalidTests,
					want:   false,
				},
			},
			country: &turkey.VAT,
		},
		{
			name: "Ukraine",
			args: []args{
				{
					values: mocks.UkraineValidTests,
					want:   true,
				},
				{
					values: mocks.UkraineInvalidTests,
					want:   false,
				},
			},
			country: &ukraine.VAT,
		},
		{
			name: "United Kingdom",
			args: []args{
				{
					values: mocks.UnitedKingdomValidTests,
					want:   true,
				},
				{
					values: mocks.UnitedKingdomInvalidTests,
					want:   false,
				},
			},
			country: &united_kingdom.VAT,
		},
	}
	for _, tt := range tests {
		for _, arg := range tt.args {

			count := 0

			for _, value := range arg.values {

				count++

				valid := "VALID"

				if !arg.want {
					valid = "INVALID"
				}

				name := fmt.Sprintf("%s %s nº %d: %s", tt.name, valid, count, value)

				t.Run(
					name, func(t *testing.T) {

						got, _ := CheckVAT(value, tt.country)
						if got.IsValid != arg.want {
							t.Errorf("CheckVAT() got = %v, want %v", got, arg.want)
						}
					},
				)

			}

		}

	}
}
