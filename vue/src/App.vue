<template>
  <div class="wrapper">
    <AppNavHeader />
    <main>
      <RouterView />
    </main>
    <footer>
      <p>Give me money: <RouterLink to="/">Donate Now!</RouterLink></p>
    </footer>
  </div>
</template>

<script setup lang="ts">
import AppNavHeader from "./components/AppNavHeader/AppNavHeader.vue";
</script>

<style scoped>
.wrapper {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  height: 100vh;
  width: 100vw;
}

footer {
  display: flex;
  justify-content: center;
  width: 100%;
  align-items: center;
  background-color: #303030;
  height: 40px;
}
</style>
