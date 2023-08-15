package pkg_test

import (
	"reflect"
	"testing"

	"github.com/crashdump/libguardian/pkg"
	"github.com/crashdump/libguardian/pkg/gomod"
)

func TestConfig_SaveAndLoad(t *testing.T) {
	type args struct {
		path string
	}
	type fields struct {
		Version   string
		Catalogue map[string]pkg.LibrarySet[gomod.Library]
	}
	tests := []struct {
		name        string
		args        args
		fields      fields
		wantConfig  pkg.Config[gomod.Library]
		wantLoadErr bool
		wantSaveErr bool
	}{
		{
			name: "valid",
			args: args{
				path: "/tmp/test.json",
			},
			fields: fields{
				Version: "1.0",
				Catalogue: map[string]pkg.LibrarySet[gomod.Library]{
					"gomod": {
						{
							Module: "foo",
						},
					},
				},
			},
			wantConfig: pkg.Config[gomod.Library]{
				Version: "1.0",
				Catalogue: map[string]pkg.LibrarySet[gomod.Library]{
					"gomod": {
						{
							Module: "foo",
						},
					},
				},
			},
			wantLoadErr: false,
			wantSaveErr: false,
		},
		{
			name: "invalid-version",
			args: args{
				path: "/tmp/test.json",
			},
			fields: fields{
				Version: "0.0",
				Catalogue: map[string]pkg.LibrarySet[gomod.Library]{
					"gomod": {
						{
							Module: "foo",
						},
					},
				}},
			wantLoadErr: true,
			wantSaveErr: false,
		},
		{
			name: "missing-inventory",
			args: args{
				path: "/tmp/test.json",
			},
			fields: fields{
				Version: "1.0",
			},
			wantLoadErr: true,
			wantSaveErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &pkg.Config[gomod.Library]{
				Version:   tt.fields.Version,
				Catalogue: tt.fields.Catalogue,
			}
			if err := c.Save(tt.args.path); (err != nil) != tt.wantSaveErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantSaveErr)
			}
			if err := c.Load(tt.args.path); (err != nil) != tt.wantLoadErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantLoadErr)
			}
			if !tt.wantLoadErr && !tt.wantSaveErr {
				if !reflect.DeepEqual(*c, tt.wantConfig) {
					t.Errorf("Enumerate() gotLibraries = %v, want %v", c, tt.wantConfig)
				}
			}

		})
	}
}
