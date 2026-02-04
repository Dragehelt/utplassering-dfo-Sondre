<template>
  <div class="min-h-screen">
    <p class="text-7xl text-center p-50 font-bold text-black">
      {{ userDisplay }}
    </p>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from "vue-router";
import { computed } from "vue";

const route = useRoute();

const userDisplay = computed(() => {
  const id = Number(route.params.id);

  if (id >= 1 && id < 10000000) {
    return `User: ${id}`;
  } else {
    return "Incorrect ID, please use a number higher than 0 and lower than 10 million.";
  }
});
</script>
