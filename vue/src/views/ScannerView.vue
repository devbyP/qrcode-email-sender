<template>
  <div class="scanner-page">
    <!-- scanner -->
    <h1>Scanner Page</h1>
  </div>
</template>

<script setup lang="ts">
import { useHeaderStore } from "@/store/header";
import { onMounted } from "vue";

const { grantBack, revokeScan } = useHeaderStore();

onMounted(() => {
  grantBack();
  revokeScan();
});
</script>

<style scoped></style>
