package upload

import "testing"

func TestUpload(t *testing.T) {
	type args struct {
		username       string
		password       string
		filePath       string
		repositoryName string
		nexusURL       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test upload",
			args: args{
				nexusURL:       "http://10.121.218.184:31696",
				username:       "admin",
				password:       "@2Pohcfz",
				repositoryName: "al-cloud",
				filePath:       "test.yaml",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Upload(tt.args.username, tt.args.password, tt.args.filePath, tt.args.repositoryName, tt.args.nexusURL); (err != nil) != tt.wantErr {
				t.Errorf("Upload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
