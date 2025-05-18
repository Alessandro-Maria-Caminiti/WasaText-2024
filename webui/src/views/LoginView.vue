<template>
  <div class="login-container">
    <h1>Login</h1>
    <!-- Input field for the username -->
    <input v-model="username" placeholder="Enter your username" />
    <!-- Button to trigger the login function -->
    <button @click="login">Login</button>
    <!-- Display an error message if there is an error -->
    <p v-if="error" class="error">{{ error }}</p>
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
.login-container {
  max-width: 400px;
  margin: 100px auto;
  padding: 2.5rem 2rem 2rem 2rem;
  background: #fff;
  border-radius: 18px;
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.18);
  text-align: center;
  transition: box-shadow 0.2s;
}

.login-container h1 {
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
  color: #22223b;
  letter-spacing: 1px;
}

input {
  width: 100%;
  padding: 12px 14px;
  margin: 12px 0;
  border: 1.5px solid #e0e1dd;
  border-radius: 8px;
  background: #f8f9fa;
  font-size: 1rem;
  transition: border 0.2s, box-shadow 0.2s;
  outline: none;
}

input:focus {
  border: 1.5px solid #4f8cff;
  box-shadow: 0 0 0 2px #e3f0ff;
  background: #fff;
}

button {
  width: 100%;
  padding: 12px;
  background: linear-gradient(90deg, #4f8cff 0%, #4361ee 100%);
  color: #fff;
  font-weight: 600;
  font-size: 1.1rem;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  margin-top: 10px;
  box-shadow: 0 2px 8px rgba(79, 140, 255, 0.08);
  transition: background 0.2s, transform 0.1s;
}

button:hover {
  background: linear-gradient(90deg, #4361ee 0%, #4f8cff 100%);
  transform: translateY(-2px) scale(1.02);
}

.error {
  color: #e63946;
  background: #ffeaea;
  border-radius: 6px;
  padding: 8px 0;
  margin-top: 14px;
  font-size: 1rem;
  font-weight: 500;
  letter-spacing: 0.2px;
  box-shadow: 0 1px 4px rgba(230, 57, 70, 0.06);
}
</style>
