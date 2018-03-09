package Leet

import (
	"math/rand"
	"testing"
)

func TestTranslator_Translate(t *testing.T) {
	source = func() rand.Source { return rand.NewSource(0) }
	type fields struct {
		rep   Replacer
		level Level
		rand  *rand.Rand
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "default",
			fields: fields{
				level: max,
				rep:   defaultReplace(),
				rand:  rand.New(source()),
			},
			args: args{
				s: "This is a test",
			},
			want: "This is 4 t3st",
		},
		{
			name: "default",
			fields: fields{
				level: Lam3,
				rep:   defaultReplace(),
				rand:  rand.New(source()),
			},
			args: args{
				s: "This is a test",
			},
			want: "This is a test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trans := &Translator{
				rep:   tt.fields.rep,
				level: tt.fields.level,
				rand:  tt.fields.rand,
			}
			if got := trans.Translate(tt.args.s); got != tt.want {
				t.Errorf("Translator.Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
