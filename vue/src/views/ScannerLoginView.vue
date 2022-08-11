<template>
  <div class="login">
    <h1>Login to use scanner</h1>
    <form class="login-form">
      <div class="username-field">
        <label for="username">Username</label>
        <input id="username" type="text" v-model="login.username" />
      </div>
      <div>
        <label for="password">Password</label>
        <input id="password" type="password" v-model="login.password" />
      </div>
      <button @click.prevent="loginSubmit"></button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive } from "vue";
import { useHeaderStore } from "@/store/header";

const { grantBack, revokeScan } = useHeaderStore();

const login = reactive<{
  username: string;
  password: string;
}>({
  username: "",
  password: "",
});
function loginSubmit() {
  // login logic
}

onMounted(() => {
  grantBack();
  revokeScan();
});
</script>

<style scoped>
.login-form {
  display: flex;
  flex-direction: column;
}
</style>
