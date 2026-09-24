<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { api, type Recipe } from './api';
import RecipeCard from './components/RecipeCard.vue';

const recipes = ref<Recipe[]>([]);
const newRecipeTitle = ref('');

const loadRecipes = async () => {
  recipes.value = await api.getRecipes();
};

const handleCreateRecipe = async () => {
  if (!newRecipeTitle.value.trim()) return;
  await api.createRecipe(newRecipeTitle.value);
  newRecipeTitle.value = '';
  await loadRecipes();
};

onMounted(() => {
  loadRecipes();
});
</script>

<template>
  <main class="container">
    <pre class="ascii-title">
  ____  _____ ____ ___ ____  _____   __  __  ____ ____ 
 |  _ \| ____/ ___|_ _|  _ \| ____| |  \/  |/ ___|  _ \
 | |_) |  _|| |    | || |_) |  _|   | |\/| | |  _| |_) |
 |  _ <| |__| |___ | ||  __/| |___  | |  | | |_| |  _ <
 |_| \_\_____\____|___|_|   |_____| |_|  |_|\____|_| \_\
    </pre>
    
    <form @submit.prevent="handleCreateRecipe" class="create-form">
      <span class="prompt">&gt;</span>
      <input 
        v-model="newRecipeTitle" 
        placeholder="ENTER_NEW_RECIPE_TITLE" 
        type="text" 
      />
      <button type="submit">[ CREATE ]</button>
    </form>

    <div class="recipe-list">
      <RecipeCard 
        v-for="recipe in recipes" 
        :key="recipe.id" 
        :recipe="recipe" 
        @refresh="loadRecipes"
      />
    </div>
  </main>
</template>

<style>
/* Global styles to turn the whole page into a terminal */
body {
  background-color: #0c0c0c;
  color: #33ff00;
  font-family: 'Courier New', Courier, monospace;
  margin: 0;
  padding: 0;
}
</style>

<style scoped>
.container {
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem;
}
.ascii-title {
  color: #33ff00;
  font-weight: bold;
  text-shadow: 0 0 5px rgba(51, 255, 0, 0.4);
  margin-bottom: 2rem;
  font-size: 0.8rem;
  overflow-x: hidden;
}
.create-form {
  display: flex;
  gap: 1rem;
  margin-bottom: 2rem;
  padding-bottom: 2rem;
  border-bottom: 1px dashed #33ff00;
  align-items: center;
}
.prompt {
  font-weight: bold;
}
.create-form input { 
  flex-grow: 1; 
  padding: 0.5rem; 
  font-size: 1rem;
  background: transparent;
  color: #33ff00;
  border: none;
  border-bottom: 1px solid #33ff00;
  font-family: 'Courier New', Courier, monospace;
  outline: none;
}
.create-form input::placeholder {
  color: #1a8000;
}
.create-form button { 
  padding: 0.5rem 1rem; 
  font-size: 1rem; 
  cursor: pointer; 
  background: transparent; 
  color: #33ff00; 
  border: 1px solid #33ff00; 
  font-family: 'Courier New', Courier, monospace;
  transition: all 0.2s;
}
.create-form button:hover {
  background: #33ff00;
  color: #0c0c0c;
}
</style>