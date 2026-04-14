<script setup>
import Keyboard from "./components/Keyboard.vue";
</script>

<template>
  <div class="keyboard-wrapper">
    <Keyboard />
  </div>
</template>

<style scoped>
.keyboard-wrapper {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background: #f2f3f5;
  width: 100wh;
  height: 100vh;
}
</style>
