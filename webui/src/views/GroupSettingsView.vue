<template>
  <div class="settings-wrapper">
    <div class="settings-card">
      <h1 class="title">Group Settings</h1>

      <img class="avatar" :src="groupPhotoURL" alt="Group Avatar">
      <input type="file" class="file-input" accept="image/*" @change="handleFileUpload">
      <button class="btn primary" @click="uploadGroupPhoto">Upload Group Photo</button>

      <input v-model="groupname" type="text" class="group-input" placeholder="Group name">

      <div class="actions">
        <button class="btn secondary" @click="saveGroupName">Save</button>
        <button class="btn info" @click="addUserToGroup">Add User</button>
        <button class="btn danger" @click="leaveGroup">Leave Group</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "@/services/axios";

export default {
  data() {
    return {
      groupname: this.$route.query.groupname || "Unnamed Group",
      groupPhotoURL: "https://ui-avatars.com/api/?name=Group&size=100",
      imageFile: null,
    };
  },
  async mounted() {
    await this.fetchGroupData();
  },
  methods: {
    async fetchGroupData() {
      try {
        const response = await axios.get(`/groups/${this.groupname}`);
        if (response.data) {
          this.groupPhotoURL = response.data.group_photo_url || this.groupPhotoURL;
        }
      } catch (error) {
        console.error("Error fetching group data:", error);
      }
    },
    handleFileUpload(event) {
      this.imageFile = event.target.files[0];
    },
    async uploadGroupPhoto() {
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

        await axios.put(`/groups/${this.groupname}/photo`, { photo_url: imageUrl });

        this.groupPhotoURL = imageUrl;
        alert("Group photo updated!");
      } catch (error) {
        console.error("Error uploading group photo:", error);
        alert("Failed to upload group photo.");
      }
    },
    async saveGroupName() {
      try {
        await axios.put(`/groups/${this.groupname}/rename`, { newname: this.groupname });
        alert("Group name updated!");
      } catch (error) {
        console.error("Error updating group name:", error);
        alert("Failed to update group name.");
      }
    },
    async addUserToGroup() {
      const newUser = prompt("Enter username to add:");
      if (!newUser) return;
      try {
        await axios.post(`/groups/${this.groupname}/add-user`, { username: newUser });
        alert(`${newUser} added to the group!`);
      } catch (error) {
        console.error("Error adding user to group:", error);
        alert("Failed to add user.");
      }
    },
    async leaveGroup() {
      if (!confirm("Are you sure you want to leave this group?")) return;
      try {
        await axios.post(`/groups/${this.groupname}/leave`);
        alert("You have left the group.");
        this.$router.push("/home");
      } catch (error) {
        console.error("Error leaving group:", error);
        alert("Failed to leave group.");
      }
    },
  },
};
</script>

<style scoped>
.settings-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f3f4f6;
  font-family: 'Inter', sans-serif;
}

.settings-card {
  background: white;
  border-radius: 16px;
  padding: 48px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
  text-align: center;
  width: 100%;
  max-width: 400px;
}

.title {
  font-size: 1.8rem;
  font-weight: 600;
  margin-bottom: 24px;
  color: #1f2937;
}

.avatar {
  width: 120px;
  height: 120px;
  object-fit: cover;
  border-radius: 50%;
  border: 3px solid #e5e7eb;
  margin-bottom: 16px;
}

.file-input {
  margin-bottom: 16px;
}

.group-input {
  width: 100%;
  padding: 12px;
  margin: 16px 0;
  font-size: 1rem;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  outline: none;
  transition: border-color 0.2s;
}

.group-input:focus {
  border-color: #6366f1;
}

.actions {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.btn {
  padding: 10px 16px;
  border-radius: 8px;
  font-weight: 500;
  font-size: 1rem;
  border: none;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.btn.primary {
  background-color: #6366f1;
  color: white;
}

.btn.primary:hover {
  background-color: #4f46e5;
}

.btn.secondary {
  background-color: #4b5563;
  color: white;
}

.btn.secondary:hover {
  background-color: #374151;
}

.btn.info {
  background-color: #0ea5e9;
  color: white;
}

.btn.info:hover {
  background-color: #0284c7;
}

.btn.danger {
  background-color: #ef4444;
  color: white;
}

.btn.danger:hover {
  background-color: #dc2626;
}

</style>
