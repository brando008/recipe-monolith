package main

import (
	"encoding/json"
	"net/http"
)

type RecipeHandler struct {
	dao RecipeDao
}

func NewRecipeHandler(dao RecipeDao) *RecipeHandler {
	return &RecipeHandler{dao: dao}
}

func (h *RecipeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.dao.GetAll()
	if err != nil {
		http.Error(w, "Failed to get recipes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func (h *RecipeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var recipe Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if recipe.Title == "" {
		http.Error(w, "Recipe title cannot be empty", http.StatusBadRequest)
		return
	}

	if err := h.dao.Create(&recipe); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recipe)
}

func (h *RecipeHandler) AddIngredient(w http.ResponseWriter, r *http.Request) {
	var ingredient Ingredient
	if err := json.NewDecoder(r.Body).Decode(&ingredient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Missing recipe_id parameter", http.StatusBadRequest)
		return
	}

	if err := h.dao.AddIngredient(id, &ingredient); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ingredient)
}

func (h *RecipeHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.dao.DeleteRecipe(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *RecipeHandler) DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.dao.DeleteIngredient(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
