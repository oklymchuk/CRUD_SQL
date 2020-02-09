package main

import "testing"

func TestConfigData_InitConfig(t *testing.T) {
	type fields struct {
		DbAdmin      string
		DbPWD        string
		DbPath       string
		DbService    string
		DbHost       string
		DbPort       int
		Port         int
		LinesPerPage int
	}
	type args struct {
		pathToConfigFile string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &ConfigData{
				DbAdmin:      tt.fields.DbAdmin,
				DbPWD:        tt.fields.DbPWD,
				DbPath:       tt.fields.DbPath,
				DbService:    tt.fields.DbService,
				DbHost:       tt.fields.DbHost,
				DbPort:       tt.fields.DbPort,
				Port:         tt.fields.Port,
				LinesPerPage: tt.fields.LinesPerPage,
			}
			if got := conf.InitConfig(tt.args.pathToConfigFile); got != tt.want {
				t.Errorf("ConfigData.InitConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigData_ConnectString(t *testing.T) {
	type fields struct {
		DbAdmin      string
		DbPWD        string
		DbPath       string
		DbService    string
		DbHost       string
		DbPort       int
		Port         int
		LinesPerPage int
	}
	tests := []struct {
		name         string
		fields       fields
		wantDbdriver string
		wantConnstr  string
	}{
		//		{"first", args{"/home/ostap/CRUD_SQL/config.json"}, true, &ConfigData{"root", "root", "mysql", "MySQL", 8080, 20}}, // TODO: Add test cases.
		//		{"sec", args{"config.json"}, true, &ConfigData{"root", "root", "mysql", "MySQL", 8080, 20}},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &ConfigData{
				DbAdmin:      tt.fields.DbAdmin,
				DbPWD:        tt.fields.DbPWD,
				DbPath:       tt.fields.DbPath,
				DbService:    tt.fields.DbService,
				DbHost:       tt.fields.DbHost,
				DbPort:       tt.fields.DbPort,
				Port:         tt.fields.Port,
				LinesPerPage: tt.fields.LinesPerPage,
			}
			gotDbdriver, gotConnstr := conf.ConnectString()
			if gotDbdriver != tt.wantDbdriver {
				t.Errorf("ConfigData.ConnectString() gotDbdriver = %v, want %v", gotDbdriver, tt.wantDbdriver)
			}
			if gotConnstr != tt.wantConnstr {
				t.Errorf("ConfigData.ConnectString() gotConnstr = %v, want %v", gotConnstr, tt.wantConnstr)
			}
		})
	}
}
