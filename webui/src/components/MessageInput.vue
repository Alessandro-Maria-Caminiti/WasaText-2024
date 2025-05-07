<template>
  <div class="message-input">
    <!-- Image Upload -->
    <label class="file-upload">
      📷
      <input type="file" @change="handleFileUpload" accept="image/*" hidden />
    </label>

    <!-- Send Image Button -->
    <button @click="sendImage" class="send-btn">Send Image</button>

    <!-- Message Input -->
    <input
      v-model="message"
      placeholder="Type a message..."
      class="text-input"
      @keyup.enter="sendMessage"
    />

    <!-- Send Message Button -->
    <button @click="sendMessage" :disabled="!message.trim()" class="send-btn">Send</button>
  </div>
</template>

<script>
export default {
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
.message-input {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 10px;
  padding: 10px;
  background: #f5f5f5;
  border-top: 1px solid #ccc;
}

.file-upload {
  background-color: #6c5ce7;
  color: white;
  padding: 6px 12px;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
}

.text-input {
  flex: 1;
  padding: 8px 12px;
  border: 1px solid #ccc;
  border-radius: 5px;
  font-size: 16px;
  min-width: 150px;
}

.send-btn {
  background-color: #005047;
  color: white;
  padding: 6px 14px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.send-btn:disabled {
  background-color: #999;
  cursor: not-allowed;
}
</style>
