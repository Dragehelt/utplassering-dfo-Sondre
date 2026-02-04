<template>
  <div class="min-h-screen p-20">
    <h2 class="text-7xl text-center font-bold text-black m-10">Users</h2>

    <div class="text-black text-center text-4xl flex flex-col justify-center items-center gap-5">
      <input class="border-2 text-center w-[20vw] border-black rounded-3xl bg-white" v-model="userId" type="number" placeholder="Enter user ID" />
      <button class="border-black border-2 rounded-2xl text-3xl p-2 bg-slate-600 text-white font-bold" @click="goToUser">Go to User Profile</button>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { navigateTo } from "#app";

const userId = ref(null);

const goToUser = () => {
  if (userId.value) {
    navigateTo(`/users/${userId.value}`);
  }
};
</script>
