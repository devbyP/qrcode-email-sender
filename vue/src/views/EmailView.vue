<template>
  <div class="page-wrapper">
    <form class="user-form">
      <label for="fName">First Name</label>
      <input
        id="fName"
        type="text"
        name="firstName"
        :value="userInfo.firstName"
      />
      <label for="lName">Last Name</label>
      <input
        id="lName"
        type="text"
        name="lastName"
        :value="userInfo.lastName"
      />
      <label for="uEmail">Email</label>
      <input
        id="uEmail"
        type="email"
        name="user-email"
        :value="userInfo.email"
      />
    </form>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive } from "vue";
import { useStore } from "@/store/back";

interface IUserInfo {
  firstName: string;
  lastName: string;
  email: string;
}

const userInfo = reactive<IUserInfo>({
  firstName: "",
  lastName: "",
  email: "",
});
const { grantBack } = useStore();
onMounted(() => {
  grantBack();
});
</script>

<style scoped></style>
