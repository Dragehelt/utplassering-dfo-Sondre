<template>
  <footer>
    <div
      class="bg-slate-600 flex h-20 items-center justify-center p-5 text-gray-300 bottom-0 left-0 w-full text-"
    >
      <div>
        <p class="100vw text-center">{{ username }}</p>
        <p class="100vw text-center">{{ year }}</p>
      </div>
    </div>
  </footer>
</template>

<script setup>
  const username = "Sondre";
  const year = new Date().getFullYear();
</script>
