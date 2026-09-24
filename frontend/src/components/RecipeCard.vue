<script setup lang="ts">
import { ref } from 'vue';
import { api, type Recipe } from '../api';

const props = defineProps<{ recipe: Recipe }>();
const emit = defineEmits(['refresh']);

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
      <button @click="handleDeleteRecipe" class="danger">[ DELETE_RECIPE ]</button>
    </div>

    <ul class="ingredient-list">
      <li v-for="ing in recipe.ingredients" :key="ing.id">
        <span>- {{ ing.name }}</span>
        <button @click="handleDeleteIngredient(ing.id)" class="small-danger">[x]</button>
      </li>
    </ul>
    
    <p v-if="!recipe.ingredients || recipe.ingredients.length === 0" class="empty">
      // NO_INGREDIENTS_FOUND
    </p>

    <form @submit.prevent="handleAddIngredient" class="add-ingredient-form">
      <span class="prompt">~</span>
      <input 
        v-model="newIngredient" 
        placeholder="ADD_INGREDIENT..." 
        type="text" 
      />
      <button type="submit">[ ADD ]</button>
    </form>
  </div>
</template>

<style scoped>
.card {
  border: 1px dashed #33ff00;
  padding: 1rem;
  margin-bottom: 1.5rem;
  background: #000;
  color: #33ff00;
}
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
  border-bottom: 1px dashed #1a8000;
  padding-bottom: 0.5rem;
}
.header h2 { 
  margin: 0; 
  font-size: 1.2rem;
  text-transform: uppercase;
}
.ingredient-list {
  list-style-type: none;
  padding: 0;
}
.ingredient-list li {
  display: flex;
  justify-content: space-between;
  padding: 0.5rem 0;
}
.add-ingredient-form {
  display: flex;
  gap: 0.5rem;
  margin-top: 1rem;
  align-items: center;
}
.prompt {
  color: #33ff00;
}
.add-ingredient-form input { 
  flex-grow: 1; 
  padding: 0.5rem;
  background: transparent;
  border: none;
  border-bottom: 1px dotted #33ff00;
  color: #33ff00;
  font-family: 'Courier New', Courier, monospace;
  outline: none;
}
.add-ingredient-form input::placeholder {
  color: #1a8000;
}
.danger, .small-danger, .add-ingredient-form button {
  background: transparent;
  border: 1px solid transparent;
  color: #33ff00;
  font-family: 'Courier New', Courier, monospace;
  cursor: pointer;
}
.danger { border-color: #ff3333; color: #ff3333; }
.danger:hover { background: #ff3333; color: #000; }
.small-danger { color: #ff3333; font-weight: bold; }
.small-danger:hover { color: #fff; background: #ff3333; }
.add-ingredient-form button { border: 1px solid #33ff00; }
.add-ingredient-form button:hover { background: #33ff00; color: #000; }
.empty { font-style: italic; color: #1a8000; font-size: 0.9rem;}
</style>