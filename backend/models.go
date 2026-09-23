package main

import "time"

type Recipe struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	CreatedAt   time.Time    `json:"created_at"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	ID       string `json:"id"`
	RecipeID string `json:"recipe_id"`
	Name     string `json:"name"`
}
