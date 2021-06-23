<template>
  <div>
    <v-parallax src="https://cdn.vuetifyjs.com/images/parallax/material.jpg"></v-parallax>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'

export default {
  name: 'home',
  components: {
    // HelloWorld
  },
  methods: {
    isAuthenticated() {
      if (this.$parent.authenticated) {
        return true;
      } else {
        return false;
      }
    }
  }
}
</script>

<style scoped>
.jumbotron-bg {
  background: linear-gradient(rgba(255, 245, 235, 0.75)), url("https://content.parsableweb.com/bg-pattern.png") repeat center center;
  background-size: cover;
  background-position-y: 60px;
  background-size: 600px;
  
  /* background-position-y: -150px; */
  /* background-repeat: repeat; */
}
.jumbo-shadow {
  text-shadow: 1px 1px 5px rgba(255, 255, 255, 0.4);
}
.lead-shadow {
}
.lead-color {
}
</style>