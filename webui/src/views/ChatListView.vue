<template>
  <div class="chat-list-wrapper">
    <header class="chat-header">
      <h1>Chats</h1>
      <button class="btn btn-profile" @click="goToProfile">Profile Settings</button>
    </header>

    <section class="chat-list">
      <ChatCard v-for="chat in conversations" :key="chat.name" :chat="chat" />
    </section>

    <button class="btn btn-add" @click="openCreatedDialog">+</button>

    <div v-if="showCreatedDialog" class="dialog-overlay">
      <div class="dialog-content">
        <h2>Create New</h2>
        <button class="btn dialog-action" @click="createConversation">Start Conversation</button>
        <button class="btn dialog-action" @click="createGroup">Create Group</button>
        <button class="btn dialog-cancel" @click="closeCreatedDialog">Cancel</button>
      </div>
    </div>
  </div>
</template>



<script>
import ChatCard from "@/components/ChatCard.vue";
import axios from "@/services/axios";

export default {
  components: {
    ChatCard,
  },
  data() {
    return {
      conversations: [],
      showCreatedDialog: false,
      updateInterval: null,
    };
  },
  methods: {
    async fetchConversations() {
      try {
        const response = await axios.get("/user-profile");
        const newConversations = response.data || [];

        // Only reloads if there is new data
        if (JSON.stringify(this.conversations) !== JSON.stringify(newConversations)) {
          this.conversations = newConversations;
        }
      } catch (error) {
        console.error("Error fetching conversations", error);
        this.conversations = [];
      }
    },
    async goToProfile() {
      this.$router.push("/profile")
      //this.closeCreatedDialog();
    },
    openCreatedDialog() {
      this.showCreatedDialog = true;
    },
    async createConversation() {
      this.$router.push("/search")
      //this.closeCreatedDialog();
    },
    async createGroup() {
      this.$router.push("/create-group")
      //this.closeCreatedDialog();
    },
    closeCreatedDialog() {
      this.showCreatedDialog = false;
    },
  },
  mounted() {
    this.fetchConversations();

    this.updateInterval = setInterval(() => {
      this.fetchConversations();
    }, 5000);
  },

  beforeUnmount() {
    // Delete the Interval
    if (this.updateInterval) {
      clearInterval(this.updateInterval);
    }
  },
};
</script>


<style scoped>
.chat-list-wrapper {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background-color: #f3f4f6; /* light gray */
  font-family: 'Inter', sans-serif;
  color: #1f2937; /* slate-800 */
  padding: 20px;
  position: relative;
}

.chat-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.chat-header h1 {
  font-size: 2rem;
  font-weight: 600;
}

.chat-list {
  background: #ffffff;
  border-radius: 12px;
  padding: 16px;
  flex-grow: 1;
  overflow-y: auto;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
}

.btn {
  cursor: pointer;
  border: none;
  font-size: 1rem;
  border-radius: 8px;
  transition: background-color 0.3s ease;
}

.btn-profile {
  background-color: #3b82f6; /* blue-500 */
  color: #ffffff;
  padding: 10px 16px;
}

.btn-profile:hover {
  background-color: #2563eb; /* blue-600 */
}

.btn-add {
  position: absolute;
  bottom: 24px;
  right: 24px;
  background-color: #10b981; /* emerald-500 */
  color: #fff;
  width: 64px;
  height: 64px;
  font-size: 2rem;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.btn-add:hover {
  background-color: #059669; /* emerald-600 */
}

.dialog-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(31, 41, 55, 0.6); /* slate-800/60% */
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.dialog-content {
  background: #ffffff;
  padding: 24px;
  border-radius: 12px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
  width: 300px;
  text-align: center;
}

.dialog-content h2 {
  margin-bottom: 16px;
  font-size: 1.5rem;
  font-weight: 500;
}

.dialog-action {
  background-color: #6366f1; /* indigo-500 */
  color: white;
  padding: 10px 20px;
  margin: 8px 0;
  width: 100%;
}

.dialog-action:hover {
  background-color: #4f46e5; /* indigo-600 */
}

.dialog-cancel {
  background-color: #ef4444; /* red-500 */
  color: white;
  padding: 10px 20px;
  width: 100%;
  margin-top: 12px;
}

.dialog-cancel:hover {
  background-color: #dc2626; /* red-600 */
}

</style>
