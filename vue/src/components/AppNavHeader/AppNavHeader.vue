<template>
  <header>
    <h1>Get Email with QRcode!!</h1>
    <div v-show="hasBack" @click="goback" class="back">&lt; Go Back</div>
    <div v-show="toScan" @click="toScanPage" class="scanBtn">Scan Code</div>
  </header>
</template>

<script setup lang="ts">
import { storeToRefs } from "pinia";
import { useHeaderStore } from "@/store/header";
import { useRouter } from "vue-router";

const store = useHeaderStore();
const router = useRouter();
const { hasBack, toScan } = storeToRefs(store);
function goback() {
  router.push("/");
}
function toScanPage() {
  router.push("/scanner");
}
</script>

<style scoped>
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
.scanBtn {
  position: absolute;
  right: 2rem;
  transition: 0.4s;
  background-color: transparent;
  cursor: pointer;
}
.scanBtn:hover {
  background-color: hsla(160, 100%, 37%, 1);
}
</style>
