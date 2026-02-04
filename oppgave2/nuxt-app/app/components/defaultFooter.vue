<template>
  <footer>
    <div
      class="bg-gray-300 flex h-20 items-center justify-center p-5 text-black bottom-0 left-0 w-full border-t-2 border-black"
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
