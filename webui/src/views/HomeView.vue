<!--

Pagina princiaple del sito in cui vengono mostrate le conversazioni dell'utente loggato.

In questa pagina l'utente può:
- visualizzare le conversazioni con altri utenti o gruppi
- aggiornare la lista delle conversazioni
- creare un nuovo gruppo
- creare una nuova conversazione con un utente

-->

<script>
import Modal from '../components/Modal.vue';
import Group from '../components/ModalGroup.vue';

export default {
  data: function () {
    return {
      errormsg: null,

      // Dati delle conversazioni
      some_data: [],

      // Utilizzato per mostrare o nascondere il modal di ricerca
      searchModalIsVisible: false,
      // Utilizzato per mostrare o nascondere il modal di creazione di un gruppo
      createGroupModalIsVisible: false,

      // Lista di utenti del nuovo gruppo
      users: [],

      // Id dell'intervallo
      intervalId: null,
    }
  },
  emits: ['login-success', 'username-changed'],
  methods: {
    // Fiunzione per ottente la conversazioni dell'utente
    async getConversations() {
      this.errormsg = null;
      try {
        // Effettua la richiesta al server per ottenere le conversazioni dell'utente
        let response = await this.$axios.get(`/chats/${sessionStorage.userID}`, { headers: { 'Authorization': sessionStorage.token } });

        
        // Salva i dati delle conversazioni
        this.some_data = response.data;
      }catch (e) {
        this.errormsg = e.toString();
    }
    },
    // Funzione utilizzata per poratre l'utente alla conversazione selezionata
    goToConversation(response) {
  // Sicurezza: controlla se conversation esiste
  if (response.conversation) {
    console.log(response.conversation)
    this.$router.push(`/chats/${sessionStorage.userID}/${response.conversation.conversationId}`);
  } else {
    console.log(response.conversation)
    console.warn("La conversazione non è definita:", response);
    this.errormsg = "Errore: conversazione non trovata.";
  }
  // Se è una conversazione di gruppo
  if (response.group && response.group.groupId !== 0) {
    localStorage.userID = response.group.groupId;
    localStorage.username = response.group.groupName;
    localStorage.photo = response.data.group.photo;
    localStorage.users = JSON.stringify(response.groupUsers);
  } 
  // Se è una conversazione privata
  else if (response.user) {
    localStorage.userID = response.data.user.UserId;
    localStorage.username = response.data.user.username;
    localStorage.photo = response.data.user.Photo;
  }


},

    // Funzione per mostrare o nascondere il modal di ricerca di utenti per iniziare una nuova conversazione
    handleSearchModalToggle() {
      this.searchModalIsVisible = !this.searchModalIsVisible;
    },
    // Funzione per mostare o nascondere il modale per la creazione di un nuovo gruppo
    handleCreateGroupModalToggle() {
      this.createGroupModalIsVisible = !this.createGroupModalIsVisible;
    },
  },
  mounted() {
    // Se l'utente non è loggato, reindirizza alla pagina di login
    if (!sessionStorage.token) {
      this.$router.push("/");
      return;
    }
    this.getConversations();
    this.intervalId = setInterval(async () => {
      clearInterval(this.intervalId);
      await this.getConversations();
      this.intervalId = setInterval(this.getConversations, 1000);
    }, 1000);
  },
  beforeUnmount() {
    // Pulisci l'intervallo quando il componente viene distrutto
    if (this.intervalId) {
      clearInterval(this.intervalId);
    }
  },
  components: { Modal, Group }
}
</script>

<template>
  <div>
    <div
      class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <h1 class="h2">Home page</h1>

      <!-- Modale utilizzato per la creazione di un nuovo gruppo -->
      <Group :show="createGroupModalIsVisible" @close="handleCreateGroupModalToggle" title="search">
        <template v-slot:header>
          <h3>Select users</h3>
        </template>
      </Group>
      <!-- Modale utilzzato per la ricerca degli utenti con cui aprire una nuova conversazione -->
      <Modal :show="searchModalIsVisible" @close="handleSearchModalToggle" title="search">
        <template v-slot:header>
          <h3>Users</h3>
        </template>
      </Modal>

      <!-- Pulsanti per aggiornare la lista delle conversazioni, creare un nuovo gruppo e cercare nuovi utenti -->
      <div class="btn-toolbar mb-2 mb-md-0">
        <div class="btn-group me-2">
          <!-- Pulsante per aggiornare la lista delle conversazioni -->
          <button type="button" class="btn btn-sm btn-outline-secondary" @click="getConversations">
            Refresh
          </button>
        </div>
        <!-- Pulsante per creare un nuovo gruppo -->
        <div class="btn-group me-2">
          <button type="button" class="btn btn-sm btn-outline-primary" @click="handleCreateGroupModalToggle">
            New Group
          </button>
        </div>
        <!-- Pulsante per cercare nuovi utenti e aprire un nuova conversazione -->
        <div class="btn-group me-2">
          <button type="button" class="btn btn-sm btn-outline-primary" @click="handleSearchModalToggle">
            New Chat
          </button>
        </div>
      </div>
    </div>

    <!-- Lista delle conversazioni -->
    <div v-if="some_data.length !== 0">
      <!-- Mostra le conversazioni dell'utente iterando all'interno di some_data dove sono salvate tutte le conversazioni -->
      <div class="conversations" v-for="response in some_data" :key="response.conversation.conversationId">
        <!-- Controlla se la conversazione non è con un gruppo -->
        <div v-if="response.group.groupName == ''">
          <!-- Mostra il nome dell'utente con cui si sta conversando, l'ultimo messaggio e chi lo ha inviato -->
          <!-- Se il messaggio è una foto, mostra "Photo" al posto del testo -->
          <button v-if="response.message.photo == ''" type="button" class="btn btn-sm btn-outline-primary"
            @click="goToConversation(response)">
            {{ response.User_to_send.username }} <br> {{ response.message.senderuserName }}: {{ response.message.text }}
          </button>
          <!-- Altrimenti mostra il contenuto del messaggio -->
          <button type="button" class="btn btn-sm btn-outline-primary" @click="goToConversation(response)" v-else>
            {{ response.User_to_send.username }} <br> {{ response.message.senderuserName}}: Photo
          </button>
        </div>
        <!-- Se la conversazione è con un gruppo -->
        <div v-else>
          <!-- Mostra il nome del gruppo, l'ultimo messaggio e chi lo ha inviato come per il singolo utente -->
          <button v-if="response.message.photo == ''" type="button" class="btn btn-sm btn-outline-primary"
            @click="goToConversation(response)">
            {{ response.group.groupName }} <br> {{ response.message.senderuserName }}: {{ response.message.text }}
          </button>
          <button type="button" class="btn btn-sm btn-outline-primary" @click="goToConversation(response)" v-else>
            {{ response.group.groupName }} <br> {{ response.message.senderuserName}}: Photo
          </button>
        </div>
        <hr>
      </div>
    </div>

  



    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
  </div>
</template>

<style></style>
