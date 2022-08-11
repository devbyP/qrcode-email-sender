<template>
  <div class="recheck-page">
    <form>
      <div class="email-field">
        <label for="email">Email</label>
        <input id="email" type="email" v-model="state.email" />
      </div>
      <div class="orderId-field">
        <label for="orderId">Order ID</label>
        <input id="orderId" type="text" v-model="state.orderId" />
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive } from "vue";
import { useHeaderStore } from "@/store/header";

const { grantBack, grantScan } = useHeaderStore();

const state = reactive<{
  email: string;
  orderId: string;
}>({ email: "", orderId: "" });

onMounted(() => {
  grantBack();
  grantScan();
});
</script>

<style scoped>
form {
  display: flex;
  flex-direction: column;
}
</style>
