<script setup lang="ts">
import { RouterLink } from "vue-router";
import { useHeaderStore } from "@/store/header";
import { onMounted } from "vue";

const { revokeBack, grantScan } = useHeaderStore();

onMounted(() => {
  revokeBack();
  grantScan();
});
</script>

<template>
  <div class="btns">
    <RouterLink to="/start">
      <div class="big-btn start">Start</div>
    </RouterLink>
    <RouterLink to="/recheck">
      <div class="big-btn check">Check My Code</div>
    </RouterLink>
  </div>
</template>

<style scoped>
.big-btn {
  display: flex;
  justify-content: center;
  align-items: center;
  color: red;
  width: 170px;
  height: 75px;
  border: 1px solid hsla(160, 100%, 37%, 1);
}
.start {
  font-size: 1.5rem;
}
.check {
  font-size: 1.2rem;
}
a {
  margin-bottom: 10px;
  margin-top: 10px;
}

.btns {
  display: flex;
  position: relative;
  flex-direction: column;
  top: -50%;
}
</style>
