<!-- 

Pagina utilizzata per visulizzare una conversazione con un utente o con un gruppo

L'utente loggato può:
- inviare messaggi testuali o con foto
- eliminare messaggi
- inoltrare messaggi
- lasciare commenti ai messaggi e eliminarli
- visualizzare i commenti ai messaggi
- visualizzare i messaggi della conversazione

-->

<script>
import Modal from '../components/ModalConv.vue';
import Comments from '../components/ModalComments.vue';

export default {
  data() {
    return {
      errormsg: null,

      // Dati della conversazione
      userToSend: localStorage.username,  // Nome dell'utente con cui si sta conversando o del gruppo
      userIdToSend: localStorage.userID,  // ID dell'utente con cui si sta conversando o del gruppo
      proPic64: localStorage.photo, // Foto dell'utente con cui si sta conversando o del gruppo
      convId: this.$route.params.convId,  // ID della conversazione (se è una nuova conversazione, sarà undefined)

      // Testo del messaggio da inviare
      text: null,
      // Foto del messaggio da inviare
      photo: null,

      // Lista dei messaggi della conversazione
      some_data: [],

      message: [],
      
      isConversationEmpty: Boolean,

      // Utilizzata per verificare se si sta conversando è un gruppo o no
      isGroup: localStorage.users ? true : false,

      // Utilizzato per mostrare il modale di ricerca di una conversazione a cui inoltrare un messaggio selezionato
      searchModalIsVisible: false,

      // Utilizzato per mostrare il modale per lasciare un commento al messaggio selezionato
      commentModalIsVisible: false,

      // Messaggio da inoltrare
      messageToFordward: null,

      // Messaggio da commentare
      messageToComment: null,

      // Id dell'utente loggato
      userId: sessionStorage.userID,

      // Commenti del messaggio selezionato
      comments: null,

      // Id dell'intervallo
      intervalId: null,
    }
  },
  emits: ['login-success'],
  methods: {
    // Funzione utilizzata per controllare se il file inserito dall'utente è del formato corretto
    async handleFileChange(event) {
      this.errorMsg = "";
      const file = event.target.files[0]; // Prende il file inserito dall'utente
      if (file.type !== "jpeg" || file.type !== "png") {
        this.errorMsg = "File type not supported, only jpg and jpeg are allowed";
        return
      }
      if (file.size > 5242880) {
        this.errorMsg = "File size is too big. Max size is 5MB";
        return
      }
      this.photo = file;  // Assegna il file inserito dall'utente alla variabile photo
    },
    // Funzione che prende i messaggi della conversazione
    async getConversation() {
  this.errormsg = null;
  
  // Verifica che convid esista
  if (!this.convId) {
    console.error('ID conversazione non presente');
    return;
  }
  
  try {
    const response = await this.$axios.get(`/messages/${this.convId}`, { 
      headers: { 'Authorization': sessionStorage.token } 
    });
    console.log(response.data)
    
    // Gestisci il caso di conversazione vuota
    if (!response.data || response.data.length === 0) {
      console.log('Nessun messaggio nella conversazione');
      // Puoi impostare un messaggio di default o uno stato specifico
      this.some_data = [];
      this.isConversationEmpty = true;
    } else {
      this.some_data = response.data;
      this.isConversationEmpty = false;
    }
  } catch (e) {
    this.errormsg = e.toString();
    console.error('Errore nel recupero messaggi:', e);
  }
},
    // Funzione che controlla se la convId è stata presa correttamente dai parametri
    async check() {
  this.errormsg = null;

  if (this.convId &&  this.convId !== undefined  || !isNaN(this.convId)) {
    console.log("oh no o fatp dano")
  this.sendMessage();
  return;
}

try {
    console.log("creo nuova conversazione perché convId non valido:", this.convId);
    
    await this.createConversation();
    this.sendMessage();
  } catch (e) {
    this.errormsg = e.toString();
  }
},

createConversation() {
 this.errormsg = null;
this.$axios.put(`/chats/${sessionStorage.userID}/CreatePrivateConversation/${this.userIdToSend}`, {
headers: { 'Authorization': sessionStorage.token }
 }) .then(async response => {
// Assegna la convId
this.convId = response.data.conversationId;
await this.$router.push(`/chats/${sessionStorage.userID}/${this.convId}`);
window.location.reload();
 }) .catch (e => {
this.errormsg = e.toString();
 })},

// Funzione che elimina un messaggio
async deleteMessage(msgId) {
this.errormsg = null;
// Effettua la richiesta al server per eliminare il messaggio selezionato
this.$axios.delete(`/chats/${sessionStorage.userID}/${this.convId}/delete_message/${msgId}`, { headers: { 'Authorization': sessionStorage.token } })
 .then(() => {
 })
 .catch(e => {
this.errormsg = e.toString();
 });
 },
    async sendMessage() {
  this.errormsg = null;
  console.log("Invio messaggio con text:", this.text, "photo:", this.photo);

  const formData = new FormData();
  formData.append('content', this.text);
  if (this.photo) {
    formData.append('file', this.photo);
  }
  console.log(formData.entries)

  this.$axios.post(`/chats/${sessionStorage.userID}/${this.convId}/messages`, formData, { 
    headers: { 'Authorization': sessionStorage.token } 
  })
  .then(response => {
    console.log("Messaggio inviato correttamente:", response.data);
  })
  .catch(e => {
    console.error("Errore nel POST /messages:", e.response);
    this.errormsg = e.toString();
  });
},
    // Funzione utilizzata per mostrare o nascondere il modale per lasciare un commento al messaggio selezionato
    handleCommentModalToggle(cmt) {
      // Assegna il messaggio selezionato e i commenti del messaggio selezionato alle variabili messageToComment e comments
      this.messageToComment = cmt;
      // Mostra o nasconde il modale
      this.commentModalIsVisible = !this.commentModalIsVisible;
    },
    // Funzione utilizzata per mostrare o nascondere il modale per inoltrare un messaggio selezionato
    handleSearchModalToggle(msg) {
      // Assegna il messaggio selezionato alla variabile messageToFordward
      this.messageToFordward = msg;
      // Mostra o nasconde il modale
      this.searchModalIsVisible = !this.searchModalIsVisible;
    },
    // Funzione utlizzata per andare alla pagina delle infromazioni di un gruppo (utilizzata solo nel caso in cui la conversazione è con un gruppo)
    goToGroupInfo() {
      // Reindirizza alla pagina delle informazioni del gruppo
      this.$router.push(`/chats/${sessionStorage.userID}/${this.convId}/GetGroup`);
    },
    // Funzione che rimuove il commento dell'utente loggato
    async uncommentMessage(cmtId) {
      this.errormsg = null;
      // Effettua una richietsa DELETE per rimuovere il commento dell'utente loggato
      const url = `/chats/${sessionStorage.userID}/${this.convId}/comment/${cmtId}`;
      this.$axios.delete(url, { headers: { 'Authorization': sessionStorage.token } })
        .then(() => {
        })
        .catch(e => {
          this.errormsg = e.toString();
        });
    },
  },

  mounted() {
  if (!sessionStorage.token) {
    this.$router.push("/");
    return;
  }
  
  // Assicurati che convid esista prima di chiamare getConversation
  if (this.convId) {
    this.getConversation();
    
    // Riduce l'intervallo e aggiunge un controllo
    this.intervalId = setInterval(() => {
      if (this.convId) {
        this.getConversation();
      }
    }, 1000); 
  } else {
    console.error('ID conversazione mancante');
  }
},
beforeUnmount() {
    // Pulisci l'intervallo quando il componente viene distrutto
    if (this.intervalId) {
      clearInterval(this.intervalId);
    }
  },
  components: { Modal, Comments },
}
</script>

