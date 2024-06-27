package storage

import (
	"reflect"
	"tasks_app/pkg/model"
	"testing"
)

func TestStorage_NewTask(t *testing.T) {
	type args struct {
		task model.Task
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.NewTask(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.NewTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetTasks(t *testing.T) {
	tests := []struct {
		name    string
		s       *Storage
		want    []model.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetTasks()
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}
