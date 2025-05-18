<template>
  <div class="profile-settings">
    <h1 class="profile-title">Profile</h1>
    <div class="profile-card">
      <!-- Profilbild -->
      <img class="profile-avatar" :src="profilePhotoURL" alt="User Avatar" />

      <!-- Datei-Upload für das Profilbild -->
      <input type="file" @change="handleFileUpload" accept="image/*" />
      <button class="upload-button" @click="uploadProfilePicture">Upload Picture</button>

      <input type="text" class="profile-name" v-model="username" placeholder="Enter username" />
      <ErrorMsg v-if="errorMsg" :message="errorMsg" />
      <div class="button-group">
        <button class="save-button" @click="saveProfile">Save</button>
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
.profile-settings {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  min-height: 100vh;
  background: linear-gradient(135deg, #6a82fb 0%, #fc5c7d 100%);
  font-family: 'Inter', 'Roboto', Arial, sans-serif;
  color: #222;
  padding-top: 60px;
}

.profile-title {
  font-size: 2.5rem;
  font-weight: 700;
  margin-bottom: 32px;
  color: #fff;
  letter-spacing: 1px;
  text-shadow: 0 2px 8px rgba(0,0,0,0.12);
}

.profile-card {
  background: rgba(255,255,255,0.95);
  border-radius: 18px;
  padding: 48px 40px 36px 40px;
  display: flex;
  flex-direction: column;
  align-items: center;
  box-shadow: 0 8px 32px rgba(80, 80, 120, 0.18);
  min-width: 340px;
  max-width: 95vw;
}

.profile-avatar {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  margin-bottom: 24px;
  object-fit: cover;
  border: 4px solid #6a82fb;
  box-shadow: 0 2px 12px rgba(106,130,251,0.12);
  background: #f0f0f0;
}

input[type="file"] {
  margin-bottom: 10px;
  font-size: 1rem;
  color: #444;
  background: none;
  border: none;
}

.profile-name {
  font-size: 1.2rem;
  padding: 14px;
  border: 1.5px solid #e0e0e0;
  border-radius: 8px;
  text-align: center;
  width: 260px;
  margin-bottom: 18px;
  background: #f7f8fa;
  transition: border 0.2s;
}

.profile-name:focus {
  border: 1.5px solid #6a82fb;
  outline: none;
  background: #fff;
}

.upload-button, .save-button {
  margin-top: 10px;
  padding: 12px 28px;
  border: none;
  border-radius: 8px;
  background: linear-gradient(90deg, #6a82fb 0%, #fc5c7d 100%);
  color: #fff;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  box-shadow: 0 2px 8px rgba(106,130,251,0.10);
  transition: background 0.2s, transform 0.1s;
}

.upload-button:hover, .save-button:hover {
  background: linear-gradient(90deg, #fc5c7d 0%, #6a82fb 100%);
  transform: translateY(-2px) scale(1.03);
}

.button-group {
  display: flex;
  gap: 16px;
  margin-top: 8px;
}

@media (max-width: 500px) {
  .profile-card {
    padding: 24px 10px 20px 10px;
    min-width: unset;
    width: 98vw;
  }
  .profile-name {
    width: 90vw;
    font-size: 1rem;
  }
}
</style>
