package pkg_test

import (
	"reflect"
	"testing"

	"github.com/crashdump/libguardian/pkg"
)

var fixturesDir = "../test/fixtures"

func Test_findManifests(t *testing.T) {
	type args struct {
		filename string
		dir      string
	}
	tests := []struct {
		name    string
		args    args
		wantOut []string
		wantErr bool
	}{
		{
			name: "invalid-SourcePath",
			args: args{
				filename: "foo",
				dir:      "XXX",
			},
			wantErr: true,
		},
		{
			name: "valid-maven",
			args: args{
				filename: "pom.xml",
				dir:      fixturesDir,
			},
			wantOut: []string{
				"../test/fixtures/maven/pom.xml",
			},
			wantErr: false,
		},
		{
			name: "valid-gomod",
			args: args{
				filename: "go.mod",
				dir:      fixturesDir,
			},
			wantOut: []string{
				"../test/fixtures/gomod/go.mod",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := pkg.FindManifests(tt.args.filename, tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindManifests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("FindManifests() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
