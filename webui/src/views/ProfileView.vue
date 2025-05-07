<template>
  <div class="profile-wrapper">
    <div class="profile-card">
      <h1 class="profile-title">Profile Settings</h1>

      <img class="profile-avatar" :src="profilePhotoURL" alt="User Avatar" />

      <input type="file" @change="handleFileUpload" accept="image/*" />
      <button class="primary-button" @click="uploadProfilePicture">Upload Picture</button>

      <input
        type="text"
        class="profile-input"
        v-model="username"
        placeholder="Enter username"
      />

      <ErrorMsg v-if="errorMsg" :message="errorMsg" />

      <div class="button-group">
        <button class="primary-button" @click="saveProfile">Save</button>
      </div>
    </div>
  </div>
</template>


<script>
import axios from "@/services/axios";
import ErrorMsg from '@/components/ErrorMsg.vue';

export default {
  components: { ErrorMsg },
  data() {
    return {
      username: "",
      profilePhotoURL: 'https://ui-avatars.com/api/?name=empty&size=100',
      errorMsg: "",
      imageFile: null
    };
  },
  methods: {
    async fetchUser() {
      try {
        const storedUsername = sessionStorage.getItem("currentUser");
        
        if (!storedUsername) {
          console.error("No user found in sessionStorage.");
          return;
        }

        this.username = storedUsername;

        const response = await axios.get(`/users`, {
          params: { username: this.username },
        });


        if (response.data && response.data[0].username) {
          this.username = response.data[0].username;
          sessionStorage.setItem("currentUser", this.username);
          this.profilePhotoURL = response.data[0].profile_photo_url || this.profilePhotoURL;
        }
      } catch (error) {
        console.error("Error fetching user:", error);
      }
    },

    async saveProfile() {
      try {
        this.errorMsg = "";
        const response = await axios.put("/user-profile", { newusername: this.username });
        
        if (response.data.message === 'username successfully changed') {
          alert('Username updated successfully!');
        }
      } catch (error) {
        console.error('Error saving profile:', error);
        this.errorMsg = error.response?.data?.error || "An unexpected error occurred.";
      }
    },

    handleFileUpload(event) {
      this.imageFile = event.target.files[0];
    },

    async uploadProfilePicture() {
      if (!this.imageFile) {
        alert("Please select an image first.");
        return;
      }

      try {
        const formData = new FormData();
        formData.append("image", this.imageFile);

        const uploadResponse = await axios.post("/upload", formData, {
          headers: { "Content-Type": "multipart/form-data" },
        });

        const imageUrl = uploadResponse.data.imageUrl;

        // API-Call zum Aktualisieren des Profilbildes
        await this.changeProfilePicture(imageUrl);

      } catch (error) {
        console.error("Error uploading profile picture:", error);
        alert("Failed to upload profile picture.");
      }
    },

    async changeProfilePicture(newPhotoURL) {
      try {
        const response = await axios.put('/profile-picture', { photo_url: newPhotoURL });
        if (response.data.message === 'profile picture successfully updated') {
          this.profilePhotoURL = newPhotoURL;
          alert('Profile picture updated!');
        }
      } catch (error) {
        console.error('Error changing profile picture:', error);
        alert('Failed to update profile picture.');
      }
    },

  },

  mounted() {

    this.fetchUser();

  },
};
</script>

<style scoped>
.profile-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f3f4f6;
  font-family: 'Inter', sans-serif;
}

.profile-card {
  background: white;
  border-radius: 12px;
  padding: 40px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 420px;
  text-align: center;
}

.profile-title {
  font-size: 1.8rem;
  margin-bottom: 24px;
  color: #111827;
}

.profile-avatar {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  margin-bottom: 20px;
  object-fit: cover;
}

input[type="file"] {
  margin-bottom: 12px;
}

.profile-input {
  font-size: 1rem;
  padding: 10px;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  width: 100%;
  margin-bottom: 16px;
}

.primary-button {
  background-color: #6366f1;
  color: white;
  border: none;
  padding: 10px 16px;
  border-radius: 8px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.primary-button:hover {
  background-color: #4f46e5;
}

.button-group {
  display: flex;
  justify-content: center;
  gap: 10px;
}
</style>
