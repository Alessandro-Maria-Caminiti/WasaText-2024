<template>
  <div class="chat-list-container">
    <h1 class="chat-title">Chats</h1>
    <button class="profile-button" @click="goToProfile">Profile Settings</button>
    <div class="chat-list">
      <ChatCard v-for="chat in conversations" :key="chat.name" :chat="chat" />
    </div>
    <button class="add-chat-button" @click="openCreatedDialog">+</button>

    <!-- Dialog-Box -->
    <div v-if="showCreatedDialog" class="dialog-overlay">
      <div class="dialog-box">
        <h2>Create New</h2>
        <button class="dialog-button" @click="createConversation">Start New Conversation</button>
        <button class="dialog-button" @click="createGroup">Create New Group</button>
        <button class="dialog-button cancel" @click="closeCreatedDialog">Cancel</button>
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
  methods: {
    async fetchConversations() {
      try {
        const response = await axios.get("/user_profile");
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
};
</script>

<style scoped>
.chat-list-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #65558f 0%, #21005d 100%);
  font-family: 'Inter', 'Roboto', Arial, sans-serif;
  color: #fff;
  position: relative;
  padding: 40px 0 0 0;
}

.chat-title {
  font-size: 2.5rem;
  font-weight: 700;
  margin-bottom: 24px;
  letter-spacing: 1px;
  text-shadow: 0 2px 16px rgba(33,0,93,0.15);
}

.profile-button {
  position: absolute;
  top: 32px;
  right: 40px;
  background: rgba(255,255,255,0.12);
  color: #fff;
  padding: 12px 28px;
  border-radius: 24px;
  cursor: pointer;
  border: none;
  font-size: 1rem;
  font-weight: 500;
  box-shadow: 0 2px 8px rgba(33,0,93,0.08);
  transition: background 0.2s, box-shadow 0.2s;
}
.profile-button:hover {
  background: rgba(255,255,255,0.22);
  box-shadow: 0 4px 16px rgba(33,0,93,0.15);
}

.chat-list {
  background: rgba(255,255,255,0.95);
  border-radius: 18px;
  padding: 18px 0;
  width: 90%;
  max-width: 480px;
  min-width: 320px;
  height: 65vh;
  overflow-y: auto;
  box-shadow: 0 8px 32px rgba(33,0,93,0.18);
  margin-bottom: 32px;
  border: 1px solid rgba(101,85,143,0.08);
}

.add-chat-button {
  position: fixed;
  bottom: 48px;
  right: 48px;
  background: linear-gradient(135deg, #005047 60%, #00bfae 100%);
  color: #fff;
  font-size: 2.2rem;
  border-radius: 50%;
  width: 72px;
  height: 72px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  border: none;
  box-shadow: 0 6px 24px rgba(0,80,71,0.18);
  transition: background 0.2s, box-shadow 0.2s;
  z-index: 10;
}
.add-chat-button:hover {
  background: linear-gradient(135deg, #00bfae 60%, #005047 100%);
  box-shadow: 0 12px 32px rgba(0,80,71,0.22);
}

/* Dialog Overlay */
.dialog-overlay {
  position: fixed;
  top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(33,0,93,0.32);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 100;
}

/* Dialog Box */
.dialog-box {
  background: #fff;
  color: #21005d;
  padding: 32px 28px 24px 28px;
  border-radius: 18px;
  text-align: center;
  box-shadow: 0 8px 32px rgba(33,0,93,0.22);
  min-width: 320px;
  max-width: 90vw;
  animation: dialog-pop 0.18s cubic-bezier(.4,2,.6,1) both;
}
@keyframes dialog-pop {
  0% { transform: scale(0.85); opacity: 0; }
  100% { transform: scale(1); opacity: 1; }
}

.dialog-box h2 {
  margin-bottom: 24px;
  font-size: 1.6rem;
  font-weight: 700;
  color: #21005d;
}

/* Dialog Buttons */
.dialog-button {
  background: linear-gradient(135deg, #005047 60%, #00bfae 100%);
  color: #fff;
  border: none;
  padding: 12px 28px;
  margin: 12px 0;
  border-radius: 8px;
  cursor: pointer;
  font-size: 1.15rem;
  font-weight: 500;
  width: 100%;
  transition: background 0.18s, box-shadow 0.18s;
  box-shadow: 0 2px 8px rgba(0,80,71,0.10);
  display: block;
}
.dialog-button:not(.cancel):hover {
  background: linear-gradient(135deg, #00bfae 60%, #005047 100%);
}

.dialog-button.cancel {
  background: linear-gradient(135deg, #b22222 60%, #ff5252 100%);
  color: #fff;
  margin-top: 18px;
  font-weight: 600;
}
.dialog-button.cancel:hover {
  background: linear-gradient(135deg, #ff5252 60%, #b22222 100%);
}
</style>
