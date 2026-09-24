const API_URL = 'http://localhost:8080';

export interface Ingredient {
  id?: string;
  recipe_id?: string;
  name: string;
}

export interface Recipe {
  id?: string;
  title: string;
  created_at?: string;
  ingredients?: Ingredient[];
}

export const api = {
  async getRecipes(): Promise<Recipe[]> {
    const res = await fetch(`${API_URL}/recipes`);
    if (!res.ok) throw new Error('Failed to fetch recipes');
    const data = await res.json();
    return data || []; // Go returns null for empty slices, fallback to []
  },

  async createRecipe(title: string): Promise<Recipe> {
    const res = await fetch(`${API_URL}/recipes`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ title }),
    });
    if (!res.ok) throw new Error('Failed to create recipe');
    return res.json();
  },

  async addIngredient(recipeId: string, name: string): Promise<void> {
    const res = await fetch(`${API_URL}/recipes/${recipeId}/ingredients/`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name }),
    });
    if (!res.ok) throw new Error('Failed to add ingredient');
  },

  async deleteRecipe(id: string): Promise<void> {
    const res = await fetch(`${API_URL}/recipes/${id}`, {
      method: 'DELETE',
    });
    if (!res.ok) throw new Error('Failed to delete recipe');
  },

  async deleteIngredient(id: string): Promise<void> {
    const res = await fetch(`${API_URL}/ingredients/${id}`, {
      method: 'DELETE',
    });
    if (!res.ok) throw new Error('Failed to delete ingredient');
  }
};