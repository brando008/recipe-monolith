package main

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type RecipeDaoPg struct {
	conn *sql.DB
}

func NewRecipeDaoPg(conn *sql.DB) *RecipeDaoPg {
	return &RecipeDaoPg{conn: conn}
}

func (dao *RecipeDaoPg) GetAll() ([]*Recipe, error) {
	rows, err := dao.conn.Query(`
		SELECT r.id, r.title, r.created_at, i.id, i.name
		FROM recipes r
		LEFT JOIN ingredients i ON r.id = i.recipe_id
		ORDER BY r.created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}

	recipeMap := make(map[string]*Recipe)
	var recipes []*Recipe

	for rows.Next() {
		var rID, rTitle string
		var rCreatedAt time.Time
		var iID, iName sql.NullString

		if err := rows.Scan(&rID, &rTitle, &rCreatedAt, &iID, &iName); err != nil {
			return nil, err
		}

		recipe, exists := recipeMap[rID]
		if !exists {
			recipe = &Recipe{
				ID:          rID,
				Title:       rTitle,
				CreatedAt:   rCreatedAt,
				Ingredients: []Ingredient{},
			}
			recipeMap[rID] = recipe
			recipes = append(recipes, recipe)
		}

		if iID.Valid {
			recipe.Ingredients = append(recipe.Ingredients, Ingredient{
				ID:       iID.String,
				RecipeID: rID,
				Name:     iName.String,
			})
		}
	}
	return recipes, nil
}

func (dao *RecipeDaoPg) Create(recipe *Recipe) error {
	_, err := dao.conn.Exec(`INSERT INTO recipes (id, title, created_at) VALUES ($1, $2, $3)`, recipe.ID, recipe.Title, recipe.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (dao *RecipeDaoPg) AddIngredient(recipeID string, ingredient *Ingredient) error {
	_, err := dao.conn.Exec(`INSERT INTO ingredients (id, recipe_id, name) VALUES ($1, $2, $3)`, ingredient.ID, recipeID, ingredient.Name)
	if err != nil {
		return err
	}

	return nil
}

func (dao *RecipeDaoPg) DeleteRecipe(recipeID string) error {
	_, err := dao.conn.Exec(`DELETE FROM ingredients WHERE recipe_id = $1`, recipeID)
	if err != nil {
		return err
	}
	_, err = dao.conn.Exec(`DELETE FROM recipes WHERE id = $1`, recipeID)
	if err != nil {
		return err
	}

	return nil
}
func (dao *RecipeDaoPg) DeleteIngredient(ingredientID string) error {
	_, err := dao.conn.Exec(`DELETE FROM ingredients WHERE id = $1`, ingredientID)
	if err != nil {
		return err
	}

	return nil
}
