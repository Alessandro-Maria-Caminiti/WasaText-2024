<template>
  <div v-if="message" class="error-msg" role="alert">
    <p>{{ message }}</p>
  </div>
</template>

<script>
export default {
  props: {
    message: {
      type: String,
      required: true
    }
  }
};
</script>

<style scoped>
.error-msg {
  background-color: #ffe6e6;
  color: #b00020;
  padding: 12px 16px;
  border: 1px solid #b00020;
  border-radius: 6px;
  margin-top: 12px;
  font-weight: 500;
  text-align: center;
  transition: all 0.3s ease;
}
</style>
