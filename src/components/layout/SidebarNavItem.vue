<script setup lang="ts">
const props = defineProps<{
  label: string
  path: string
}>()
</script>

<template>
  <div
    class="m-2 rounded-md text-white/80 transition-colors duration-200 hover:bg-blue-800 hover:text-white cursor-pointer"
  >
    <RouterLink :to="props.path">
      <h2 class="px-4 py-2">{{ props.label }}</h2>
    </RouterLink>
  </div>
</template>
