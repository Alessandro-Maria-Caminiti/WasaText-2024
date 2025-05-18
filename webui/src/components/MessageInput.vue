<template>
  <div class="message-input">
    <input type="file" accept="image/*" @change="handleFileUpload">
    <button @click="sendImage">Send Image</button>

    <input v-model="message" placeholder="Type a message...">
    <button @click="sendMessage">Send</button>
  </div>
</template>

<script>
export default {
  emits: ["send", "send-image"],
  data() {
    return { message: '', imageFile: null };
  },
  methods: {
    sendMessage() {
      this.$emit('send', this.message);
      this.message = '';
    },

    handleFileUpload(event) {
      this.imageFile = event.target.files[0];
    },

    sendImage() {
      if (!this.imageFile) {
        alert("Please select an image first.");
        return;
      }

      // Emit the image file to ChatView
      this.$emit("send-image", this.imageFile);
    }
  }
};
</script>

<style scoped>
.message-input { display: flex; gap: 10px; }
</style>
