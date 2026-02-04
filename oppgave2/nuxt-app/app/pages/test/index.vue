<template>
  <!-- <button @click="refresh()">Duck</button>
  <p class="text-black" v-if="data">{{ data }}</p>

  <div v-if="pending">Loading data</div> -->

  <p>{{ data }}</p>
</template>

<script setup lang="ts">
const { data } = await useFetch("/api/fetch");
console.log(data);
</script>
