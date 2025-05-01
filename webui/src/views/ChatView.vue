<template>
  <div class="chat-view">
    <!-- Header with Title and Optional Group Settings -->
    <header class="chat-header">
      <h1 class="chat-title">{{ $route.query.username }}</h1>
      <button v-if="$route.query.isGroup" @click="goToGroupSettings" class="btn btn-group-settings">
        Group Settings
      </button>
    </header>

    <!-- Messages -->
    <main class="messages-container">
      <div v-for="msg in messages" :key="msg.message_id" class="message-wrapper">
        <IncomingMessage 
          v-if="msg.sender === $route.query.username" 
          :username="msg.sender"
          :content="msg.content"
          :timestamp="msg.timestamp"
          :is-photo="msg.is_photo"
          :is-forwarded="msg.is_forwarded"
          :reactions="msg.reactions"
          @reaction-added="handleReaction(msg.message_id, $event)"
        />
        <OutgoingMessage 
          v-else 
          :content="msg.content" 
          :timestamp="msg.timestamp"
          :is-photo="msg.is_photo"
          :is-forwarded="msg.is_forwarded"
          :reactions="msg.reactions"
          :fully-received="msg.fully_received"
          :fully-read="msg.fully_read"
          @reaction-added="handleReaction(msg.message_id, $event)"
        />
      </div>
    </main>

    <!-- Message Input -->
    <footer class="message-input-container">
      <MessageInput @send="handleSend" @send-image="handleImageUpload" />
    </footer>
  </div>
</template>


<script>
import IncomingMessage from '@/components/IncomingMessage.vue';
import OutgoingMessage from '@/components/OutgoingMessage.vue';
import MessageInput from '@/components/MessageInput.vue';
import axios from "@/services/axios";

export default {
  components: { IncomingMessage, OutgoingMessage, MessageInput },
  data() {
    return {
      messages: [],
      updateInterval: null,
    };
  },
  methods: {

    async fetchMessages() {
      try {
        const partnerUsername = this.$route.query.username;
        
        const response = await axios.get(`/conversations/${partnerUsername}`);
        this.messages = response.data;
      } catch (error) {
        console.error("Error fetching messages:", error);
      }
    },

    goToGroupSettings() {
      this.$router.push({
        path: "/group-settings",
        query: { groupname: this.$route.query.username }
      });
    },

    async handleSend(content) {
      try {
        const partnerUsername = this.$route.query.username;
        let isPhoto = false;

        // Überprüfen, ob der Inhalt eine URL zu einem Bild ist
        const urlMatch = content.match(/https?:\/\/.*\.(jpg|jpeg|png|gif|bmp|svg)/i);
        if (urlMatch) {
          isPhoto = true;
          content = urlMatch[0];
        }

        const newMessage = {
          message: content,
          is_photo: isPhoto, 
        };

        const response = await axios.post(`/conversations/${partnerUsername}`, newMessage);
        if (!this.messages) {
          this.messages = [];
        }
    
        this.messages.push(response.data);

        this.fetchMessages()
        
      } catch (error) {
        console.error("Error sending message:", error);
      }
    },

    async handleImageUpload(imageFile) {
      try {
        const formData = new FormData();
        formData.append("image", imageFile);

        // Upload an das Backend
        const uploadResponse = await axios.post("/upload", formData, {
          headers: { "Content-Type": "multipart/form-data" },
        });

        const imageUrl = uploadResponse.data.imageUrl;

        // Bild als Nachricht senden
        this.handleSend(imageUrl);
      } catch (error) {
        console.error("Error uploading image:", error);
      }
    },

    async handleReaction(messageId, reaction) {

      // Find message and add Reaktion
      const message = this.messages.find(msg => msg.message_id === messageId);

      if (!message) {
        console.error(`Message with ID ${messageId} not found!`);
        return;
      }

      // Sicherstellen, dass `reactions` existiert
      if (!message.reactions) {
        message.reactions = [];
      }

      // Reaktion hinzufügen
      message.reactions.push(reaction);

      // Send reaction to server
      try {

        await axios.put(`/conversations/messages/${messageId}/comment`, reaction);

        await this.fetchMessages();

      } catch (error) {

        console.error("Error sending reaction:", error);

      }
    },

  },

  mounted() {
    this.fetchMessages();

    this.updateInterval = setInterval(() => {
      this.fetchMessages();
    }, 10000);
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
.chat-view {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background: #f3f4f6; /* Light gray background */
  font-family: 'Inter', sans-serif;
  color: #1f2937;
}

.chat-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 24px;
  background: #ffffff;
  border-bottom: 1px solid #e5e7eb;
  position: sticky;
  top: 0;
  z-index: 10;
}

.chat-title {
  font-size: 1.5rem;
  font-weight: 600;
  margin: 0;
}

.btn {
  border: none;
  border-radius: 8px;
  padding: 8px 16px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.btn-group-settings {
  background-color: #3b82f6; /* blue-500 */
  color: white;
}

.btn-group-settings:hover {
  background-color: #2563eb; /* blue-600 */
}

.messages-container {
  flex-grow: 1;
  overflow-y: auto;
  padding: 16px 24px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  background-color: #e5e7eb; /* Light gray */
}

.message-wrapper {
  display: flex;
}

.message-input-container {
  padding: 12px 24px;
  background: #ffffff;
  border-top: 1px solid #e5e7eb;
}

</style>
