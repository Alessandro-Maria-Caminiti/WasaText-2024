<template>
  <div class="search-container">
    <h1 class="title">Choose User</h1>
    <input
      v-model="searchQuery"
      type="text"
      placeholder="Search..."
      class="search-bar"
      @input="onSearch"
    />

    <div class="user-list">
      <template v-if="filteredUsers.length">
        <UserCard
          v-for="user in filteredUsers"
          :key="user.username"
          :user="user"
        />
      </template>
      <p v-else class="no-results">No users found.</p>
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
  background-color: #f3f4f6;
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 40px;
  font-family: 'Inter', sans-serif;
}

.title {
  font-size: 2rem;
  color: #111827;
  margin-bottom: 20px;
}

.search-bar {
  width: 80%;
  max-width: 400px;
  padding: 12px;
  font-size: 16px;
  border-radius: 8px;
  border: 1px solid #d1d5db;
  margin-bottom: 20px;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}

.user-list {
  width: 80%;
  max-width: 500px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.no-results {
  color: #6b7280;
  font-style: italic;
  margin-top: 20px;
}
</style>