<template>
  <div>
    <div
      class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <!-- User photo -->
      <div class="top-profile-container">
        <img :src="`data:image/jpg;base64,${proPic64}`">
      </div>
      <!-- Controlla se la conversazione è con un gruppo o con un utente -->
      <div v-if="isGroup">
        <!-- Se è un gruppo mostra il nome del gruppo -->
        <h1 class="h1 clickable" @click="goToGroupInfo">{{ this.userToSend }}</h1>
      </div>
      <div v-else>
        <!-- Se è un utente mostra il nome dell'utente -->
        <h1 class="h1">{{ this.userToSend }}</h1>
      </div>

      <!-- Modali della pagina -->

      <!-- Modale utilizzato per lasciare un commento a un messaggio -->
      <Comments :show="commentModalIsVisible" :msg="messageToComment"
        @close="handleCommentModalToggle" title="comments">
        <template>
          <h3>Comments</h3>
        </template>
      </Comments>
      <!-- Modale utilizzato per selezionare una conversazione in cui inoltrare un messaggio -->
      <Modal :show="searchModalIsVisible" :msg="messageToFordward" @close="handleSearchModalToggle"
        title="conversations">
        <template v-slot:header>
          <h3>Conversations</h3>
        </template>
      </Modal>

      <!-- Body della pagina -->
      <div class="btn-toolbar mb-2 mb-md-0">
        <!-- Form per inviare una foto -->
        <input type="file" ref="file" accept=".jpg,.jpeg" @change="handleFileChange" />
        <!-- Input per invaire un messaggio testuale -->
        <div class="input-group">
          <input type="text" class="form-control" v-model="text" placeholder="Type your message here">
          <button class="btn btn-outline-primary" @click="check">Send</button>
        </div>
      </div>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

    <!-- Resto del codice -->
    
    <!-- Messaggio per conversazioni vuote -->
    <div v-if="isConversationEmpty" class="empty-conversation-message">
      Non ci sono ancora messaggi in questa conversazione. 
      Inizia a chattare!
    </div>

