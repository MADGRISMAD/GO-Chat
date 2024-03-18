<template>
  <div class="container">
    <div class="d-flex flex-column align-items-stretch flex-shrink-0 bg-body-tertiary">
      <div class="d-flex align-items-center flex-shrink-0 p-3 link-body-emphasis text-decoration-none border-bottom">

        <span class="fs-5 fw-semibold">List group</span>
        <input class="fs-5 fw-semibold" v-model="username" />
      </div>
      <div class="list-group list-group-flush border-bottom scrollarea">
        <div class="list-group-item list-group-item-action  py-3 lh-sm" aria-current="true" v-for="message in messages"
          :key="message">
          <div class="d-flex w-100 align-items-center justify-content-between">
            <strong class="mb-1">{{ message.username }}</strong>
          </div>
          <div class="col-10 mb-1 small">{{ message.message }}</div>
        </div>
      </div>
    </div>
    <form @submit.prevent="submit">
      <input class="form-control" placeholder="escribe un mensaje" v-model="message" />
    </form>

  </div>

</template>

<script>
import { ref, onMounted } from 'vue'
import Pusher from 'pusher-js';
export default {
  name: 'App',
  setup() {
    const username = ref('username');
    const messages = ref([]);
    const message = ref('')

    onMounted(() => {
      // Enable pusher logging - don't include this in production
      Pusher.logToConsole = true;

      const pusher = new Pusher('a19dafb420f8035c9a03', {
        cluster: 'us2'
      });

      const channel = pusher.subscribe('chat');
      channel.bind('message', data => {
        messages.value.push(data);
      });

    })

    const submit = async () => {
      await fetch('http://localhost:8000/api/messages', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username: username.value, message: message.value })
      })
    

      message.value = '';
      
    }

    return { username, messages, message, submit }
  }
}

</script>

<style>
.scrollarea {
  overflow-y: auto;
  min-height: 500px;
}
</style>
