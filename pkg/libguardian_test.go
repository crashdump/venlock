package pkg_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/crashdump/libguardian/pkg"
	"github.com/crashdump/libguardian/pkg/gomod"
)

func TestEnumerate(t *testing.T) {
	type args struct {
		sourcePath string
	}
	tests := []struct {
		name    string
		args    args
		wantRes pkg.LibrarySet[gomod.Library]
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				sourcePath: fixturesDir,
			},
			wantRes: pkg.LibrarySet[gomod.Library]{
				{
					Module: "github.com/PuerkitoBio/goquery",
				},
				{
					Module: "github.com/avelino/slugify",
				},
				{
					Module: "github.com/otiai10/copy",
				},
				{
					Module: "github.com/yuin/goldmark",
				},
				{
					Module: "golang.org/x/oauth2",
				},
				{
					Module: "github.com/andybalholm/cascadia",
				},
				{
					Module: "github.com/golang/protobuf",
				},
				{
					Module: "golang.org/x/net",
				},
				{
					Module: "golang.org/x/sys",
				},
				{
					Module: "golang.org/x/text",
				},
				{
					Module: "google.golang.org/appengine",
				},
				{
					Module: "google.golang.org/protobuf",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lg := pkg.NewLibGuardian[gomod.Library](tt.args.sourcePath, gomod.GoMod[gomod.Library]{})
			gotLibraries, err := lg.Enumerate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Enumerate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, gotLibraries, tt.wantRes,
				"Enumerate() gotLibraries = %v, want %v", gotLibraries, tt.wantRes)
		})
	}
}

func TestEnforce(t *testing.T) {
	type args struct {
		configPath string
		sourcePath string
	}
	tests := []struct {
		name    string
		args    args
		wantRes pkg.LibrarySet[gomod.Library]
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				configPath: fixturesDir + "/config.json",
				sourcePath: fixturesDir,
			},
			wantRes: make(pkg.LibrarySet[gomod.Library], 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lg := pkg.NewLibGuardian[gomod.Library](tt.args.sourcePath, gomod.GoMod[gomod.Library]{})
			gotRes, err := lg.Enforce()
			if (err != nil) != tt.wantErr {
				t.Errorf("Enforce() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, gotRes, tt.wantRes, "Enforce() gotRes = %v, want %v", gotRes, tt.wantRes)
		})
	}
}
