<template>
  <div class="chat-card" @click="goToConversation">
    <img 
      :src="chat.photo_url?.String || defaultAvatar" 
      alt="Profile Picture" 
      class="profile-picture" 
    >

    <div class="chat-details">
      <div class="chat-header">
        <span class="convo-name">{{ chat.name }}</span>
        <span class="chat-time">{{ formatTime(chat.last_message_time?.Time) }}</span>
      </div>
      <p class="last-message">{{ chat.last_message?.String || "No messages yet" }}</p>
    </div>
  </div>
</template>


<script>
export default {
  props: {
    chat: Object
  },
  methods: {
    formatTime(timestamp) {
      if (!timestamp) return "";
      return new Date(timestamp).toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });
    },
    goToConversation() {
      this.$router.push({ path: '/chat', query: { username: this.chat.name, isGroup: this.chat.is_group } });
    }
  }
};
</script>

<style scoped>
.chat-card {
  display: flex;
  align-items: center;
  background: #e8def8;
  border-radius: 10px;
  padding: 12px;
  margin-bottom: 10px;
  cursor: pointer;
  transition: background 0.2s ease;
}

.chat-card:hover {
  background: #d7c7ef;
}

.profile-picture {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 12px;
}

.chat-details {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.chat-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 4px;
}

.convo-name {
  font-size: 18px;
  font-weight: 600;
  color: #1d192b;
  max-width: 60%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.chat-time {
  font-size: 12px;
  color: #6c6c6c;
  white-space: nowrap;
}

.last-message {
  font-size: 14px;
  color: #333;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>

