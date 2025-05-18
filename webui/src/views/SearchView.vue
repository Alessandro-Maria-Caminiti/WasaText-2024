<template>
  <div class="search-container">
    <h1 class="title">Choose User</h1>
    <input v-model="searchQuery" type="text" placeholder="Search..." class="search-bar" @input="fetchUsers" />
    <div class="user-list">
      <UserCard v-for="user in filteredUsers" :key="user.username" :user="user" />
    </div>
  </div>
</template>

<script>
import axios from "@/services/axios";
import UserCard from "@/components/UserCard.vue";

export default {
  components: {
    UserCard,
  },
  computed: {
    filteredUsers() {
      return this.users.filter(user => user.username !== sessionStorage.getItem("currentUser"));
    }
  },
  data() {
    return {
      searchQuery: "",
      users: [],
    };
  },
  methods: {
    async fetchUsers() {
      try {
        const response = await axios.get("/users", {
          params: { username: this.searchQuery },
        });
        this.users = response.data;
      } catch (error) {
        console.error("Error fetching users", error);
      }
    },
  },
  mounted() {
    this.fetchUsers();
  },
};
</script>

<style scoped>
.search-container {
  text-align: center;
  background: linear-gradient(135deg, #6a82fb 0%, #fc5c7d 100%);
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 40px;
  font-family: 'Inter', 'Segoe UI', Arial, sans-serif;
}

.title {
  font-size: 2.5rem;
  color: #fff;
  font-weight: 700;
  letter-spacing: 1px;
  margin-bottom: 24px;
  text-shadow: 0 2px 12px rgba(0,0,0,0.12);
}

.search-bar {
  width: 100%;
  max-width: 400px;
  padding: 14px 18px;
  font-size: 1.1rem;
  border-radius: 16px;
  border: none;
  margin: 0 0 24px 0;
  box-shadow: 0 2px 16px rgba(0,0,0,0.08);
  outline: none;
  transition: box-shadow 0.2s;
  background: #fff;
}

.search-bar:focus {
  box-shadow: 0 4px 24px rgba(108, 99, 255, 0.18);
}

.user-list {
  width: 100%;
  max-width: 500px;
  display: flex;
  flex-direction: column;
  gap: 18px;
  background: rgba(255,255,255,0.18);
  border-radius: 18px;
  padding: 24px 18px;
  box-shadow: 0 4px 32px rgba(0,0,0,0.10);
  backdrop-filter: blur(2px);
}
</style>
