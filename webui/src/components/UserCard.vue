<template>
  <div 
    class="user-card" 
    @click="passChoosenUser"
  >
    <img 
      :src="user.profile_photo_url || `https://ui-avatars.com/api/?name=${encodeURIComponent(user.username)}&size=40`" 
      alt="Profile Picture" 
      class="profile-picture" 
    >
    <span class="username">{{ user.username }}</span>
  </div>
</template>

<script>
export default {
  props: {
    user: Object,
  },
  methods: {
    passChoosenUser() {
        this.$router.push({ path: '/chat', query: { username: this.user.username } });
    },
  },
};
</script>

<style scoped>
.user-card {
  display: flex;
  align-items: center;
  gap: 16px;
  background: linear-gradient(135deg, #e0c3fc 0%, #8ec5fc 100%);
  padding: 18px 24px;
  border-radius: 16px;
  font-size: 18px;
  font-weight: 600;
  cursor: pointer;
  box-shadow: 0 4px 16px rgba(140, 94, 255, 0.08);
  transition: transform 0.12s, box-shadow 0.12s, background 0.2s;
}

.user-card:hover {
  transform: translateY(-2px) scale(1.02);
  box-shadow: 0 8px 24px rgba(140, 94, 255, 0.15);
  background: linear-gradient(135deg, #c3e0fc 0%, #d6c3ff 100%);
}

.profile-picture {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #fff;
  box-shadow: 0 2px 8px rgba(140, 94, 255, 0.08);
}

.username {
  color: #3d246c;
  letter-spacing: 0.02em;
}

.user-card.disabled {
  pointer-events: none;
  opacity: 0.5;
}
</style>
