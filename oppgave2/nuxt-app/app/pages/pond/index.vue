<template>
  <!-- pic div -->
  <div
    v-if="visible === true"
    class="absolute flex justify-center items-center w-full pt-30"
  >
    <div
      class="bg-slate-600 border-2 flex border-black rounded-3xl max-w-[50vw] max-h-[50vh] p-3 justify-center items-center"
    >
      <img class="max-w-[40vw] max-h-[40vh] rounded-2xl" :src="data" />
    </div>
  </div>
  <div v-else class="text-black"></div>

  <div>
    <!-- html div -->
    <div
      class="w-full flex justify-center items-center flex-col pt-10 pb-30 max-h-[50vh]"
    >
      <h2 class="text-black text-5xl font-bold">The Duckpond</h2>
    </div>
    <!-- Duckpond div -->
    <div class="w-full flex justify-center items-center pb-80">
      <img
        class="w-[70%] -z-10 absolute"
        src="assets/img/Duckpond.png"
        alt=""
      />
    </div>
    <!-- button div -->
    <div class="w-full flex justify-center items-center flex-col pb-10">
      <button
        v-if="!handsFull"
        class="border-black border-2 rounded-2xl text-3xl p-2 bg-slate-600 text-white font-bold"
        @click="
          // refreshData();
          duckPickedUp();
          visible = !visible;
        "
      >
        Pick up duck(s)
      </button>
      <button
        v-else
        class="border-black border-2 rounded-2xl text-3xl p-2 bg-slate-600 text-white font-bold"
        @click="
          refreshData();
          duckPickedUp();
          visible = !visible;
        "
      >
        Release duck(s)
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
const visible = ref(false);

let handsFull = false;

function duckPickedUp() {
  handsFull = !handsFull;
}

const refreshData = async () => {
  await refresh();
};
const { data, refresh, pending } = await useFetch("/api/fetch");
if (!data) {
  throw new Error("no data");
}
</script>
