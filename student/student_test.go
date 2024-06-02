package student

import (
	"testing"
)

func TestAddGetStudent(t *testing.T) {
    s := Student{ID: 1, Name: "Alice", Age: 20, Grade: "A"}
    AddStudent(s)
    
    got, err := GetStudent(1)
    if err != nil {
        t.Errorf("GetStudent(1) error: %v", err)
    }
    if got != s {
        t.Errorf("GetStudent(1) = %v, want %v", got, s)
    }
}

func TestUpdateDeleteStudent(t *testing.T) {
    s := Student{ID: 1, Name: "Alice", Age: 20, Grade: "A"}
    AddStudent(s)

    sUpdated := Student{ID: 1, Name: "Alice Smith", Age: 21, Grade: "A+"}
    if err := UpdateStudent(1, sUpdated); err != nil {
        t.Errorf("UpdateStudent(1) error: %v", err)
    }

    got, err := GetStudent(1)
    if err != nil {
        t.Errorf("GetStudent(1) error: %v", err)
    }
    if got != sUpdated {
        t.Errorf("GetStudent(1) = %v, want %v", got, sUpdated)
    }

    if err := DeleteStudent(1); err != nil {
        t.Errorf("DeleteStudent(1) error: %v", err)
    }

    if _, err := GetStudent(1); err == nil {
        t.Errorf("expected error for GetStudent(1) after deletion, got nil")
    }
}
