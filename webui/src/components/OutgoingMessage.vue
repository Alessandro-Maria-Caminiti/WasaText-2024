<template>
  <div class="outgoing-message">
    <!-- Forwarded -->
    <div v-if="isForwarded" class="forwarded">Forwarded Message</div>
    <!-- Photo display -->
    <div v-if="isPhoto">
      <img :src="content" alt="Sent Photo" class="message-photo">
    </div>
    <span v-else class="content">{{ content }}</span>
    <!-- Message Status and Timestamp -->
    <div class="message-status">
      <span class="timestamp">{{ formatTime(timestamp) }}</span>
      <span v-if="fullyRead" class="status">✔✔</span>
      <span v-else-if="fullyReceived" class="status">✔</span>
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
      default: false
    },
    isForwarded: {
      type: Boolean,
      default: false
    },
    reactions: {
      type: Array,
      default: () => []
    },
    fullyReceived: {
      type: Boolean,
      default: false
    },
    fullyRead: {
      type: Boolean,
      default: false
    },
    username: {
      type: String,
      required: true
    }
  },
  emits: ["reaction-added"],
  data(){
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
  font-size: 12px;
  color: #8a8a8a;
  margin-bottom: 6px;
  font-style: italic;
  letter-spacing: 0.02em;
}

.content {
  display: inline-block;
  font-size: 16px;
  color: #232323;
  margin-bottom: 6px;
  word-break: break-word;
}

.message-status {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #7a7a7a;
  margin-top: 6px;
  justify-content: flex-end;
}

.status {
  font-weight: 600;
  color: #4caf50;
  font-size: 15px;
}

.reactions {
  margin-top: 8px;
  font-size: 15px;
  display: flex;
  gap: 10px;
  flex-wrap: wrap;
}

.reactions > div {
  background: #f3f3fa;
  border-radius: 12px;
  padding: 2px 10px;
  display: flex;
  align-items: center;
  box-shadow: 0 1px 3px rgba(80, 80, 120, 0.04);
}

.reaction-button {
  background: linear-gradient(135deg, #8ec5fc 0%, #e0c3fc 100%);
  color: #232323;
  padding: 6px 14px;
  border-radius: 50%;
  cursor: pointer;
  border: none;
  font-size: 20px;
  position: absolute;
  bottom: 10px;
  left: -36px;
  box-shadow: 0 2px 8px rgba(80, 80, 120, 0.10);
  transition: background 0.2s;
}

.reaction-button:hover {
  background: linear-gradient(135deg, #e0c3fc 0%, #8ec5fc 100%);
}

.reaction-popup {
  position: absolute;
  bottom: 60px;
  right: 0;
  background: #fff;
  border: 1px solid #e0e0e0;
  border-radius: 14px;
  padding: 12px 16px;
  display: flex;
  gap: 16px;
  box-shadow: 0 6px 24px rgba(80, 80, 120, 0.12);
  z-index: 10;
}

.reaction-popup button {
  background-color: transparent;
  border: none;
  font-size: 24px;
  cursor: pointer;
  transition: color 0.15s, transform 0.15s;
  padding: 4px 8px;
}

.reaction-popup button:hover {
  color: #4f46e5;
  transform: scale(1.2);
}
</style>
