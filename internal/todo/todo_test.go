package todo_test

import (
	"My-first-api/internal/db"
	"My-first-api/internal/todo"
	"context"
	"reflect"
	"testing"
)

type MockDB struct {
	items []db.Item
}

func (m MockDB) InsertItem(ctx context.Context, item db.Item) error {
	//TODO implement me
	m.items = append(m.items, item)
	return nil

}

func (m MockDB) GetAllItems(ctx context.Context) ([]db.Item, error) {
	//TODO implement me
	return m.items, nil

}

func TestService_Search(t *testing.T) {

	tests := []struct {
		name           string
		toDosToAdd     []string
		query          string
		expectedResult []string
	}{
		// TODO: Add test cases.
		{
			name:           "Given a todo of shop and a search of sh, i should get shop back",
			toDosToAdd:     []string{"shop"},
			query:          "sh",
			expectedResult: []string{"shop"},
		},
		{
			name:           "Still returns shop, even if the case doesnt match",
			toDosToAdd:     []string{"Shopping"},
			query:          "sh",
			expectedResult: []string{"Shopping"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDB{}
			svc := todo.NewService(m)
			for _, toAdd := range tt.toDosToAdd {
				err := svc.Add(toAdd)
				if err != nil {
					t.Error(err)
				}
			}
			got, err := svc.Search(tt.query)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("Search() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}
