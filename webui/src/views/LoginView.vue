<template>
  <div class="login-wrapper">
    <div class="login-card">
      <h1>Login</h1>
      <input
        v-model="username"
        class="login-input"
        type="text"
        placeholder="Enter your username"
      />
      <button class="login-button" @click="login">Login</button>
      <p v-if="error" class="error-message">{{ error }}</p>
    </div>
  </div>
</template>


<script>
import axios from "@/services/axios"; // Import the Axios instance for making HTTP requests

export default {
  // Data properties for the component
  data() {
    return {
      username: "", // Stores the username entered by the user
      error: null, // Stores any error message to be displayed to the user
    };
  },
  methods: {
    // Login function to handle user login
    async login() {
      // Check if username is provided
      if (!this.username) {
        this.error = "Username is required"; // Set error message if username is missing
        return;
      }
      try {
        
        // Make a POST request to the /session endpoint with the username
        const response = await axios.post("/session", { username: this.username });

        // safe Identifier in sessionStorage
        sessionStorage.setItem("identifier", response.data.identifier);

        // safe current username in sessionStorage
        sessionStorage.setItem("currentUser", this.username)

        // On success, redirect to the chat list page
        this.$router.push("/chats");
      } catch (err) {
        // Error handling for different error types
        if (err.response) {
          // Error coming from the backend (API response)
          if (err.response.data && err.response.data.error) {
            this.error = err.response.data.error; // Custom error message from the API
          } else {
            this.error = "An unexpected error occurred"; // General error message for unknown backend errors
          }
        } else if (err.request) {
          // Error due to network issues (e.g., no response from the server)
          this.error = "Network error, please try again later.";
        } else {
          // Other errors (e.g., misconfigured request or unexpected client-side issues)
          this.error = "An unexpected error occurred";
        }
      }
    },
  },
};
</script>

<style scoped>
.login-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f9fafb;
  font-family: 'Inter', sans-serif;
}

.login-card {
  background: white;
  padding: 40px;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 360px;
  text-align: center;
}

.login-card h1 {
  font-size: 1.8rem;
  margin-bottom: 24px;
  color: #111827;
}

.login-input {
  width: 100%;
  padding: 12px;
  margin-bottom: 16px;
  font-size: 1rem;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  outline: none;
  transition: border 0.2s ease;
}

.login-input:focus {
  border-color: #6366f1;
}

.login-button {
  width: 100%;
  padding: 12px;
  background-color: #6366f1;
  color: white;
  font-weight: 500;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.login-button:hover {
  background-color: #4f46e5;
}

.error-message {
  margin-top: 12px;
  color: #ef4444;
  font-size: 0.95rem;
}

</style>
