package pkg_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/crashdump/venlock/pkg"
	"github.com/crashdump/venlock/pkg/gomod"
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
			lg := pkg.NewVenlock[gomod.Library](tt.args.sourcePath, gomod.GoMod[gomod.Library]{})
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
		sourcePath string
		librarySet pkg.LibrarySet[gomod.Library]
	}
	tests := []struct {
		name    string
		args    args
		wantRes pkg.LibrarySet[gomod.Library]
		wantErr bool
	}{
		{
			name: "valid-no-gap",
			args: args{
				sourcePath: fixturesDir + "/gomod",
				librarySet: pkg.LibrarySet[gomod.Library]{
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
			},
			wantRes: pkg.LibrarySet[gomod.Library]{},
			wantErr: false,
		},
		{
			name: "valid-gap",
			args: args{
				sourcePath: fixturesDir + "/gomod",
				librarySet: pkg.LibrarySet[gomod.Library]{
					{
						Module: "github.com/PuerkitoBio/goquery",
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
			},
			wantRes: pkg.LibrarySet[gomod.Library]{
				{
					Module: "github.com/avelino/slugify",
				},
				{
					Module: "github.com/golang/protobuf",
				},
				{
					Module: "golang.org/x/net",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lg := pkg.NewVenlock[gomod.Library](tt.args.sourcePath, gomod.GoMod[gomod.Library]{})
			gotRes, err := lg.Enforce(tt.args.librarySet)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enforce() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantRes, gotRes, "Enforce() gotRes = %v, want %v", gotRes, tt.wantRes)
		})
	}
}
