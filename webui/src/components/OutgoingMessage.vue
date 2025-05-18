<template>
  <div class="outgoing-message">
    <!-- Forwarded Label -->
    <div v-if="isForwarded" class="forwarded">Forwarded Message</div>

    <!-- Photo or Text Message -->
    <div v-if="isPhoto">
      <img :src="content" alt="Sent Photo" class="message-photo">
    </div>
    <span v-else class="content">{{ content }}</span>

    <!-- Timestamp & Status (✔ / ✔✔) -->
    <div class="message-status">
      <span class="timestamp">{{ formatTime(timestamp) }}</span>
      <span class="status">
        <template v-if="fullyRead">✔✔</template>
        <template v-else-if="fullyReceived">✔</template>
      </span>
    </div>

    <!-- Reaction Trigger Button -->
    <button class="reaction-button" @click="toggleReactionPopup">+</button>

    <!-- Display Reactions -->
    <div class="reactions">
      <div v-for="(reaction, index) in reactions" :key="index">
        <span>{{ reaction.reactor }}:</span> <span>{{ reaction.content }}</span>
      </div>
    </div>

    <!-- Emoji Popup -->
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
      required: false,
      default: false
    },
    reactions: {
      type: Array,
      required: false,
      default: () => []
    },
    fullyReceived: {
      type: Boolean,
      required: false,
      default: false
    },
    fullyRead: {
      type: Boolean,
      required: false,
      default: false
    },
    username: {
      type: String,
      required: false,
      default: ""
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
.outgoing-message {
  background: #d6c5f0;
  padding: 10px;
  border-radius: 10px;
  text-align: right;
  position: relative;
  margin-left: auto;
  max-width: 80%;
}

.message-photo {
  max-width: 100%;
  border-radius: 10px;
}

.forwarded {
  font-size: 12px;
  color: gray;
}

.content {
  font-size: 16px;
  display: inline-block;
  word-wrap: break-word;
  max-width: 100%;
}

.message-status {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 5px;
  font-size: 14px;
  color: #555;
  margin-top: 4px;
}

.status {
  font-weight: bold;
  color: green;
}

.reactions {
  margin-top: 5px;
  font-size: 14px;
  text-align: left;
}

.reaction-button {
  background-color: #005047;
  color: white;
  padding: 5px 10px;
  border-radius: 5px;
  cursor: pointer;
  border: none;
  margin-top: 5px;
}

.reaction-popup {
  position: absolute;
  bottom: 50px;
  right: 0;
  background: white;
  border: 1px solid #ccc;
  border-radius: 5px;
  padding: 10px;
  display: flex;
  gap: 10px;
  z-index: 10;
}

.reaction-popup button {
  background-color: transparent;
  border: none;
  font-size: 20px;
  cursor: pointer;
}

.reaction-popup button:hover {
  color: #005047;
}
</style>

