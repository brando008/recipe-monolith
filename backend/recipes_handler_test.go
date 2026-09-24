package main

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockRecipeDao struct {
	RecipeDao
	ForceError bool
}

func (m *MockRecipeDao) Create(recipe *Recipe) error {
	if m.ForceError {
		return errors.New("simulated database crash")
	}

	recipe.ID = "mock-id"
	return nil
}

func (m *MockRecipeDao) AddIngredient(recipeID string, ingredient *Ingredient) error {
	if m.ForceError {
		return errors.New("simulated database crash")
	}

	ingredient.ID = "mock-ingredient-id"
	ingredient.RecipeID = recipeID
	return nil
}

func TestCreateRecipe_DatabaseCrash(t *testing.T) {
	mockDao := &MockRecipeDao{ForceError: true}
	handler := NewRecipeHandler(mockDao)

	body := []byte(`{"title": "Valid Title", "ingredients": []}`)
	req := httptest.NewRequest(http.MethodPost, "/recipes", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.Create(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500 Internal Server Error on DB crash, got %d", w.Code)
	}
}

func TestCreateRecipe_Success(t *testing.T) {
	mockDao := &MockRecipeDao{ForceError: false}
	handler := NewRecipeHandler(mockDao)

	body := []byte(`{"title": "My First Recipe", "ingredients": []}`)
	req := httptest.NewRequest(http.MethodPost, "/recipes", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.Create(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %d", w.Code)
	}
}

func TestCreateRecipe_EmptyTitle(t *testing.T) {
	mockDao := &MockRecipeDao{}
	handler := NewRecipeHandler(mockDao)

	body := []byte(`{"title": "", "ingredients": []}`)
	req := httptest.NewRequest(http.MethodPost, "/recipes", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.Create(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 Bad Request for empty title, got %d", w.Code)
	}
}

func TestAddIngredient_MissingQueryParam(t *testing.T) {
	mockDao := &MockRecipeDao{}
	handler := NewRecipeHandler(mockDao)

	body := []byte(`{"name": "Salt"}`)
	req := httptest.NewRequest(http.MethodPost, "/ingredients", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.AddIngredient(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 Bad Request for missing recipe_id, got %d", w.Code)
	}
}
