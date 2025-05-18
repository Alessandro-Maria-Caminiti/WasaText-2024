<template>
  <div 
    class="incoming-message"
  >
    <!-- Sender Username -->
    <span class="username" aria-label="Sender username">{{ username }}</span>
    <!-- Forwarded Label -->
    <div v-if="isForwarded" class="forwarded">Forwarded Message</div>
    <!-- Message Content -->
    <div class="message-content">
      <img
        v-if="isPhoto"
        :src="content"
        alt="Received Photo"
        class="message-photo"
      >
      <span v-else class="content">{{ content }}</span>
    </div>
    <!-- Timestamp -->
    <div class="message-info">
      <span class="timestamp">{{ formatTime(timestamp) }}</span>
    </div>
    <!-- Reaction Button -->
    <button aria-label="React to message" class="reaction-button" @click="toggleReactionPopup">+</button>
    <!-- Reactions Display -->
    <div v-if="reactions.length" class="reactions">
      <div
        v-for="(reaction, index) in reactions"
        :key="index"
        class="reaction"
      >
        <span class="reactor">{{ reaction.reactor }}:</span>
        <span class="emoji">{{ reaction.content }}</span>   
      </div>
    </div>
    <!-- Emoji Selection Popup -->
    <div v-if="isReacting" class="reaction-popup">
      <button @click="addReaction(':D')">:D</button>
      <button @click="addReaction('D:')">D:</button>
      <button @click="addReaction(':|')">:|</button>
      <button @click="addReaction(':O')">:O</button>
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
      type: [String, Number, Date],
      required: true
    },
    isPhoto: {
      type: Boolean,
      required: true
    },
    isForwarded: {
      type: Boolean,
      required: true
    },
    reactions: {
      type: Array,
      required: true
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
  background: #e6dfff;
  padding: 12px;
  border-radius: 12px;
  text-align: left;
  position: relative;
  margin-bottom: 12px;
  font-family: 'Arial', sans-serif;
}

.username {
  font-weight: bold;
  color: #4b0082;
  margin-bottom: 4px;
  display: block;
}

.forwarded {
  font-size: 12px;
  color: gray;
  margin-bottom: 6px;
}

.message-content {
  margin-bottom: 6px;
}

.message-photo {
  max-width: 100%;
  border-radius: 10px;
}

.content {
  font-size: 16px;
  color: #333;
}

.message-info {
  font-size: 13px;
  color: #555;
  margin-bottom: 6px;
}

.reaction-button {
  background-color: #005047;
  color: white;
  padding: 4px 10px;
  border-radius: 5px;
  cursor: pointer;
  border: none;
  font-size: 16px;
}

.reactions {
  margin-top: 8px;
  font-size: 14px;
}

.reaction {
  display: flex;
  gap: 4px;
}

.reactor {
  font-weight: bold;
  color: #333;
}

.reaction-popup {
  position: absolute;
  bottom: 40px;
  left: 10px;
  background: white;
  border: 1px solid #ccc;
  border-radius: 6px;
  padding: 8px;
  display: flex;
  gap: 8px;
  z-index: 10;
  box-shadow: 0 2px 8px rgba(0,0,0,0.2);
}

.reaction-popup button {
  background-color: transparent;
  border: none;
  font-size: 20px;
  cursor: pointer;
  transition: transform 0.2s;
}

.reaction-popup button:hover {
  color: #005047;
  transform: scale(1.2);
}
</style>
