<template>
  <div class="wrapper">
    <header>
      <h1>Get Email and QRcode!!</h1>
      <div v-show="hasBack === true" @click="goback" class="back">&lt; go back</div>
    </header>
    <main>
      <RouterView />
    </main>
    <footer>
      <p>Give me money: <RouterLink to="/">Donate Now!</RouterLink></p>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { RouterView } from "vue-router";
import { storeToRefs } from "pinia";
import { useStore } from "./store/back";
import { useRouter } from "vue-router";

const store = useStore();
const router = useRouter();
const { hasBack } = storeToRefs(store);
function goback() {
  router.push("/");
}
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

header {
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #303030;
  width: 100%;
  height: 75px;
}

.back {
  position: absolute;
  left: 2rem;
  transition: 0.4s;
  background-color: transparent;
  cursor: pointer;
}
.back:hover {
  background-color: hsla(160, 100%, 37%, 1);
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
