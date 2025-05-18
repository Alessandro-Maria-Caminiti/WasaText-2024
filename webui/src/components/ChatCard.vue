<template>
  <div class="chat-card" @click="goToConversation">
    <img 
      :src="chat.photo_url.String || `https://ui-avatars.com/api/?name=${encodeURIComponent(chat.name)}&size=40`" 
      alt="Profile Picture" 
      class="profile-picture" 
    >
    <div class="chat-info">
      <strong class="convo-name">{{ chat.name }}</strong>
      <p class="last-message">{{ chat.last_message.String }}</p>
    </div>
    <span class="chat-time">{{ formatTime(chat.last_message_time.Time) }}</span>
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
  background: #fff;
  border-radius: 16px;
  padding: 16px 20px;
  margin-bottom: 14px;
  box-shadow: 0 2px 12px 0 rgba(80, 80, 180, 0.08);
  cursor: pointer;
  transition: box-shadow 0.18s, transform 0.18s;
  gap: 16px;
}
.chat-card:hover {
  box-shadow: 0 6px 24px 0 rgba(80, 80, 180, 0.16);
  transform: translateY(-2px) scale(1.01);
}

.profile-picture {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #e0e0f0;
  background: #f3f3fa;
}

.chat-info {
  display: flex;
  flex-direction: column;
  justify-content: center;
  flex: 1 1 auto;
  min-width: 0;
  gap: 4px;
}

.convo-name {
  font-size: 1.1rem;
  font-weight: 600;
  color: #2d2250;
  margin-bottom: 2px;
  letter-spacing: 0.01em;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.last-message {
  font-size: 0.98rem;
  color: #6c5a99;
  opacity: 0.85;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  max-width: 100%;
}

.chat-time {
  font-size: 0.85rem;
  color: #b0a7d6;
  margin-left: 16px;
  align-self: flex-start;
  font-variant-numeric: tabular-nums;
}
</style>