<!-- Lista dei messaggi della conversazione -->
<div class="messages" v-for="msg in some_data" :key="msg.messageId">

  <!-- Mittente -->
  <p v-if="msg.text !== 'null' || msg.photo !== ''">
    {{ msg.SenderUserId }}
  </p>

  <!-- Testo del messaggio -->
  <p v-if="msg.text && msg.text !== 'null'">
    {{ msg.text }}
  </p>

  <!-- Foto allegata -->
  <img class="msg_photo" v-if="msg.photo" :src="`data:image/jpg;base64,${msg.photo}`" alt="Message Photo">

  <!-- Data e stato -->
  <p v-if="msg.text !== 'null' || msg.photo !== ''">
    {{ msg.timeMsg }}
    <span v-if="msg.status === 'Sended'">✔️</span>
    <span v-if="msg.status === 'read'">✔️✔️</span>
  </p>

  <!-- Commenti -->
  <div v-for="cmt in msg.comments" :key="cmt.commentId">
    <p>{{ cmt.comment }} : {{ cmt.commentUsername }}</p>
    <button v-if="cmt.commentUserId == userId" type="button" class="btn btn-sm btn-outline-secondary"
      @click="uncommentMessage(cmt.commentId)">
      Delete comment
    </button>
  </div>

  <!-- Azioni -->
  <div class="btn-group me-2">
    <button type="button" class="btn btn-sm btn-outline-secondary" @click="handleSearchModalToggle(msg)">
      Forward Message
    </button>
    <button type="button" class="btn btn-sm btn-outline-secondary" @click="deleteMessage(msg.MessageId)">
      Delete message
    </button>
    <button v-if="msg.SenderUserId != userId" type="button" class="btn btn-sm btn-outline-secondary"
      @click="handleCommentModalToggle(msg)">
      Comment message
    </button>
  </div>

  <hr v-if="msg.text !== 'null' || msg.photo !== ''">
</div>
</div>
</template>

<style>
/* Stile utilizzato per visualizzare l'immagine profilo dell'utente o del gruppo con cui si sta conversando */
.top-profile-container {
  width: auto;
  height: auto;

  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
}

/* Stile utilizzato per visualizzare la foto di un messaggio */
.msg_photo {
  width: 25%;
  height: 25%;
}

/* Stile utilizzato nel caso in cui la conversazione è con un gruppo */
.clickable {
  cursor: pointer;
  color: blue;
  text-decoration: underline;
}

.clickable:hover {
  color: darkblue;
}
</style>
