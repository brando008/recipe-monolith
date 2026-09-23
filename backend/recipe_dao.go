package main

type RecipeDao interface {
	GetAll() ([]*Recipe, error)
	Create(recipe *Recipe) error
	AddIngredient(recipeID string, ingredient *Ingredient) error
	DeleteRecipe(recipeID string) error
	DeleteIngredient(ingredientID string) error
}
