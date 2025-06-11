<template>
  <div class="incoming-message">
    <!-- Sender Username -->
    <span class="username">{{ username }}</span>
    <!-- Forwarded -->
    <div v-if="isForwarded" class="forwarded">Forwarded Message</div>
    <!-- Photo display -->
    <div v-if="isPhoto">
      <img :src="content" alt="Received Photo" class="message-photo">
    </div>
    <span v-else class="content">{{ content }}</span>
    <!-- Timestamp -->
    <div class="message-info">
      <span class="timestamp">{{ formatTime(timestamp) }}</span>
    </div>
    <!-- Reaction Button -->
    <button class="reaction-button" @click="toggleReactionPopup">+</button>
    <!-- Reactions Display -->
    <div class="reactions">
      <div v-for="(reaction, index) in reactions" :key="index">
        <span>{{ reaction.reactor }}: </span><span>{{ reaction.content }}</span>
      </div>
    </div>
    <!-- Emoji Selection Popup -->
    <div v-if="isReacting" class="reaction-popup">
      <button @click="addReaction(':D')">:D</button>
      <button @click="addReaction('D:')">D:</button>
      <button @click="addReaction(':|')">:|</button>    
    </div>

  </div>
</template>

<script>
export default {
  props: {
  username: {
    type: String,
    required: true
  },
  content: {
    type: String,
    required: true
  },
  timestamp: {
    type: [String, Number, Date], // Depending on how it's passed
    required: true
  },
  isPhoto: {
    type: Boolean,
    default: false
  },
  isForwarded: {
    type: Boolean,
    default: false
  },
  reactions: {
    type: Array,
    default: () => []
  }
},

  emits: ["reaction-added"],
  data() {
    return {
      isReacting: false,  // Flag to toggle the emoji popup visibility
    };
  },
  methods: {
    toggleReactionPopup() {
      // Toggle the visibility of the emoji popup
      this.isReacting = !this.isReacting;
    },
    addReaction(emoji) {

      // Add the selected emoji to the reactions list along with the current username
      const newReaction = {
        content: emoji
      };

      // emit newReaction
      this.$emit("reaction-added", newReaction);

      this.isReacting = false; // close Popup

    },

    formatTime(timestamp) {
      if (!timestamp) return "";
      return new Date(timestamp).toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });
    },

  }
};
</script>

<style scoped>
.incoming-message {
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  padding: 18px 48px; /* Increased padding for a bigger box */
  border-radius: 20px;
  box-shadow: 0 2px 16px rgba(80, 80, 160, 0.10);
  text-align: left;
  position: relative;
  margin-bottom: 24px;
  transition: box-shadow 0.2s;
  min-width: 340px;   /* Minimum width for bigger messages */
  max-width: 820px;   /* Maximum width for large images */
}

.incoming-message:hover {
  box-shadow: 0 6px 32px rgba(80, 80, 160, 0.18);
}

.message-photo {
  max-width: 200%;
  max-height: 380px;   /* Allow taller images */
  border-radius: 16px;
  margin: 48px 0;
  box-shadow: 0 2px 12px rgba(80, 80, 160, 0.13);
  display: block;
  margin-left: auto;
  margin-right: auto;
}

.forwarded {
  font-size: 13px;
  color: #888;
  background: #f0f0f8;
  border-radius: 8px;
  padding: 2px 8px;
  display: inline-block;
  margin-bottom: 6px;
}

.message-info {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #888;
  margin-top: 6px;
}

.username {
  font-weight: 600;
  color: #5a189a;
  display: block;
  font-size: 16px;
  margin-bottom: 2px;
  letter-spacing: 0.02em;
}

.content {
  font-size: 15px;
  color: #22223b;
  margin: 6px 0;
  line-height: 1.6;
}

.reactions {
  margin-top: 8px;
  font-size: 15px;
  display: flex;
  gap: 10px;
  flex-wrap: wrap;
}

.reactions > div {
  background: #e0e7ff;
  border-radius: 16px;
  padding: 2px 10px;
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 15px;
}

.reaction-button {
  background: linear-gradient(135deg, #7b2ff2 0%, #f357a8 100%);
  color: #fff;
  padding: 6px 14px;
  border-radius: 50%;
  cursor: pointer;
  border: none;
  font-size: 20px;
  position: absolute;
  right: 16px;
  bottom: 16px;
  box-shadow: 0 2px 8px rgba(123, 47, 242, 0.10);
  transition: background 0.2s, box-shadow 0.2s;
}

.reaction-button:hover {
  background: linear-gradient(135deg, #f357a8 0%, #7b2ff2 100%);
  box-shadow: 0 4px 16px rgba(123, 47, 242, 0.18);
}

.reaction-popup {
  position: absolute;
  bottom: 56px;
  left: 16px;
  background: #fff;
  border: 1px solid #e0e0e0;
  border-radius: 12px;
  padding: 12px 16px;
  display: flex;
  gap: 16px;
  box-shadow: 0 4px 24px rgba(80, 80, 160, 0.12);
  z-index: 10;
}

.reaction-popup button {
  background-color: transparent;
  border: none;
  font-size: 22px;
  cursor: pointer;
  transition: color 0.2s, transform 0.2s;
  padding: 4px 8px;
  border-radius: 8px;
}

.reaction-popup button:hover {
  color: #7b2ff2;
  background: #f3e8ff;
  transform: scale(1.15);
}
.forward-button,
.delete-button {
  background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);
  color: #fff;
  padding: 6px 18px;
  border-radius: 16px;
  cursor: pointer;
  border: none;
  font-size: 15px;
  margin-top: 10px;
  margin-right: 10px;
  box-shadow: 0 2px 8px rgba(67, 233, 123, 0.10);
  transition: background 0.2s, box-shadow 0.2s, transform 0.2s;
}

.forward-button:hover {
  background: linear-gradient(135deg, #38f9d7 0%, #43e97b 100%);
  box-shadow: 0 4px 16px rgba(67, 233, 123, 0.18);
  transform: translateY(-2px) scale(1.05);
}

.delete-button {
  background: linear-gradient(135deg, #ff5858 0%, #f09819 100%);
  box-shadow: 0 2px 8px rgba(255, 88, 88, 0.10);
}

.delete-button:hover {
  background: linear-gradient(135deg, #f09819 0%, #ff5858 100%);
  box-shadow: 0 4px 16px rgba(255, 88, 88, 0.18);
  transform: translateY(-2px) scale(1.05);
}
</style>
