<template>
  <div class="user-card" @click="goToChat">
    <img 
      :src="user.profile_photo_url || defaultAvatar" 
      alt="Profile Picture" 
      class="profile-picture" 
    >
    <span class="username">{{ user.username }}</span>
  </div>
</template>

<script>
export default {
  props: {
    user: {
      type: Object,
      required: true
    }
  },
  computed: {
    defaultAvatar() {
      return `https://ui-avatars.com/api/?name=${encodeURIComponent(this.user.username)}&size=40`;
    }
  },
  methods: {
    goToChat() {
      this.$router.push({
        path: '/chat',
        query: { username: this.user.username }
      });
    }
  }
};
</script>

<style scoped>
.user-card {
  display: flex;
  align-items: center;
  gap: 12px;
  background: #d6c3ff;
  padding: 12px 16px;
  border-radius: 10px;
  font-size: 18px;
  font-weight: bold;
  cursor: pointer;
  transition: background 0.2s;
}

.user-card:hover {
  background: #c5aef0;
}

.profile-picture {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}

.username {
  flex-grow: 1;
  color: #1d192b;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.user-card.disabled {
  pointer-events: none;
  opacity: 0.5;
}
</style>
