<template>
  <div class="chat-view">
    <!-- Partner Username as Title -->
    <h1 class="chat-title">{{ $route.query.username }}</h1>

    <!-- Group Settings Button (nur wenn es eine Gruppe ist) -->
    <div v-if="$route.query.isGroup">
      <button class="group-settings-button" @click="goToGroupSettings">
        Group Settings
      </button>
    </div>

    <!-- Show the messages List -->
    <div v-for="msg in messages" :key="msg.message_id">
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

    <!-- New Message Input -->
    <MessageInput @send="handleSend" @send-image="handleImageUpload" />
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

        // Upload Backend
        const uploadResponse = await axios.post("/upload", formData, {
          headers: { "Content-Type": "multipart/form-data" },
        });

        const imageUrl = uploadResponse.data.imageUrl;

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

      if (!message.reactions) {
        message.reactions = [];
      }

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
};
</script>
<style scoped>
.chat-view {
  background: linear-gradient(135deg, #6a82fb 0%, #fc5c7d 100%);
  padding: 32px 0 0 0;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.chat-title {
  font-size: 2.2rem;
  font-weight: 700;
  color: #fff;
  margin-bottom: 24px;
  letter-spacing: 1px;
  text-shadow: 0 2px 8px rgba(0,0,0,0.15);
}

.group-settings-button {
  position: absolute;
  top: 32px;
  right: 40px;
  background: linear-gradient(90deg, #43e97b 0%, #38f9d7 100%);
  color: #fff;
  padding: 12px 22px;
  border: none;
  border-radius: 24px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 600;
  box-shadow: 0 4px 16px rgba(67,233,123,0.15);
  transition: background 0.2s, transform 0.15s;
}

.group-settings-button:hover {
  background: linear-gradient(90deg, #38f9d7 0%, #43e97b 100%);
  transform: translateY(-2px) scale(1.04);
  box-shadow: 0 6px 20px rgba(67,233,123,0.22);
}
</style>
