<template>
  
  <div class="min-h-[80vh] flex justify-center items-center w-full flex-col">
    <p class="text-7xl text-center font-bold text-black pb-5">
      {{ userDisplay }}
    </p>
    <div
      class="bg-slate-600 border-2 flex border-black rounded-3xl max-w-[50vw] max-h-[50vh] p-3 justify-center items-center"
    >
      <img class="max-w-[40vw] max-h-[40vh] rounded-2xl" :src="data" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from "vue-router";
import { computed } from "vue";

const route = useRoute();

const userDisplay = computed(() => {
  const id = Number(route.params.id);

  if (id >= 1 && id < 1000) {
    return `Bruker ${id}`;
  } else {
    return "Incorrect ID, please use a number higher than 0 and lower than 10 million.";
  }
});

const { data } = await useFetch("/api/fetch");
if (!data) {
  throw new Error("no data");
}
</script>
