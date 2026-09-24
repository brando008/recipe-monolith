<script setup lang="ts">
import { ref } from 'vue';
import { api, type Recipe } from '../api';

const props = defineProps<{ recipe: Recipe }>();
const emit = defineEmits(['refresh']); // Tells the parent to reload data

const newIngredient = ref('');

const handleAddIngredient = async () => {
  if (!newIngredient.value.trim() || !props.recipe.id) return;
  
  await api.addIngredient(props.recipe.id, newIngredient.value);
  newIngredient.value = ''; 
  emit('refresh'); 
};

const handleDeleteIngredient = async (ingredientId?: string) => {
  if (!ingredientId) return;
  await api.deleteIngredient(ingredientId);
  emit('refresh');
};

const handleDeleteRecipe = async () => {
  if (!props.recipe.id) return;
  await api.deleteRecipe(props.recipe.id);
  emit('refresh');
};
</script>

<template>
  <div class="card">
    <div class="header">
      <h2>{{ recipe.title }}</h2>
      <button @click="handleDeleteRecipe" class="danger">Delete Recipe</button>
    </div>

    <ul class="ingredient-list">
      <li v-for="ing in recipe.ingredients" :key="ing.id">
        {{ ing.name }}
        <button @click="handleDeleteIngredient(ing.id)" class="small-danger">x</button>
      </li>
    </ul>
    
    <p v-if="!recipe.ingredients || recipe.ingredients.length === 0" class="empty">
      No ingredients yet. Add one below!
    </p>

    <form @submit.prevent="handleAddIngredient" class="add-ingredient-form">
      <input 
        v-model="newIngredient" 
        placeholder="Type an ingredient and hit Enter" 
        type="text" 
      />
      <button type="submit">Add</button>
    </form>
  </div>
</template>

<style scoped>
.card {
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 1rem;
  margin-bottom: 1.5rem;
  background: #f9f9f9;
}
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}
.header h2 { margin: 0; }
.ingredient-list {
  list-style-type: none;
  padding: 0;
}
.ingredient-list li {
  display: flex;
  justify-content: space-between;
  padding: 0.5rem;
  background: white;
  border: 1px solid #eee;
  margin-bottom: 0.25rem;
  border-radius: 4px;
}
.add-ingredient-form {
  display: flex;
  gap: 0.5rem;
  margin-top: 1rem;
}
.add-ingredient-form input { flex-grow: 1; padding: 0.5rem; }
.danger { background: #ff4c4c; color: white; border: none; padding: 0.5rem; border-radius: 4px; cursor: pointer;}
.small-danger { background: transparent; color: #ff4c4c; border: none; cursor: pointer; font-weight: bold;}
.empty { font-style: italic; color: #666; font-size: 0.9rem;}
</style>