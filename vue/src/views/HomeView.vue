<script setup lang="ts">
import { RouterLink } from "vue-router";
import { useStore } from "@/store/back";
import { onMounted } from "vue";

const { revokeBack } = useStore();

onMounted(() => {
  revokeBack();
});
</script>

<template>
  <RouterLink to="/start">Start</RouterLink>
</template>

<style scoped>
a {
  display: flex;
  justify-content: center;
  align-items: center;
  color: red;
  width: 125px;
  height: 75px;
  border: 1px solid hsla(160, 100%, 37%, 1);
  font-size: 1.5rem;
}
</style>
